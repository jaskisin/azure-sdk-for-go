//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"github.com/stretchr/testify/suite"
)

type LinkedServicesTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	cluserId          string
	clusterName       string
	workspaceName     string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *LinkedServicesTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/operationalinsights/armoperationalinsights/testdata")

	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.clusterName = testutil.GenerateAlphaNumericID(testsuite.T(), "linkedservicecluser", 6)
	testsuite.workspaceName = testutil.GenerateAlphaNumericID(testsuite.T(), "oilinkedservice", 6)
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *LinkedServicesTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestLinkedServicesTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedServicesTestSuite))
}

func (testsuite *LinkedServicesTestSuite) Prepare() {
	var err error
	// From step Workspaces_Create
	workspacesClient, err := armoperationalinsights.NewWorkspacesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	workspacesClientCreateOrUpdateResponsePoller, err := workspacesClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.workspaceName, armoperationalinsights.Workspace{
		Location: to.Ptr(testsuite.location),
		Properties: &armoperationalinsights.WorkspaceProperties{
			RetentionInDays: to.Ptr[int32](30),
			SKU: &armoperationalinsights.WorkspaceSKU{
				Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, workspacesClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step Cluster_Create
	clustersClient, err := armoperationalinsights.NewClustersClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	clustersClientCreateOrUpdateResponsePoller, err := clustersClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.clusterName, armoperationalinsights.Cluster{
		Location: to.Ptr(testsuite.location),
		Identity: &armoperationalinsights.Identity{
			Type: to.Ptr(armoperationalinsights.IdentityTypeSystemAssigned),
		},
		SKU: &armoperationalinsights.ClusterSKU{
			Name:     to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
			Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
		},
	}, nil)
	testsuite.Require().NoError(err)
	var clustersClientCreateOrUpdateResponse *armoperationalinsights.ClustersClientCreateOrUpdateResponse
	clustersClientCreateOrUpdateResponse, err = testutil.PollForTest(testsuite.ctx, clustersClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
	testsuite.cluserId = *clustersClientCreateOrUpdateResponse.ID
}

// Microsoft.OperationalInsights/workspaces/linkedServices
func (testsuite *LinkedServicesTestSuite) TestLinkedService() {
	var err error
	// From step LinkedServices_CreateOrUpdate
	linkedServicesClient, err := armoperationalinsights.NewLinkedServicesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	linkedServicesClientCreateOrUpdateResponsePoller, err := linkedServicesClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.workspaceName, "Cluster", armoperationalinsights.LinkedService{
		Properties: &armoperationalinsights.LinkedServiceProperties{
			WriteAccessResourceID: to.Ptr(testsuite.cluserId),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, linkedServicesClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step LinkedServices_ListByWorkspace
	linkedServicesClientNewListByWorkspacePager := linkedServicesClient.NewListByWorkspacePager(testsuite.resourceGroupName, testsuite.workspaceName, nil)
	for linkedServicesClientNewListByWorkspacePager.More() {
		_, err := linkedServicesClientNewListByWorkspacePager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step LinkedServices_Get
	_, err = linkedServicesClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.workspaceName, "Cluster", nil)
	testsuite.Require().NoError(err)

	// From step LinkedServices_Delete
	linkedServicesClientDeleteResponsePoller, err := linkedServicesClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, testsuite.workspaceName, "Cluster", nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, linkedServicesClientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.OperationalInsights/workspaces/availableServiceTiers
func (testsuite *LinkedServicesTestSuite) TestAvailableServiceTiers() {
	var err error
	// From step AvailableServiceTiers_ListByWorkspace
	availableServiceTiersClient, err := armoperationalinsights.NewAvailableServiceTiersClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = availableServiceTiersClient.ListByWorkspace(testsuite.ctx, testsuite.resourceGroupName, testsuite.workspaceName, nil)
	testsuite.Require().NoError(err)
}

// Microsoft.OperationalInsights/workspaces/managementGroups
func (testsuite *LinkedServicesTestSuite) TestManagementGroups() {
	var err error
	// From step ManagementGroups_List
	managementGroupsClient, err := armoperationalinsights.NewManagementGroupsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	managementGroupsClientNewListPager := managementGroupsClient.NewListPager(testsuite.resourceGroupName, testsuite.workspaceName, nil)
	for managementGroupsClientNewListPager.More() {
		_, err := managementGroupsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}
}

// Microsoft.OperationalInsights/workspaces/schema
func (testsuite *LinkedServicesTestSuite) TestSchema() {
	var err error
	// From step Schema_Get
	schemaClient, err := armoperationalinsights.NewSchemaClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = schemaClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.workspaceName, nil)
	testsuite.Require().NoError(err)
}

// Microsoft.OperationalInsights/workspaces/sharedKeys
func (testsuite *LinkedServicesTestSuite) TestSharedKeys() {
	var err error
	// From step SharedKeys_GetSharedKeys
	sharedKeysClient, err := armoperationalinsights.NewSharedKeysClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = sharedKeysClient.GetSharedKeys(testsuite.ctx, testsuite.resourceGroupName, testsuite.workspaceName, nil)
	testsuite.Require().NoError(err)
}

// Microsoft.OperationalInsights/workspaces/usages
func (testsuite *LinkedServicesTestSuite) TestUsages() {
	var err error
	// From step Usages_List
	usagesClient, err := armoperationalinsights.NewUsagesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	usagesClientNewListPager := usagesClient.NewListPager(testsuite.resourceGroupName, testsuite.workspaceName, nil)
	for usagesClientNewListPager.More() {
		_, err := usagesClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}
}