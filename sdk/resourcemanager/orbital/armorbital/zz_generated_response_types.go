//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

// AvailableGroundStationsClientGetResponse contains the response from method AvailableGroundStationsClient.Get.
type AvailableGroundStationsClientGetResponse struct {
	AvailableGroundStation
}

// AvailableGroundStationsClientListByCapabilityResponse contains the response from method AvailableGroundStationsClient.ListByCapability.
type AvailableGroundStationsClientListByCapabilityResponse struct {
	AvailableGroundStationListResult
}

// ContactProfilesClientCreateOrUpdateResponse contains the response from method ContactProfilesClient.CreateOrUpdate.
type ContactProfilesClientCreateOrUpdateResponse struct {
	ContactProfile
}

// ContactProfilesClientDeleteResponse contains the response from method ContactProfilesClient.Delete.
type ContactProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ContactProfilesClientGetResponse contains the response from method ContactProfilesClient.Get.
type ContactProfilesClientGetResponse struct {
	ContactProfile
}

// ContactProfilesClientListBySubscriptionResponse contains the response from method ContactProfilesClient.ListBySubscription.
type ContactProfilesClientListBySubscriptionResponse struct {
	ContactProfileListResult
}

// ContactProfilesClientListResponse contains the response from method ContactProfilesClient.List.
type ContactProfilesClientListResponse struct {
	ContactProfileListResult
}

// ContactProfilesClientUpdateTagsResponse contains the response from method ContactProfilesClient.UpdateTags.
type ContactProfilesClientUpdateTagsResponse struct {
	ContactProfile
}

// ContactsClientCreateResponse contains the response from method ContactsClient.Create.
type ContactsClientCreateResponse struct {
	Contact
}

// ContactsClientDeleteResponse contains the response from method ContactsClient.Delete.
type ContactsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContactsClientGetResponse contains the response from method ContactsClient.Get.
type ContactsClientGetResponse struct {
	Contact
}

// ContactsClientListResponse contains the response from method ContactsClient.List.
type ContactsClientListResponse struct {
	ContactListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// OperationsResultsClientGetResponse contains the response from method OperationsResultsClient.Get.
type OperationsResultsClientGetResponse struct {
	OperationResult
}

// SpacecraftsClientCreateOrUpdateResponse contains the response from method SpacecraftsClient.CreateOrUpdate.
type SpacecraftsClientCreateOrUpdateResponse struct {
	Spacecraft
}

// SpacecraftsClientDeleteResponse contains the response from method SpacecraftsClient.Delete.
type SpacecraftsClientDeleteResponse struct {
	// placeholder for future response values
}

// SpacecraftsClientGetResponse contains the response from method SpacecraftsClient.Get.
type SpacecraftsClientGetResponse struct {
	Spacecraft
}

// SpacecraftsClientListAvailableContactsResponse contains the response from method SpacecraftsClient.ListAvailableContacts.
type SpacecraftsClientListAvailableContactsResponse struct {
	AvailableContactsListResult
}

// SpacecraftsClientListBySubscriptionResponse contains the response from method SpacecraftsClient.ListBySubscription.
type SpacecraftsClientListBySubscriptionResponse struct {
	SpacecraftListResult
}

// SpacecraftsClientListResponse contains the response from method SpacecraftsClient.List.
type SpacecraftsClientListResponse struct {
	SpacecraftListResult
}

// SpacecraftsClientUpdateTagsResponse contains the response from method SpacecraftsClient.UpdateTags.
type SpacecraftsClientUpdateTagsResponse struct {
	Spacecraft
}