package security

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ConnectorGovernanceRulesClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type ConnectorGovernanceRulesClient struct {
	BaseClient
}

// NewConnectorGovernanceRulesClient creates an instance of the ConnectorGovernanceRulesClient client.
func NewConnectorGovernanceRulesClient(subscriptionID string) ConnectorGovernanceRulesClient {
	return NewConnectorGovernanceRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConnectorGovernanceRulesClientWithBaseURI creates an instance of the ConnectorGovernanceRulesClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewConnectorGovernanceRulesClientWithBaseURI(baseURI string, subscriptionID string) ConnectorGovernanceRulesClient {
	return ConnectorGovernanceRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or update a security GovernanceRule on the given security connector.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// securityConnectorName - the security connector name.
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
// governanceRule - governanceRule over a subscription scope
func (client ConnectorGovernanceRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, governanceRule GovernanceRule) (result GovernanceRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorGovernanceRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: governanceRule,
			Constraints: []validation.Constraint{{Target: "governanceRule.GovernanceRuleProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "governanceRule.GovernanceRuleProperties.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "governanceRule.GovernanceRuleProperties.RulePriority", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "governanceRule.GovernanceRuleProperties.RulePriority", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
							{Target: "governanceRule.GovernanceRuleProperties.RulePriority", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
						}},
					{Target: "governanceRule.GovernanceRuleProperties.SourceResourceType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "governanceRule.GovernanceRuleProperties.ConditionSets", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "governanceRule.GovernanceRuleProperties.OwnerSource", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("security.ConnectorGovernanceRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, securityConnectorName, ruleID, governanceRule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConnectorGovernanceRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, governanceRule GovernanceRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"ruleId":                autorest.Encode("path", ruleID),
		"securityConnectorName": autorest.Encode("path", securityConnectorName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}", pathParameters),
		autorest.WithJSON(governanceRule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorGovernanceRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConnectorGovernanceRulesClient) CreateOrUpdateResponder(resp *http.Response) (result GovernanceRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a GovernanceRule over a given scope
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// securityConnectorName - the security connector name.
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
func (client ConnectorGovernanceRulesClient) Delete(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorGovernanceRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.ConnectorGovernanceRulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, securityConnectorName, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConnectorGovernanceRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"ruleId":                autorest.Encode("path", ruleID),
		"securityConnectorName": autorest.Encode("path", securityConnectorName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorGovernanceRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConnectorGovernanceRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a specific governanceRule for the requested scope by ruleId
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// securityConnectorName - the security connector name.
// ruleID - the security GovernanceRule key - unique key for the standard GovernanceRule
func (client ConnectorGovernanceRulesClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string) (result GovernanceRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorGovernanceRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.ConnectorGovernanceRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, securityConnectorName, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ConnectorGovernanceRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConnectorGovernanceRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"ruleId":                autorest.Encode("path", ruleID),
		"securityConnectorName": autorest.Encode("path", securityConnectorName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorGovernanceRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConnectorGovernanceRulesClient) GetResponder(resp *http.Response) (result GovernanceRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}