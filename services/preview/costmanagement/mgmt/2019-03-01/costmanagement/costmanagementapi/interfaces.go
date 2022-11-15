package costmanagementapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/costmanagement/mgmt/2019-03-01/costmanagement"
	"github.com/Azure/go-autorest/autorest"
)

// DimensionsClientAPI contains the set of methods on the DimensionsClient type.
type DimensionsClientAPI interface {
	ListByBillingAccount(ctx context.Context, billingAccountID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
	ListByDepartment(ctx context.Context, billingAccountID string, departmentID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
	ListByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
	ListByManagementGroup(ctx context.Context, managementGroupID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
	ListBySubscription(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ DimensionsClientAPI = (*costmanagement.DimensionsClient)(nil)

// QueryClientAPI contains the set of methods on the QueryClient type.
type QueryClientAPI interface {
	UsageByBillingAccount(ctx context.Context, billingAccountID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByDepartment(ctx context.Context, billingAccountID string, departmentID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByManagementGroup(ctx context.Context, managementGroupID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByResourceGroup(ctx context.Context, resourceGroupName string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageBySubscription(ctx context.Context, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
}

var _ QueryClientAPI = (*costmanagement.QueryClient)(nil)

// ForecastClientAPI contains the set of methods on the ForecastClient type.
type ForecastClientAPI interface {
	UsageByBillingAccount(ctx context.Context, billingAccountID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByDepartment(ctx context.Context, billingAccountID string, departmentID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByExternalBillingAccount(ctx context.Context, externalBillingAccountName string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByManagementGroup(ctx context.Context, managementGroupID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageByResourceGroup(ctx context.Context, resourceGroupName string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	UsageBySubscription(ctx context.Context, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
}

var _ ForecastClientAPI = (*costmanagement.ForecastClient)(nil)

// CloudConnectorClientAPI contains the set of methods on the CloudConnectorClient type.
type CloudConnectorClientAPI interface {
	CreateOrUpdate(ctx context.Context, connectorName string, connector costmanagement.ConnectorDefinition) (result costmanagement.ConnectorDefinition, err error)
	Delete(ctx context.Context, connectorName string) (result autorest.Response, err error)
	Get(ctx context.Context, connectorName string, expand string) (result costmanagement.ConnectorDefinition, err error)
	List(ctx context.Context) (result costmanagement.ConnectorDefinitionListResult, err error)
	Update(ctx context.Context, connectorName string, connector costmanagement.ConnectorDefinition) (result costmanagement.ConnectorDefinition, err error)
}

var _ CloudConnectorClientAPI = (*costmanagement.CloudConnectorClient)(nil)

// ConnectorClientAPI contains the set of methods on the ConnectorClient type.
type ConnectorClientAPI interface {
	CheckEligibility(ctx context.Context, connectorCredentials costmanagement.CheckEligibilityDefinition) (result costmanagement.ConnectorDefinition, err error)
}

var _ ConnectorClientAPI = (*costmanagement.ConnectorClient)(nil)

// ExternalBillingAccountClientAPI contains the set of methods on the ExternalBillingAccountClient type.
type ExternalBillingAccountClientAPI interface {
	Get(ctx context.Context, externalBillingAccountName string, expand string) (result costmanagement.ExternalBillingAccountDefinition, err error)
	List(ctx context.Context) (result costmanagement.ExternalBillingAccountDefinitionListResult, err error)
}

var _ ExternalBillingAccountClientAPI = (*costmanagement.ExternalBillingAccountClient)(nil)

// ExternalSubscriptionClientAPI contains the set of methods on the ExternalSubscriptionClient type.
type ExternalSubscriptionClientAPI interface {
	Get(ctx context.Context, externalSubscriptionName string, expand string) (result costmanagement.ExternalSubscriptionDefinition, err error)
	List(ctx context.Context) (result costmanagement.ExternalSubscriptionDefinitionListResult, err error)
	ListByExternalBillingAccount(ctx context.Context, externalBillingAccountName string) (result costmanagement.ExternalSubscriptionDefinitionListResult, err error)
	ListByManagementGroup(ctx context.Context, managementGroupID string, recurse *bool) (result costmanagement.ExternalSubscriptionDefinitionListResult, err error)
	UpdateManagementGroup(ctx context.Context, managementGroupID string, externalSubscriptionName string) (result autorest.Response, err error)
}

var _ ExternalSubscriptionClientAPI = (*costmanagement.ExternalSubscriptionClient)(nil)

// ShowbackRulesClientAPI contains the set of methods on the ShowbackRulesClient type.
type ShowbackRulesClientAPI interface {
	List(ctx context.Context, billingAccountID string) (result costmanagement.ShowbackRuleListResult, err error)
}

var _ ShowbackRulesClientAPI = (*costmanagement.ShowbackRulesClient)(nil)

// ShowbackRuleClientAPI contains the set of methods on the ShowbackRuleClient type.
type ShowbackRuleClientAPI interface {
	CreateUpdateRule(ctx context.Context, billingAccountID string, ruleName string, showbackRule costmanagement.ShowbackRule) (result costmanagement.ShowbackRule, err error)
	GetBillingAccountID(ctx context.Context, billingAccountID string, ruleName string) (result costmanagement.ShowbackRule, err error)
}

var _ ShowbackRuleClientAPI = (*costmanagement.ShowbackRuleClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result costmanagement.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result costmanagement.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*costmanagement.OperationsClient)(nil)