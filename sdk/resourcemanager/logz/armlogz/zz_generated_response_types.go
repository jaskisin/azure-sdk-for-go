//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

// MonitorClientListVMHostUpdateResponse contains the response from method MonitorClient.ListVMHostUpdate.
type MonitorClientListVMHostUpdateResponse struct {
	VMResourcesListResponse
}

// MonitorClientListVMHostsResponse contains the response from method MonitorClient.ListVMHosts.
type MonitorClientListVMHostsResponse struct {
	VMResourcesListResponse
}

// MonitorClientVMHostPayloadResponse contains the response from method MonitorClient.VMHostPayload.
type MonitorClientVMHostPayloadResponse struct {
	VMExtensionPayload
}

// MonitorsClientCreateResponse contains the response from method MonitorsClient.Create.
type MonitorsClientCreateResponse struct {
	MonitorResource
}

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.Delete.
type MonitorsClientDeleteResponse struct {
	// placeholder for future response values
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	MonitorResource
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.ListByResourceGroup.
type MonitorsClientListByResourceGroupResponse struct {
	MonitorResourceListResponse
}

// MonitorsClientListBySubscriptionResponse contains the response from method MonitorsClient.ListBySubscription.
type MonitorsClientListBySubscriptionResponse struct {
	MonitorResourceListResponse
}

// MonitorsClientListMonitoredResourcesResponse contains the response from method MonitorsClient.ListMonitoredResources.
type MonitorsClientListMonitoredResourcesResponse struct {
	MonitoredResourceListResponse
}

// MonitorsClientListUserRolesResponse contains the response from method MonitorsClient.ListUserRoles.
type MonitorsClientListUserRolesResponse struct {
	UserRoleListResponse
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.Update.
type MonitorsClientUpdateResponse struct {
	MonitorResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// SingleSignOnClientCreateOrUpdateResponse contains the response from method SingleSignOnClient.CreateOrUpdate.
type SingleSignOnClientCreateOrUpdateResponse struct {
	SingleSignOnResource
}

// SingleSignOnClientGetResponse contains the response from method SingleSignOnClient.Get.
type SingleSignOnClientGetResponse struct {
	SingleSignOnResource
}

// SingleSignOnClientListResponse contains the response from method SingleSignOnClient.List.
type SingleSignOnClientListResponse struct {
	SingleSignOnResourceListResponse
}

// SubAccountClientCreateResponse contains the response from method SubAccountClient.Create.
type SubAccountClientCreateResponse struct {
	MonitorResource
}

// SubAccountClientDeleteResponse contains the response from method SubAccountClient.Delete.
type SubAccountClientDeleteResponse struct {
	// placeholder for future response values
}

// SubAccountClientGetResponse contains the response from method SubAccountClient.Get.
type SubAccountClientGetResponse struct {
	MonitorResource
}

// SubAccountClientListMonitoredResourcesResponse contains the response from method SubAccountClient.ListMonitoredResources.
type SubAccountClientListMonitoredResourcesResponse struct {
	MonitoredResourceListResponse
}

// SubAccountClientListResponse contains the response from method SubAccountClient.List.
type SubAccountClientListResponse struct {
	MonitorResourceListResponse
}

// SubAccountClientListVMHostUpdateResponse contains the response from method SubAccountClient.ListVMHostUpdate.
type SubAccountClientListVMHostUpdateResponse struct {
	VMResourcesListResponse
}

// SubAccountClientListVMHostsResponse contains the response from method SubAccountClient.ListVMHosts.
type SubAccountClientListVMHostsResponse struct {
	VMResourcesListResponse
}

// SubAccountClientUpdateResponse contains the response from method SubAccountClient.Update.
type SubAccountClientUpdateResponse struct {
	MonitorResource
}

// SubAccountClientVMHostPayloadResponse contains the response from method SubAccountClient.VMHostPayload.
type SubAccountClientVMHostPayloadResponse struct {
	VMExtensionPayload
}

// SubAccountTagRulesClientCreateOrUpdateResponse contains the response from method SubAccountTagRulesClient.CreateOrUpdate.
type SubAccountTagRulesClientCreateOrUpdateResponse struct {
	MonitoringTagRules
}

// SubAccountTagRulesClientDeleteResponse contains the response from method SubAccountTagRulesClient.Delete.
type SubAccountTagRulesClientDeleteResponse struct {
	// Location contains the information returned from the location header response.
	Location *string
}

// SubAccountTagRulesClientGetResponse contains the response from method SubAccountTagRulesClient.Get.
type SubAccountTagRulesClientGetResponse struct {
	MonitoringTagRules
}

// SubAccountTagRulesClientListResponse contains the response from method SubAccountTagRulesClient.List.
type SubAccountTagRulesClientListResponse struct {
	MonitoringTagRulesListResponse
}

// TagRulesClientCreateOrUpdateResponse contains the response from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResponse struct {
	MonitoringTagRules
}

// TagRulesClientDeleteResponse contains the response from method TagRulesClient.Delete.
type TagRulesClientDeleteResponse struct {
	// Location contains the information returned from the location header response.
	Location *string
}

// TagRulesClientGetResponse contains the response from method TagRulesClient.Get.
type TagRulesClientGetResponse struct {
	MonitoringTagRules
}

// TagRulesClientListResponse contains the response from method TagRulesClient.List.
type TagRulesClientListResponse struct {
	MonitoringTagRulesListResponse
}