//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateEndpointConnections_ListByResource.json
func ExamplePrivateEndpointConnectionsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient("a0020869-4d28-422a-89f4-c2413130d73c",
		"<resource-group-name>",
		"<azure-resource-name>",
		"<private-endpoint-name>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourcePager("resourceGroup",
		"azureResourceName",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateEndpointConnections_Get.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient("a0020869-4d28-422a-89f4-c2413130d73c",
		"resourceGroup",
		"azureResourceName",
		"myPrivateEndpointName", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateEndpointConnections_Create.json
func ExamplePrivateEndpointConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient("a0020869-4d28-422a-89f4-c2413130d73c",
		"resourceGroup",
		"azureResourceName",
		"myPrivateEndpointName", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		armpowerbiprivatelinks.PrivateEndpointConnection{
			Properties: &armpowerbiprivatelinks.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armpowerbiprivatelinks.PrivateEndpoint{
					ID: to.Ptr("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpointName"),
				},
				PrivateLinkServiceConnectionState: &armpowerbiprivatelinks.ConnectionState{
					Description:     to.Ptr(""),
					ActionsRequired: to.Ptr("None"),
					Status:          to.Ptr(armpowerbiprivatelinks.PersistedConnectionStatus("Approved ")),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/PrivateEndpointConnections_Delete.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerbiprivatelinks.NewPrivateEndpointConnectionsClient("a0020869-4d28-422a-89f4-c2413130d73c",
		"resourceGroup",
		"azureResourceName",
		"myPrivateEndpointName", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}