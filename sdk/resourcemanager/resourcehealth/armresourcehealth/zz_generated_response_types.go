//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

// AvailabilityStatusesClientGetByResourceResponse contains the response from method AvailabilityStatusesClient.GetByResource.
type AvailabilityStatusesClientGetByResourceResponse struct {
	AvailabilityStatus
}

// AvailabilityStatusesClientListByResourceGroupResponse contains the response from method AvailabilityStatusesClient.ListByResourceGroup.
type AvailabilityStatusesClientListByResourceGroupResponse struct {
	AvailabilityStatusListResult
}

// AvailabilityStatusesClientListBySubscriptionIDResponse contains the response from method AvailabilityStatusesClient.ListBySubscriptionID.
type AvailabilityStatusesClientListBySubscriptionIDResponse struct {
	AvailabilityStatusListResult
}

// AvailabilityStatusesClientListResponse contains the response from method AvailabilityStatusesClient.List.
type AvailabilityStatusesClientListResponse struct {
	AvailabilityStatusListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}