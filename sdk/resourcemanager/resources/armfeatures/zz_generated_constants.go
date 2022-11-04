//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

const (
	moduleName    = "armfeatures"
	moduleVersion = "v1.0.0"
)

// SubscriptionFeatureRegistrationApprovalType - The feature approval type.
type SubscriptionFeatureRegistrationApprovalType string

const (
	SubscriptionFeatureRegistrationApprovalTypeApprovalRequired SubscriptionFeatureRegistrationApprovalType = "ApprovalRequired"
	SubscriptionFeatureRegistrationApprovalTypeAutoApproval     SubscriptionFeatureRegistrationApprovalType = "AutoApproval"
	SubscriptionFeatureRegistrationApprovalTypeNotSpecified     SubscriptionFeatureRegistrationApprovalType = "NotSpecified"
)

// PossibleSubscriptionFeatureRegistrationApprovalTypeValues returns the possible values for the SubscriptionFeatureRegistrationApprovalType const type.
func PossibleSubscriptionFeatureRegistrationApprovalTypeValues() []SubscriptionFeatureRegistrationApprovalType {
	return []SubscriptionFeatureRegistrationApprovalType{
		SubscriptionFeatureRegistrationApprovalTypeApprovalRequired,
		SubscriptionFeatureRegistrationApprovalTypeAutoApproval,
		SubscriptionFeatureRegistrationApprovalTypeNotSpecified,
	}
}

// SubscriptionFeatureRegistrationState - The state.
type SubscriptionFeatureRegistrationState string

const (
	SubscriptionFeatureRegistrationStateNotRegistered SubscriptionFeatureRegistrationState = "NotRegistered"
	SubscriptionFeatureRegistrationStateNotSpecified  SubscriptionFeatureRegistrationState = "NotSpecified"
	SubscriptionFeatureRegistrationStatePending       SubscriptionFeatureRegistrationState = "Pending"
	SubscriptionFeatureRegistrationStateRegistered    SubscriptionFeatureRegistrationState = "Registered"
	SubscriptionFeatureRegistrationStateRegistering   SubscriptionFeatureRegistrationState = "Registering"
	SubscriptionFeatureRegistrationStateUnregistered  SubscriptionFeatureRegistrationState = "Unregistered"
	SubscriptionFeatureRegistrationStateUnregistering SubscriptionFeatureRegistrationState = "Unregistering"
)

// PossibleSubscriptionFeatureRegistrationStateValues returns the possible values for the SubscriptionFeatureRegistrationState const type.
func PossibleSubscriptionFeatureRegistrationStateValues() []SubscriptionFeatureRegistrationState {
	return []SubscriptionFeatureRegistrationState{
		SubscriptionFeatureRegistrationStateNotRegistered,
		SubscriptionFeatureRegistrationStateNotSpecified,
		SubscriptionFeatureRegistrationStatePending,
		SubscriptionFeatureRegistrationStateRegistered,
		SubscriptionFeatureRegistrationStateRegistering,
		SubscriptionFeatureRegistrationStateUnregistered,
		SubscriptionFeatureRegistrationStateUnregistering,
	}
}