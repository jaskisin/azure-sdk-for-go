package azcosmos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	azlog "github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
)

const globalEndpointManagerDefaultBackgroundRefreshInterval = 5 * time.Minute

type globalEndpointManager struct {
	client           *Client
	preferredRegions []string
	locationCache    *locationCache
	refreshTicker    *time.Ticker
	refreshDone      chan struct{}
	sync.Mutex
}

// NewGlobalEndpointManager creates a new global endpoint manager.
func newGlobalEndpointManager(client *Client, preferredRegions []string, refreshInterval time.Duration) (*globalEndpointManager, error) {
	endpoint, err := url.Parse(client.endpoint)
	if err != nil {
		return nil, err
	}
	gem, err := &globalEndpointManager{
		client:           client,
		preferredRegions: preferredRegions,
		locationCache:    newLocationCache(preferredRegions, *endpoint),
	}, nil
	if err != nil {
		return nil, err
	}
	gem.stopBackgroundRefresh()
	gem.startBackgroundRefresh(refreshInterval)
	defer gem.stopBackgroundRefresh()

	return gem, nil
}

// GetWriteEndpoints returns the write endpoints from the location cache.
func (gem *globalEndpointManager) GetWriteEndpoints() ([]url.URL, error) {
	return gem.locationCache.writeEndpoints()
}

// GetReadEndpoints returns the read endpoints from the location cache.
func (gem *globalEndpointManager) GetReadEndpoints() ([]url.URL, error) {
	return gem.locationCache.readEndpoints()
}

// GetAccountProperties retrieves account properties from the Cosmos DB instance.
func (gem *globalEndpointManager) GetAccountProperties() (accountProperties, error) {
	operationContext := pipelineRequestOptions{
		resourceType:    resourceTypeDatabaseAccount,
		resourceAddress: "",
	}

	path, err := generatePathForNameBased(resourceTypeDatabaseAccount, "", false)
	if err != nil {
		return accountProperties{}, fmt.Errorf("failed to generate path for name-based request: %w", err)
	}

	azResponse, err := gem.client.sendGetRequest(path, context.Background(), operationContext, nil, nil)
	if err != nil {
		return accountProperties{}, fmt.Errorf("failed to send GET request: %w", err)
	}

	properties, err := newAccountProperties(azResponse)
	if err != nil {
		return accountProperties{}, fmt.Errorf("failed to parse account properties: %w", err)
	}

	return properties, nil
}

// newAccountProperties parses the account properties from the HTTP response.
func newAccountProperties(azResponse *http.Response) (accountProperties, error) {
	properties := accountProperties{}
	err := azruntime.UnmarshalAsJSON(azResponse, &properties)
	if err != nil {
		return properties, err
	}
	return properties, nil
}

// GetLocation returns the location for the given endpoint from the location cache.
func (gem *globalEndpointManager) GetLocation(endpoint url.URL) string {
	return gem.locationCache.getLocation(endpoint)
}

// MarkEndpointUnavailableForRead marks an endpoint as unavailable for read operations.
func (gem *globalEndpointManager) MarkEndpointUnavailableForRead(endpoint url.URL) error {
	return gem.locationCache.markEndpointUnavailableForRead(endpoint)
}

// MarkEndpointUnavailableForWrite marks an endpoint as unavailable for write operations.
func (gem *globalEndpointManager) MarkEndpointUnavailableForWrite(endpoint url.URL) error {
	return gem.locationCache.markEndpointUnavailableForWrite(endpoint)
}

// Update updates the location cache and the client's default endpoint based on the provided account properties.
func (gem *globalEndpointManager) Update() error {
	accountProperties, err := gem.GetAccountProperties()
	if err != nil {
		return fmt.Errorf("failed to retrieve account properties: %w", err)
	}

	if err := gem.locationCache.update(accountProperties.WriteRegions, accountProperties.ReadRegions, gem.preferredRegions, &accountProperties.EnableMultipleWriteLocations); err != nil {
		return fmt.Errorf("failed to update location cache: %w", err)
	}
	return nil
}

func (gem *globalEndpointManager) startBackgroundRefresh(refreshInterval time.Duration) {
	gem.Lock()
	defer gem.Unlock()

	if gem.refreshTicker != nil {
		// Background refresh already running, stop it first
		gem.refreshTicker.Stop()
		close(gem.refreshDone)
	}

	gem.refreshTicker = time.NewTicker(refreshInterval)
	gem.refreshDone = make(chan struct{})
	accountProperties, err := gem.GetAccountProperties()
	if err != nil {
		log.Write(azlog.EventResponse, fmt.Sprintf("Background refresh error: %v", err))
		return
	}
	go func() {
		for {
			select {
			case <-gem.refreshTicker.C:
				err := gem.locationCache.update(accountProperties.WriteRegions, accountProperties.ReadRegions, gem.preferredRegions, &accountProperties.EnableMultipleWriteLocations)
				if err != nil {
					log.Write(azlog.EventResponse, fmt.Sprintf("Background refresh error: %v", err))
				}
			case <-gem.refreshDone:
				gem.refreshTicker.Stop()
				close(gem.refreshDone)
				return
			}
		}
	}()
}

func (gem *globalEndpointManager) stopBackgroundRefresh() {
	gem.Lock()
	defer gem.Unlock()

	if gem.refreshTicker != nil {
		gem.refreshTicker.Stop()
		close(gem.refreshDone)
		gem.refreshTicker = nil
		gem.refreshDone = nil

		err := gem.Update()
		if err != nil {
			log.Write(azlog.EventResponse, fmt.Sprintf("Failed to update location cache: %v", err))
			return
		}
	}
}

// RefreshStaleEndpoints triggers a refresh of stale endpoints in the location cache.
func (gem *globalEndpointManager) RefreshStaleEndpoints() {
	gem.locationCache.refreshStaleEndpoints()
}

// IsEndpointUnavailable checks if an endpoint is marked as unavailable for the given requested operations.
func (gem *globalEndpointManager) IsEndpointUnavailable(endpoint url.URL, ops requestedOperations) bool {
	return gem.locationCache.isEndpointUnavailable(endpoint, ops)
}

// CanUseMultipleWriteLocations returns whether multiple write locations can be used.
func (gem *globalEndpointManager) CanUseMultipleWriteLocations() bool {
	return gem.locationCache.canUseMultipleWriteLocs()
}
