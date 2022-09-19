//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdnsresolver

// DNSForwardingRulesetsClientCreateOrUpdateResponse contains the response from method DNSForwardingRulesetsClient.CreateOrUpdate.
type DNSForwardingRulesetsClientCreateOrUpdateResponse struct {
	DNSForwardingRuleset
}

// DNSForwardingRulesetsClientDeleteResponse contains the response from method DNSForwardingRulesetsClient.Delete.
type DNSForwardingRulesetsClientDeleteResponse struct {
	// placeholder for future response values
}

// DNSForwardingRulesetsClientGetResponse contains the response from method DNSForwardingRulesetsClient.Get.
type DNSForwardingRulesetsClientGetResponse struct {
	DNSForwardingRuleset
}

// DNSForwardingRulesetsClientListByResourceGroupResponse contains the response from method DNSForwardingRulesetsClient.ListByResourceGroup.
type DNSForwardingRulesetsClientListByResourceGroupResponse struct {
	DNSForwardingRulesetListResult
}

// DNSForwardingRulesetsClientListByVirtualNetworkResponse contains the response from method DNSForwardingRulesetsClient.ListByVirtualNetwork.
type DNSForwardingRulesetsClientListByVirtualNetworkResponse struct {
	VirtualNetworkDNSForwardingRulesetListResult
}

// DNSForwardingRulesetsClientListResponse contains the response from method DNSForwardingRulesetsClient.List.
type DNSForwardingRulesetsClientListResponse struct {
	DNSForwardingRulesetListResult
}

// DNSForwardingRulesetsClientUpdateResponse contains the response from method DNSForwardingRulesetsClient.Update.
type DNSForwardingRulesetsClientUpdateResponse struct {
	DNSForwardingRuleset
}

// DNSResolversClientCreateOrUpdateResponse contains the response from method DNSResolversClient.CreateOrUpdate.
type DNSResolversClientCreateOrUpdateResponse struct {
	DNSResolver
}

// DNSResolversClientDeleteResponse contains the response from method DNSResolversClient.Delete.
type DNSResolversClientDeleteResponse struct {
	// placeholder for future response values
}

// DNSResolversClientGetResponse contains the response from method DNSResolversClient.Get.
type DNSResolversClientGetResponse struct {
	DNSResolver
}

// DNSResolversClientListByResourceGroupResponse contains the response from method DNSResolversClient.ListByResourceGroup.
type DNSResolversClientListByResourceGroupResponse struct {
	ListResult
}

// DNSResolversClientListByVirtualNetworkResponse contains the response from method DNSResolversClient.ListByVirtualNetwork.
type DNSResolversClientListByVirtualNetworkResponse struct {
	SubResourceListResult
}

// DNSResolversClientListResponse contains the response from method DNSResolversClient.List.
type DNSResolversClientListResponse struct {
	ListResult
}

// DNSResolversClientUpdateResponse contains the response from method DNSResolversClient.Update.
type DNSResolversClientUpdateResponse struct {
	DNSResolver
}

// ForwardingRulesClientCreateOrUpdateResponse contains the response from method ForwardingRulesClient.CreateOrUpdate.
type ForwardingRulesClientCreateOrUpdateResponse struct {
	ForwardingRule
}

// ForwardingRulesClientDeleteResponse contains the response from method ForwardingRulesClient.Delete.
type ForwardingRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// ForwardingRulesClientGetResponse contains the response from method ForwardingRulesClient.Get.
type ForwardingRulesClientGetResponse struct {
	ForwardingRule
}

// ForwardingRulesClientListResponse contains the response from method ForwardingRulesClient.List.
type ForwardingRulesClientListResponse struct {
	ForwardingRuleListResult
}

// ForwardingRulesClientUpdateResponse contains the response from method ForwardingRulesClient.Update.
type ForwardingRulesClientUpdateResponse struct {
	ForwardingRule
}

// InboundEndpointsClientCreateOrUpdateResponse contains the response from method InboundEndpointsClient.CreateOrUpdate.
type InboundEndpointsClientCreateOrUpdateResponse struct {
	InboundEndpoint
}

// InboundEndpointsClientDeleteResponse contains the response from method InboundEndpointsClient.Delete.
type InboundEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// InboundEndpointsClientGetResponse contains the response from method InboundEndpointsClient.Get.
type InboundEndpointsClientGetResponse struct {
	InboundEndpoint
}

// InboundEndpointsClientListResponse contains the response from method InboundEndpointsClient.List.
type InboundEndpointsClientListResponse struct {
	InboundEndpointListResult
}

// InboundEndpointsClientUpdateResponse contains the response from method InboundEndpointsClient.Update.
type InboundEndpointsClientUpdateResponse struct {
	InboundEndpoint
}

// OutboundEndpointsClientCreateOrUpdateResponse contains the response from method OutboundEndpointsClient.CreateOrUpdate.
type OutboundEndpointsClientCreateOrUpdateResponse struct {
	OutboundEndpoint
}

// OutboundEndpointsClientDeleteResponse contains the response from method OutboundEndpointsClient.Delete.
type OutboundEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// OutboundEndpointsClientGetResponse contains the response from method OutboundEndpointsClient.Get.
type OutboundEndpointsClientGetResponse struct {
	OutboundEndpoint
}

// OutboundEndpointsClientListResponse contains the response from method OutboundEndpointsClient.List.
type OutboundEndpointsClientListResponse struct {
	OutboundEndpointListResult
}

// OutboundEndpointsClientUpdateResponse contains the response from method OutboundEndpointsClient.Update.
type OutboundEndpointsClientUpdateResponse struct {
	OutboundEndpoint
}

// VirtualNetworkLinksClientCreateOrUpdateResponse contains the response from method VirtualNetworkLinksClient.CreateOrUpdate.
type VirtualNetworkLinksClientCreateOrUpdateResponse struct {
	VirtualNetworkLink
}

// VirtualNetworkLinksClientDeleteResponse contains the response from method VirtualNetworkLinksClient.Delete.
type VirtualNetworkLinksClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworkLinksClientGetResponse contains the response from method VirtualNetworkLinksClient.Get.
type VirtualNetworkLinksClientGetResponse struct {
	VirtualNetworkLink
}

// VirtualNetworkLinksClientListResponse contains the response from method VirtualNetworkLinksClient.List.
type VirtualNetworkLinksClientListResponse struct {
	VirtualNetworkLinkListResult
}

// VirtualNetworkLinksClientUpdateResponse contains the response from method VirtualNetworkLinksClient.Update.
type VirtualNetworkLinksClientUpdateResponse struct {
	VirtualNetworkLink
}