//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_get.json
func ExampleResourceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"myResourceGroup",
		"testHub",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_createOrUpdate.json
func ExampleResourceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"testHub",
		armiothub.Description{
			Location: to.Ptr("centraluseuap"),
			Tags:     map[string]*string{},
			Etag:     to.Ptr("AAAAAAFD6M4="),
			Properties: &armiothub.Properties{
				CloudToDevice: &armiothub.CloudToDeviceProperties{
					DefaultTTLAsIso8601: to.Ptr("PT1H"),
					Feedback: &armiothub.FeedbackProperties{
						LockDurationAsIso8601: to.Ptr("PT1M"),
						MaxDeliveryCount:      to.Ptr[int32](10),
						TTLAsIso8601:          to.Ptr("PT1H"),
					},
					MaxDeliveryCount: to.Ptr[int32](10),
				},
				EnableDataResidency:           to.Ptr(true),
				EnableFileUploadNotifications: to.Ptr(false),
				EventHubEndpoints: map[string]*armiothub.EventHubProperties{
					"events": {
						PartitionCount:      to.Ptr[int32](2),
						RetentionTimeInDays: to.Ptr[int64](1),
					},
				},
				Features:      to.Ptr(armiothub.CapabilitiesNone),
				IPFilterRules: []*armiothub.IPFilterRule{},
				MessagingEndpoints: map[string]*armiothub.MessagingEndpointProperties{
					"fileNotifications": {
						LockDurationAsIso8601: to.Ptr("PT1M"),
						MaxDeliveryCount:      to.Ptr[int32](10),
						TTLAsIso8601:          to.Ptr("PT1H"),
					},
				},
				MinTLSVersion: to.Ptr("1.2"),
				NetworkRuleSets: &armiothub.NetworkRuleSetProperties{
					ApplyToBuiltInEventHubEndpoint: to.Ptr(true),
					DefaultAction:                  to.Ptr(armiothub.DefaultActionDeny),
					IPRules: []*armiothub.NetworkRuleSetIPRule{
						{
							Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
							FilterName: to.Ptr("rule1"),
							IPMask:     to.Ptr("131.117.159.53"),
						},
						{
							Action:     to.Ptr(armiothub.NetworkRuleIPActionAllow),
							FilterName: to.Ptr("rule2"),
							IPMask:     to.Ptr("157.55.59.128/25"),
						}},
				},
				Routing: &armiothub.RoutingProperties{
					Endpoints: &armiothub.RoutingEndpoints{
						EventHubs:         []*armiothub.RoutingEventHubProperties{},
						ServiceBusQueues:  []*armiothub.RoutingServiceBusQueueEndpointProperties{},
						ServiceBusTopics:  []*armiothub.RoutingServiceBusTopicEndpointProperties{},
						StorageContainers: []*armiothub.RoutingStorageContainerProperties{},
					},
					FallbackRoute: &armiothub.FallbackRouteProperties{
						Name:      to.Ptr("$fallback"),
						Condition: to.Ptr("true"),
						EndpointNames: []*string{
							to.Ptr("events")},
						IsEnabled: to.Ptr(true),
						Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
					},
					Routes: []*armiothub.RouteProperties{},
				},
				StorageEndpoints: map[string]*armiothub.StorageEndpointProperties{
					"$default": {
						ConnectionString: to.Ptr(""),
						ContainerName:    to.Ptr(""),
						SasTTLAsIso8601:  to.Ptr("PT1H"),
					},
				},
			},
			SKU: &armiothub.SKUInfo{
				Name:     to.Ptr(armiothub.IotHubSKUS1),
				Capacity: to.Ptr[int64](1),
			},
		},
		&armiothub.ResourceClientBeginCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_patch.json
func ExampleResourceClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myHub",
		armiothub.TagsResource{
			Tags: map[string]*string{
				"foo": to.Ptr("bar"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_delete.json
func ExampleResourceClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"testHub",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listbysubscription.json
func ExampleResourceClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listbyrg.json
func ExampleResourceClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("myResourceGroup",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_stats.json
func ExampleResourceClient_GetStats() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetStats(ctx,
		"myResourceGroup",
		"testHub",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listehgroups.json
func ExampleResourceClient_NewListEventHubConsumerGroupsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListEventHubConsumerGroupsPager("myResourceGroup",
		"testHub",
		"events",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_getconsumergroup.json
func ExampleResourceClient_GetEventHubConsumerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetEventHubConsumerGroup(ctx,
		"myResourceGroup",
		"testHub",
		"events",
		"test",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_createconsumergroup.json
func ExampleResourceClient_CreateEventHubConsumerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateEventHubConsumerGroup(ctx,
		"myResourceGroup",
		"testHub",
		"events",
		"test",
		armiothub.EventHubConsumerGroupBodyDescription{
			Properties: &armiothub.EventHubConsumerGroupName{
				Name: to.Ptr("test"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_deleteconsumergroup.json
func ExampleResourceClient_DeleteEventHubConsumerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.DeleteEventHubConsumerGroup(ctx,
		"myResourceGroup",
		"testHub",
		"events",
		"test",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listjobs.json
func ExampleResourceClient_NewListJobsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListJobsPager("myResourceGroup",
		"testHub",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_getjob.json
func ExampleResourceClient_GetJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetJob(ctx,
		"myResourceGroup",
		"testHub",
		"test",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_quotametrics.json
func ExampleResourceClient_NewGetQuotaMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewGetQuotaMetricsPager("myResourceGroup",
		"testHub",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_routingendpointhealth.json
func ExampleResourceClient_NewGetEndpointHealthPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewGetEndpointHealthPager("myResourceGroup",
		"testHub",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/checkNameAvailability.json
func ExampleResourceClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx,
		armiothub.OperationInputs{
			Name: to.Ptr("test-request"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_testallroutes.json
func ExampleResourceClient_TestAllRoutes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.TestAllRoutes(ctx,
		"testHub",
		"myResourceGroup",
		armiothub.TestAllRoutesInput{
			Message: &armiothub.RoutingMessage{
				AppProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
				Body: to.Ptr("Body of message"),
				SystemProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
			},
			RoutingSource: to.Ptr(armiothub.RoutingSourceDeviceMessages),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_testnewroute.json
func ExampleResourceClient_TestRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.TestRoute(ctx,
		"testHub",
		"myResourceGroup",
		armiothub.TestRouteInput{
			Message: &armiothub.RoutingMessage{
				AppProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
				Body: to.Ptr("Body of message"),
				SystemProperties: map[string]*string{
					"key1": to.Ptr("value1"),
				},
			},
			Route: &armiothub.RouteProperties{
				Name: to.Ptr("Routeid"),
				EndpointNames: []*string{
					to.Ptr("id1")},
				IsEnabled: to.Ptr(true),
				Source:    to.Ptr(armiothub.RoutingSourceDeviceMessages),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listkeys.json
func ExampleResourceClient_NewListKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListKeysPager("myResourceGroup",
		"testHub",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_getkey.json
func ExampleResourceClient_GetKeysForKeyName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetKeysForKeyName(ctx,
		"myResourceGroup",
		"testHub",
		"iothubowner",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_exportdevices.json
func ExampleResourceClient_ExportDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ExportDevices(ctx,
		"myResourceGroup",
		"testHub",
		armiothub.ExportDevicesRequest{
			ExcludeKeys:            to.Ptr(true),
			ExportBlobContainerURI: to.Ptr("testBlob"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_importdevices.json
func ExampleResourceClient_ImportDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armiothub.NewResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ImportDevices(ctx,
		"myResourceGroup",
		"testHub",
		armiothub.ImportDevicesRequest{
			InputBlobContainerURI:  to.Ptr("testBlob"),
			OutputBlobContainerURI: to.Ptr("testBlob"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}