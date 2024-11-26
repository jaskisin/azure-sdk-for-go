//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/KubeEnvironments_ListBySubscription.json
func ExampleKubeEnvironmentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKubeEnvironmentsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.KubeEnvironmentCollection = armappservice.KubeEnvironmentCollection{
		// 	Value: []*armappservice.KubeEnvironment{
		// 		{
		// 			Name: to.Ptr("jlaw-demo1"),
		// 			Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/jlaw-demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armappservice.KubeEnvironmentProperties{
		// 				DefaultDomain: to.Ptr("jlaw-demo1.k4apps.io"),
		// 				InternalLoadBalancerEnabled: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("20.42.33.145"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("demo1"),
		// 			Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/DemoRG/providers/Microsoft.Web/kubeEnvironments/demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armappservice.KubeEnvironmentProperties{
		// 				DefaultDomain: to.Ptr("demo1.k4apps.io"),
		// 				InternalLoadBalancerEnabled: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("52.142.21.61"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/KubeEnvironments_ListByResourceGroup.json
func ExampleKubeEnvironmentsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKubeEnvironmentsClient().NewListByResourceGroupPager("examplerg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.KubeEnvironmentCollection = armappservice.KubeEnvironmentCollection{
		// 	Value: []*armappservice.KubeEnvironment{
		// 		{
		// 			Name: to.Ptr("jlaw-demo1"),
		// 			Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/jlaw-demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappservice.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armappservice.KubeEnvironmentProperties{
		// 				DefaultDomain: to.Ptr("jlaw-demo1.k4apps.io"),
		// 				InternalLoadBalancerEnabled: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("20.42.33.145"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("demo1"),
		// 			Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/demo1"),
		// 			Location: to.Ptr("North Central US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armappservice.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armappservice.KubeEnvironmentProperties{
		// 				DefaultDomain: to.Ptr("demo1.k4apps.io"),
		// 				InternalLoadBalancerEnabled: to.Ptr(false),
		// 				ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
		// 				StaticIP: to.Ptr("52.142.21.61"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/KubeEnvironments_Get.json
func ExampleKubeEnvironmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKubeEnvironmentsClient().Get(ctx, "examplerg", "jlaw-demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubeEnvironment = armappservice.KubeEnvironment{
	// 	Name: to.Ptr("jlaw-demo1"),
	// 	Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/jlaw-demo1"),
	// 	Location: to.Ptr("North Central US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	ExtendedLocation: &armappservice.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armappservice.KubeEnvironmentProperties{
	// 		AksResourceID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.ContainerService/managedClusters/jlaw-demo1"),
	// 		DefaultDomain: to.Ptr("jlaw-demo1.k4apps.io"),
	// 		InternalLoadBalancerEnabled: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("20.42.33.145"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/KubeEnvironments_CreateOrUpdate.json
func ExampleKubeEnvironmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKubeEnvironmentsClient().BeginCreateOrUpdate(ctx, "examplerg", "testkubeenv", armappservice.KubeEnvironment{
		Location: to.Ptr("East US"),
		Properties: &armappservice.KubeEnvironmentProperties{
			StaticIP: to.Ptr("1.2.3.4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubeEnvironment = armappservice.KubeEnvironment{
	// 	Name: to.Ptr("testkubeenv"),
	// 	Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/testkubeenv"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	ExtendedLocation: &armappservice.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armappservice.KubeEnvironmentProperties{
	// 		AksResourceID: to.Ptr("test"),
	// 		DefaultDomain: to.Ptr("testkubeenv.k4apps.io"),
	// 		InternalLoadBalancerEnabled: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("1.2.3.4"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/KubeEnvironments_Delete.json
func ExampleKubeEnvironmentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKubeEnvironmentsClient().BeginDelete(ctx, "examplerg", "examplekenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b5d78da207e9c5d8f82e95224039867271f47cdf/specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/KubeEnvironments_Update.json
func ExampleKubeEnvironmentsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKubeEnvironmentsClient().Update(ctx, "examplerg", "testkubeenv", armappservice.KubeEnvironmentPatchResource{
		Properties: &armappservice.KubeEnvironmentPatchResourceProperties{
			StaticIP: to.Ptr("1.2.3.4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubeEnvironment = armappservice.KubeEnvironment{
	// 	Name: to.Ptr("testkubeenv"),
	// 	Type: to.Ptr("Microsoft.Web/kubeEnvironments"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.Web/kubeEnvironments/testkubeenv"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	ExtendedLocation: &armappservice.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armappservice.KubeEnvironmentProperties{
	// 		AksResourceID: to.Ptr("test"),
	// 		DefaultDomain: to.Ptr("testkubeenv.k4apps.io"),
	// 		InternalLoadBalancerEnabled: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armappservice.KubeEnvironmentProvisioningStateSucceeded),
	// 		StaticIP: to.Ptr("1.2.3.4"),
	// 	},
	// }
}
