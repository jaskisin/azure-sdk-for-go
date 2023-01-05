//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_ListByFactory.json
func ExampleDatasetsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDatasetsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Create.json
func ExampleDatasetsClient_CreateOrUpdate_datasetsCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDatasetsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", armdatafactory.DatasetResource{
		Properties: &armdatafactory.AzureBlobDataset{
			Type: to.Ptr("AzureBlob"),
			LinkedServiceName: &armdatafactory.LinkedServiceReference{
				Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
				ReferenceName: to.Ptr("exampleLinkedService"),
			},
			Parameters: map[string]*armdatafactory.ParameterSpecification{
				"MyFileName": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
				"MyFolderPath": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
			},
			TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
				Format: &armdatafactory.TextFormat{
					Type: to.Ptr("TextFormat"),
				},
				FileName: map[string]interface{}{
					"type":  "Expression",
					"value": "@dataset().MyFileName",
				},
				FolderPath: map[string]interface{}{
					"type":  "Expression",
					"value": "@dataset().MyFolderPath",
				},
			},
		},
	}, &armdatafactory.DatasetsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Update.json
func ExampleDatasetsClient_CreateOrUpdate_datasetsUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDatasetsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", armdatafactory.DatasetResource{
		Properties: &armdatafactory.AzureBlobDataset{
			Type:        to.Ptr("AzureBlob"),
			Description: to.Ptr("Example description"),
			LinkedServiceName: &armdatafactory.LinkedServiceReference{
				Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
				ReferenceName: to.Ptr("exampleLinkedService"),
			},
			Parameters: map[string]*armdatafactory.ParameterSpecification{
				"MyFileName": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
				"MyFolderPath": {
					Type: to.Ptr(armdatafactory.ParameterTypeString),
				},
			},
			TypeProperties: &armdatafactory.AzureBlobDatasetTypeProperties{
				Format: &armdatafactory.TextFormat{
					Type: to.Ptr("TextFormat"),
				},
				FileName: map[string]interface{}{
					"type":  "Expression",
					"value": "@dataset().MyFileName",
				},
				FolderPath: map[string]interface{}{
					"type":  "Expression",
					"value": "@dataset().MyFolderPath",
				},
			},
		},
	}, &armdatafactory.DatasetsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Get.json
func ExampleDatasetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDatasetsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", &armdatafactory.DatasetsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Datasets_Delete.json
func ExampleDatasetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDatasetsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataset", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}