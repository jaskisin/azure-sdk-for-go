//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoep

import "time"

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string `json:"name,omitempty"`

	// The resource type.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string `json:"message,omitempty"`

	// Indicates if the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason `json:"reason,omitempty"`
}

// DataPartitionAddOrRemoveRequest - Defines the partition add/ delete action properties.
type DataPartitionAddOrRemoveRequest struct {
	// Name of the data partition
	Name *string `json:"name,omitempty"`
}

// DataPartitionNames - The list of Energy services resource's Data Partition Names.
type DataPartitionNames struct {
	Name *string `json:"name,omitempty"`
}

// DataPartitionProperties - Defines the properties of an individual data partition.
type DataPartitionProperties struct {
	// Name of the data partition
	Name *string `json:"name,omitempty"`

	// Name of the data partition
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// DataPartitionsList - List of data partitions
type DataPartitionsList struct {
	DataPartitionNames []*DataPartitionNames `json:"dataPartitionNames,omitempty"`
}

// DataPartitionsListResult - List of data partitions.
type DataPartitionsListResult struct {
	// List of data partitions along with their properties in a given OEP resource.
	Value []*DataPartitionProperties `json:"value,omitempty"`
}

// EnergyResourceUpdate - The resource model definition used for updating a tracked ARM resource.
type EnergyResourceUpdate struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

type EnergyService struct {
	// REQUIRED; Geo-location where the resource lives.
	Location   *string                  `json:"location,omitempty"`
	Properties *EnergyServiceProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// EnergyServiceList - The list of oep resources.
type EnergyServiceList struct {
	// The link used to get the next page of oep resources list.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of oep resources.
	Value []*EnergyService `json:"value,omitempty"`
}

type EnergyServiceProperties struct {
	AuthAppID          *string               `json:"authAppId,omitempty"`
	DataPartitionNames []*DataPartitionNames `json:"dataPartitionNames,omitempty"`

	// READ-ONLY
	DNSName *string `json:"dnsName,omitempty" azure:"ro"`

	// READ-ONLY
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// EnergyServicesClientBeginAddPartitionOptions contains the optional parameters for the EnergyServicesClient.BeginAddPartition
// method.
type EnergyServicesClientBeginAddPartitionOptions struct {
	// add partition action payload
	Body *DataPartitionAddOrRemoveRequest
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EnergyServicesClientBeginCreateOptions contains the optional parameters for the EnergyServicesClient.BeginCreate method.
type EnergyServicesClientBeginCreateOptions struct {
	// Request body.
	Body *EnergyService
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EnergyServicesClientBeginDeleteOptions contains the optional parameters for the EnergyServicesClient.BeginDelete method.
type EnergyServicesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EnergyServicesClientBeginRemovePartitionOptions contains the optional parameters for the EnergyServicesClient.BeginRemovePartition
// method.
type EnergyServicesClientBeginRemovePartitionOptions struct {
	// remove partition action payload
	Body *DataPartitionAddOrRemoveRequest
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EnergyServicesClientGetOptions contains the optional parameters for the EnergyServicesClient.Get method.
type EnergyServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientListByResourceGroupOptions contains the optional parameters for the EnergyServicesClient.ListByResourceGroup
// method.
type EnergyServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientListBySubscriptionOptions contains the optional parameters for the EnergyServicesClient.ListBySubscription
// method.
type EnergyServicesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientListPartitionsOptions contains the optional parameters for the EnergyServicesClient.ListPartitions
// method.
type EnergyServicesClientListPartitionsOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientUpdateOptions contains the optional parameters for the EnergyServicesClient.Update method.
type EnergyServicesClientUpdateOptions struct {
	Body *EnergyResourceUpdate
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
type LocationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}