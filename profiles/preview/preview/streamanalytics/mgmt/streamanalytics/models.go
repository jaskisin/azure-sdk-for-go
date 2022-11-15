//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package streamanalytics

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/streamanalytics/mgmt/2020-03-01-preview/streamanalytics"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AuthenticationMode = original.AuthenticationMode

const (
	ConnectionString AuthenticationMode = original.ConnectionString
	Msi              AuthenticationMode = original.Msi
	UserToken        AuthenticationMode = original.UserToken
)

type BindingType = original.BindingType

const (
	BindingTypeFunctionRetrieveDefaultDefinitionParameters BindingType = original.BindingTypeFunctionRetrieveDefaultDefinitionParameters
	BindingTypeMicrosoftMachineLearningServices            BindingType = original.BindingTypeMicrosoftMachineLearningServices
	BindingTypeMicrosoftMachineLearningWebService          BindingType = original.BindingTypeMicrosoftMachineLearningWebService
	BindingTypeMicrosoftStreamAnalyticsCLRUdf              BindingType = original.BindingTypeMicrosoftStreamAnalyticsCLRUdf
	BindingTypeMicrosoftStreamAnalyticsJavascriptUdf       BindingType = original.BindingTypeMicrosoftStreamAnalyticsJavascriptUdf
)

type ClusterProvisioningState = original.ClusterProvisioningState

const (
	Canceled   ClusterProvisioningState = original.Canceled
	Failed     ClusterProvisioningState = original.Failed
	InProgress ClusterProvisioningState = original.InProgress
	Succeeded  ClusterProvisioningState = original.Succeeded
)

type ClusterSkuName = original.ClusterSkuName

const (
	Default ClusterSkuName = original.Default
)

type CompatibilityLevel = original.CompatibilityLevel

const (
	OneFullStopZero CompatibilityLevel = original.OneFullStopZero
)

type ContentStoragePolicy = original.ContentStoragePolicy

const (
	ContentStoragePolicyJobStorageAccount ContentStoragePolicy = original.ContentStoragePolicyJobStorageAccount
	ContentStoragePolicySystemAccount     ContentStoragePolicy = original.ContentStoragePolicySystemAccount
)

type Encoding = original.Encoding

const (
	UTF8 Encoding = original.UTF8
)

type EventSerializationType = original.EventSerializationType

const (
	Avro      EventSerializationType = original.Avro
	Csv       EventSerializationType = original.Csv
	CustomClr EventSerializationType = original.CustomClr
	JSON      EventSerializationType = original.JSON
	Parquet   EventSerializationType = original.Parquet
)

type EventsOutOfOrderPolicy = original.EventsOutOfOrderPolicy

const (
	Adjust EventsOutOfOrderPolicy = original.Adjust
	Drop   EventsOutOfOrderPolicy = original.Drop
)

type JSONOutputSerializationFormat = original.JSONOutputSerializationFormat

const (
	Array         JSONOutputSerializationFormat = original.Array
	LineSeparated JSONOutputSerializationFormat = original.LineSeparated
)

type JobState = original.JobState

const (
	JobStateCreated    JobState = original.JobStateCreated
	JobStateDegraded   JobState = original.JobStateDegraded
	JobStateDeleting   JobState = original.JobStateDeleting
	JobStateFailed     JobState = original.JobStateFailed
	JobStateRestarting JobState = original.JobStateRestarting
	JobStateRunning    JobState = original.JobStateRunning
	JobStateScaling    JobState = original.JobStateScaling
	JobStateStarting   JobState = original.JobStateStarting
	JobStateStopped    JobState = original.JobStateStopped
	JobStateStopping   JobState = original.JobStateStopping
)

type JobType = original.JobType

const (
	Cloud JobType = original.Cloud
	Edge  JobType = original.Edge
)

type OutputErrorPolicy = original.OutputErrorPolicy

const (
	OutputErrorPolicyDrop OutputErrorPolicy = original.OutputErrorPolicyDrop
	OutputErrorPolicyStop OutputErrorPolicy = original.OutputErrorPolicyStop
)

type OutputStartMode = original.OutputStartMode

const (
	CustomTime          OutputStartMode = original.CustomTime
	JobStartTime        OutputStartMode = original.JobStartTime
	LastOutputEventTime OutputStartMode = original.LastOutputEventTime
)

type StreamingJobSkuName = original.StreamingJobSkuName

const (
	Standard StreamingJobSkuName = original.Standard
)

type Type = original.Type

const (
	TypeFunctionBinding                       Type = original.TypeFunctionBinding
	TypeMicrosoftMachineLearningServices      Type = original.TypeMicrosoftMachineLearningServices
	TypeMicrosoftMachineLearningWebService    Type = original.TypeMicrosoftMachineLearningWebService
	TypeMicrosoftStreamAnalyticsCLRUdf        Type = original.TypeMicrosoftStreamAnalyticsCLRUdf
	TypeMicrosoftStreamAnalyticsJavascriptUdf Type = original.TypeMicrosoftStreamAnalyticsJavascriptUdf
)

type TypeBasicFunctionProperties = original.TypeBasicFunctionProperties

const (
	TypeAggregate          TypeBasicFunctionProperties = original.TypeAggregate
	TypeFunctionProperties TypeBasicFunctionProperties = original.TypeFunctionProperties
	TypeScalar             TypeBasicFunctionProperties = original.TypeScalar
)

type TypeBasicInputProperties = original.TypeBasicInputProperties

const (
	TypeInputProperties TypeBasicInputProperties = original.TypeInputProperties
	TypeReference       TypeBasicInputProperties = original.TypeReference
	TypeStream          TypeBasicInputProperties = original.TypeStream
)

type TypeBasicOutputDataSource = original.TypeBasicOutputDataSource

const (
	TypeMicrosoftAzureFunction          TypeBasicOutputDataSource = original.TypeMicrosoftAzureFunction
	TypeMicrosoftDataLakeAccounts       TypeBasicOutputDataSource = original.TypeMicrosoftDataLakeAccounts
	TypeMicrosoftEventHubEventHub       TypeBasicOutputDataSource = original.TypeMicrosoftEventHubEventHub
	TypeMicrosoftServiceBusEventHub     TypeBasicOutputDataSource = original.TypeMicrosoftServiceBusEventHub
	TypeMicrosoftServiceBusQueue        TypeBasicOutputDataSource = original.TypeMicrosoftServiceBusQueue
	TypeMicrosoftServiceBusTopic        TypeBasicOutputDataSource = original.TypeMicrosoftServiceBusTopic
	TypeMicrosoftSQLServerDatabase      TypeBasicOutputDataSource = original.TypeMicrosoftSQLServerDatabase
	TypeMicrosoftSQLServerDataWarehouse TypeBasicOutputDataSource = original.TypeMicrosoftSQLServerDataWarehouse
	TypeMicrosoftStorageBlob            TypeBasicOutputDataSource = original.TypeMicrosoftStorageBlob
	TypeMicrosoftStorageDocumentDB      TypeBasicOutputDataSource = original.TypeMicrosoftStorageDocumentDB
	TypeMicrosoftStorageTable           TypeBasicOutputDataSource = original.TypeMicrosoftStorageTable
	TypeOutputDataSource                TypeBasicOutputDataSource = original.TypeOutputDataSource
	TypePowerBI                         TypeBasicOutputDataSource = original.TypePowerBI
)

type TypeBasicReferenceInputDataSource = original.TypeBasicReferenceInputDataSource

const (
	TypeBasicReferenceInputDataSourceTypeMicrosoftSQLServerDatabase TypeBasicReferenceInputDataSource = original.TypeBasicReferenceInputDataSourceTypeMicrosoftSQLServerDatabase
	TypeBasicReferenceInputDataSourceTypeMicrosoftStorageBlob       TypeBasicReferenceInputDataSource = original.TypeBasicReferenceInputDataSourceTypeMicrosoftStorageBlob
	TypeBasicReferenceInputDataSourceTypeReferenceInputDataSource   TypeBasicReferenceInputDataSource = original.TypeBasicReferenceInputDataSourceTypeReferenceInputDataSource
)

type TypeBasicSerialization = original.TypeBasicSerialization

const (
	TypeAvro          TypeBasicSerialization = original.TypeAvro
	TypeCsv           TypeBasicSerialization = original.TypeCsv
	TypeCustomClr     TypeBasicSerialization = original.TypeCustomClr
	TypeJSON          TypeBasicSerialization = original.TypeJSON
	TypeParquet       TypeBasicSerialization = original.TypeParquet
	TypeSerialization TypeBasicSerialization = original.TypeSerialization
)

type TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSource

const (
	TypeBasicStreamInputDataSourceTypeMicrosoftDevicesIotHubs     TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeMicrosoftDevicesIotHubs
	TypeBasicStreamInputDataSourceTypeMicrosoftEventHubEventHub   TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeMicrosoftEventHubEventHub
	TypeBasicStreamInputDataSourceTypeMicrosoftServiceBusEventHub TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeMicrosoftServiceBusEventHub
	TypeBasicStreamInputDataSourceTypeMicrosoftStorageBlob        TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeMicrosoftStorageBlob
	TypeBasicStreamInputDataSourceTypeStreamInputDataSource       TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeStreamInputDataSource
)

type UdfType = original.UdfType

const (
	Scalar UdfType = original.Scalar
)

type AggregateFunctionProperties = original.AggregateFunctionProperties
type AvroSerialization = original.AvroSerialization
type AzureDataLakeStoreOutputDataSource = original.AzureDataLakeStoreOutputDataSource
type AzureDataLakeStoreOutputDataSourceProperties = original.AzureDataLakeStoreOutputDataSourceProperties
type AzureFunctionOutputDataSource = original.AzureFunctionOutputDataSource
type AzureFunctionOutputDataSourceProperties = original.AzureFunctionOutputDataSourceProperties
type AzureMachineLearningServiceFunctionBinding = original.AzureMachineLearningServiceFunctionBinding
type AzureMachineLearningServiceFunctionBindingProperties = original.AzureMachineLearningServiceFunctionBindingProperties
type AzureMachineLearningServiceFunctionBindingRetrievalProperties = original.AzureMachineLearningServiceFunctionBindingRetrievalProperties
type AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters = original.AzureMachineLearningServiceFunctionRetrieveDefaultDefinitionParameters
type AzureMachineLearningServiceInputColumn = original.AzureMachineLearningServiceInputColumn
type AzureMachineLearningServiceInputs = original.AzureMachineLearningServiceInputs
type AzureMachineLearningServiceOutputColumn = original.AzureMachineLearningServiceOutputColumn
type AzureMachineLearningStudioFunctionBinding = original.AzureMachineLearningStudioFunctionBinding
type AzureMachineLearningStudioFunctionBindingProperties = original.AzureMachineLearningStudioFunctionBindingProperties
type AzureMachineLearningStudioFunctionBindingRetrievalProperties = original.AzureMachineLearningStudioFunctionBindingRetrievalProperties
type AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters = original.AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters
type AzureMachineLearningStudioInputColumn = original.AzureMachineLearningStudioInputColumn
type AzureMachineLearningStudioInputs = original.AzureMachineLearningStudioInputs
type AzureMachineLearningStudioOutputColumn = original.AzureMachineLearningStudioOutputColumn
type AzureSQLDatabaseDataSourceProperties = original.AzureSQLDatabaseDataSourceProperties
type AzureSQLDatabaseOutputDataSource = original.AzureSQLDatabaseOutputDataSource
type AzureSQLDatabaseOutputDataSourceProperties = original.AzureSQLDatabaseOutputDataSourceProperties
type AzureSQLReferenceInputDataSource = original.AzureSQLReferenceInputDataSource
type AzureSQLReferenceInputDataSourceProperties = original.AzureSQLReferenceInputDataSourceProperties
type AzureSynapseDataSourceProperties = original.AzureSynapseDataSourceProperties
type AzureSynapseOutputDataSource = original.AzureSynapseOutputDataSource
type AzureSynapseOutputDataSourceProperties = original.AzureSynapseOutputDataSourceProperties
type AzureTableOutputDataSource = original.AzureTableOutputDataSource
type AzureTableOutputDataSourceProperties = original.AzureTableOutputDataSourceProperties
type BaseClient = original.BaseClient
type BasicFunctionBinding = original.BasicFunctionBinding
type BasicFunctionProperties = original.BasicFunctionProperties
type BasicFunctionRetrieveDefaultDefinitionParameters = original.BasicFunctionRetrieveDefaultDefinitionParameters
type BasicInputProperties = original.BasicInputProperties
type BasicOutputDataSource = original.BasicOutputDataSource
type BasicReferenceInputDataSource = original.BasicReferenceInputDataSource
type BasicSerialization = original.BasicSerialization
type BasicStreamInputDataSource = original.BasicStreamInputDataSource
type BlobDataSourceProperties = original.BlobDataSourceProperties
type BlobOutputDataSource = original.BlobOutputDataSource
type BlobOutputDataSourceProperties = original.BlobOutputDataSourceProperties
type BlobReferenceInputDataSource = original.BlobReferenceInputDataSource
type BlobReferenceInputDataSourceProperties = original.BlobReferenceInputDataSourceProperties
type BlobStreamInputDataSource = original.BlobStreamInputDataSource
type BlobStreamInputDataSourceProperties = original.BlobStreamInputDataSourceProperties
type CSharpFunctionBinding = original.CSharpFunctionBinding
type CSharpFunctionBindingProperties = original.CSharpFunctionBindingProperties
type CSharpFunctionBindingRetrievalProperties = original.CSharpFunctionBindingRetrievalProperties
type CSharpFunctionRetrieveDefaultDefinitionParameters = original.CSharpFunctionRetrieveDefaultDefinitionParameters
type Cluster = original.Cluster
type ClusterInfo = original.ClusterInfo
type ClusterJob = original.ClusterJob
type ClusterJobListResult = original.ClusterJobListResult
type ClusterJobListResultIterator = original.ClusterJobListResultIterator
type ClusterJobListResultPage = original.ClusterJobListResultPage
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterProperties = original.ClusterProperties
type ClusterSku = original.ClusterSku
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type Compression = original.Compression
type CsvSerialization = original.CsvSerialization
type CsvSerializationProperties = original.CsvSerializationProperties
type CustomClrSerialization = original.CustomClrSerialization
type CustomClrSerializationProperties = original.CustomClrSerializationProperties
type DiagnosticCondition = original.DiagnosticCondition
type Diagnostics = original.Diagnostics
type DocumentDbOutputDataSource = original.DocumentDbOutputDataSource
type DocumentDbOutputDataSourceProperties = original.DocumentDbOutputDataSourceProperties
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorError = original.ErrorError
type ErrorResponse = original.ErrorResponse
type EventHubDataSourceProperties = original.EventHubDataSourceProperties
type EventHubOutputDataSource = original.EventHubOutputDataSource
type EventHubOutputDataSourceProperties = original.EventHubOutputDataSourceProperties
type EventHubStreamInputDataSource = original.EventHubStreamInputDataSource
type EventHubStreamInputDataSourceProperties = original.EventHubStreamInputDataSourceProperties
type EventHubV2OutputDataSource = original.EventHubV2OutputDataSource
type EventHubV2StreamInputDataSource = original.EventHubV2StreamInputDataSource
type External = original.External
type Function = original.Function
type FunctionBinding = original.FunctionBinding
type FunctionConfiguration = original.FunctionConfiguration
type FunctionInput = original.FunctionInput
type FunctionListResult = original.FunctionListResult
type FunctionListResultIterator = original.FunctionListResultIterator
type FunctionListResultPage = original.FunctionListResultPage
type FunctionOutput = original.FunctionOutput
type FunctionProperties = original.FunctionProperties
type FunctionRetrieveDefaultDefinitionParameters = original.FunctionRetrieveDefaultDefinitionParameters
type FunctionsClient = original.FunctionsClient
type FunctionsTestFuture = original.FunctionsTestFuture
type Identity = original.Identity
type Input = original.Input
type InputListResult = original.InputListResult
type InputListResultIterator = original.InputListResultIterator
type InputListResultPage = original.InputListResultPage
type InputProperties = original.InputProperties
type InputsClient = original.InputsClient
type InputsTestFuture = original.InputsTestFuture
type IoTHubStreamInputDataSource = original.IoTHubStreamInputDataSource
type IoTHubStreamInputDataSourceProperties = original.IoTHubStreamInputDataSourceProperties
type JSONSerialization = original.JSONSerialization
type JSONSerializationProperties = original.JSONSerializationProperties
type JavaScriptFunctionBinding = original.JavaScriptFunctionBinding
type JavaScriptFunctionBindingProperties = original.JavaScriptFunctionBindingProperties
type JavaScriptFunctionBindingRetrievalProperties = original.JavaScriptFunctionBindingRetrievalProperties
type JavaScriptFunctionRetrieveDefaultDefinitionParameters = original.JavaScriptFunctionRetrieveDefaultDefinitionParameters
type JobStorageAccount = original.JobStorageAccount
type OAuthBasedDataSourceProperties = original.OAuthBasedDataSourceProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Output = original.Output
type OutputDataSource = original.OutputDataSource
type OutputListResult = original.OutputListResult
type OutputListResultIterator = original.OutputListResultIterator
type OutputListResultPage = original.OutputListResultPage
type OutputProperties = original.OutputProperties
type OutputsClient = original.OutputsClient
type OutputsTestFuture = original.OutputsTestFuture
type ParquetSerialization = original.ParquetSerialization
type PowerBIOutputDataSource = original.PowerBIOutputDataSource
type PowerBIOutputDataSourceProperties = original.PowerBIOutputDataSourceProperties
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointListResult = original.PrivateEndpointListResult
type PrivateEndpointListResultIterator = original.PrivateEndpointListResultIterator
type PrivateEndpointListResultPage = original.PrivateEndpointListResultPage
type PrivateEndpointProperties = original.PrivateEndpointProperties
type PrivateEndpointsClient = original.PrivateEndpointsClient
type PrivateEndpointsDeleteFuture = original.PrivateEndpointsDeleteFuture
type PrivateLinkConnectionState = original.PrivateLinkConnectionState
type PrivateLinkServiceConnection = original.PrivateLinkServiceConnection
type PrivateLinkServiceConnectionProperties = original.PrivateLinkServiceConnectionProperties
type ProxyResource = original.ProxyResource
type ReferenceInputDataSource = original.ReferenceInputDataSource
type ReferenceInputProperties = original.ReferenceInputProperties
type Resource = original.Resource
type ResourceTestStatus = original.ResourceTestStatus
type ScalarFunctionProperties = original.ScalarFunctionProperties
type Serialization = original.Serialization
type ServiceBusDataSourceProperties = original.ServiceBusDataSourceProperties
type ServiceBusQueueOutputDataSource = original.ServiceBusQueueOutputDataSource
type ServiceBusQueueOutputDataSourceProperties = original.ServiceBusQueueOutputDataSourceProperties
type ServiceBusTopicOutputDataSource = original.ServiceBusTopicOutputDataSource
type ServiceBusTopicOutputDataSourceProperties = original.ServiceBusTopicOutputDataSourceProperties
type StartStreamingJobParameters = original.StartStreamingJobParameters
type StorageAccount = original.StorageAccount
type StreamInputDataSource = original.StreamInputDataSource
type StreamInputProperties = original.StreamInputProperties
type StreamingJob = original.StreamingJob
type StreamingJobListResult = original.StreamingJobListResult
type StreamingJobListResultIterator = original.StreamingJobListResultIterator
type StreamingJobListResultPage = original.StreamingJobListResultPage
type StreamingJobProperties = original.StreamingJobProperties
type StreamingJobSku = original.StreamingJobSku
type StreamingJobsClient = original.StreamingJobsClient
type StreamingJobsCreateOrReplaceFuture = original.StreamingJobsCreateOrReplaceFuture
type StreamingJobsDeleteFuture = original.StreamingJobsDeleteFuture
type StreamingJobsStartFuture = original.StreamingJobsStartFuture
type StreamingJobsStopFuture = original.StreamingJobsStopFuture
type SubResource = original.SubResource
type SubscriptionQuota = original.SubscriptionQuota
type SubscriptionQuotaProperties = original.SubscriptionQuotaProperties
type SubscriptionQuotasListResult = original.SubscriptionQuotasListResult
type SubscriptionsClient = original.SubscriptionsClient
type TrackedResource = original.TrackedResource
type Transformation = original.Transformation
type TransformationProperties = original.TransformationProperties
type TransformationsClient = original.TransformationsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClusterJobListResultIterator(page ClusterJobListResultPage) ClusterJobListResultIterator {
	return original.NewClusterJobListResultIterator(page)
}
func NewClusterJobListResultPage(cur ClusterJobListResult, getNextPage func(context.Context, ClusterJobListResult) (ClusterJobListResult, error)) ClusterJobListResultPage {
	return original.NewClusterJobListResultPage(cur, getNextPage)
}
func NewClusterListResultIterator(page ClusterListResultPage) ClusterListResultIterator {
	return original.NewClusterListResultIterator(page)
}
func NewClusterListResultPage(cur ClusterListResult, getNextPage func(context.Context, ClusterListResult) (ClusterListResult, error)) ClusterListResultPage {
	return original.NewClusterListResultPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewFunctionListResultIterator(page FunctionListResultPage) FunctionListResultIterator {
	return original.NewFunctionListResultIterator(page)
}
func NewFunctionListResultPage(cur FunctionListResult, getNextPage func(context.Context, FunctionListResult) (FunctionListResult, error)) FunctionListResultPage {
	return original.NewFunctionListResultPage(cur, getNextPage)
}
func NewFunctionsClient(subscriptionID string) FunctionsClient {
	return original.NewFunctionsClient(subscriptionID)
}
func NewFunctionsClientWithBaseURI(baseURI string, subscriptionID string) FunctionsClient {
	return original.NewFunctionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInputListResultIterator(page InputListResultPage) InputListResultIterator {
	return original.NewInputListResultIterator(page)
}
func NewInputListResultPage(cur InputListResult, getNextPage func(context.Context, InputListResult) (InputListResult, error)) InputListResultPage {
	return original.NewInputListResultPage(cur, getNextPage)
}
func NewInputsClient(subscriptionID string) InputsClient {
	return original.NewInputsClient(subscriptionID)
}
func NewInputsClientWithBaseURI(baseURI string, subscriptionID string) InputsClient {
	return original.NewInputsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOutputListResultIterator(page OutputListResultPage) OutputListResultIterator {
	return original.NewOutputListResultIterator(page)
}
func NewOutputListResultPage(cur OutputListResult, getNextPage func(context.Context, OutputListResult) (OutputListResult, error)) OutputListResultPage {
	return original.NewOutputListResultPage(cur, getNextPage)
}
func NewOutputsClient(subscriptionID string) OutputsClient {
	return original.NewOutputsClient(subscriptionID)
}
func NewOutputsClientWithBaseURI(baseURI string, subscriptionID string) OutputsClient {
	return original.NewOutputsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointListResultIterator(page PrivateEndpointListResultPage) PrivateEndpointListResultIterator {
	return original.NewPrivateEndpointListResultIterator(page)
}
func NewPrivateEndpointListResultPage(cur PrivateEndpointListResult, getNextPage func(context.Context, PrivateEndpointListResult) (PrivateEndpointListResult, error)) PrivateEndpointListResultPage {
	return original.NewPrivateEndpointListResultPage(cur, getNextPage)
}
func NewPrivateEndpointsClient(subscriptionID string) PrivateEndpointsClient {
	return original.NewPrivateEndpointsClient(subscriptionID)
}
func NewPrivateEndpointsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointsClient {
	return original.NewPrivateEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStreamingJobListResultIterator(page StreamingJobListResultPage) StreamingJobListResultIterator {
	return original.NewStreamingJobListResultIterator(page)
}
func NewStreamingJobListResultPage(cur StreamingJobListResult, getNextPage func(context.Context, StreamingJobListResult) (StreamingJobListResult, error)) StreamingJobListResultPage {
	return original.NewStreamingJobListResultPage(cur, getNextPage)
}
func NewStreamingJobsClient(subscriptionID string) StreamingJobsClient {
	return original.NewStreamingJobsClient(subscriptionID)
}
func NewStreamingJobsClientWithBaseURI(baseURI string, subscriptionID string) StreamingJobsClient {
	return original.NewStreamingJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClient(subscriptionID)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransformationsClient(subscriptionID string) TransformationsClient {
	return original.NewTransformationsClient(subscriptionID)
}
func NewTransformationsClientWithBaseURI(baseURI string, subscriptionID string) TransformationsClient {
	return original.NewTransformationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthenticationModeValues() []AuthenticationMode {
	return original.PossibleAuthenticationModeValues()
}
func PossibleBindingTypeValues() []BindingType {
	return original.PossibleBindingTypeValues()
}
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return original.PossibleClusterProvisioningStateValues()
}
func PossibleClusterSkuNameValues() []ClusterSkuName {
	return original.PossibleClusterSkuNameValues()
}
func PossibleCompatibilityLevelValues() []CompatibilityLevel {
	return original.PossibleCompatibilityLevelValues()
}
func PossibleContentStoragePolicyValues() []ContentStoragePolicy {
	return original.PossibleContentStoragePolicyValues()
}
func PossibleEncodingValues() []Encoding {
	return original.PossibleEncodingValues()
}
func PossibleEventSerializationTypeValues() []EventSerializationType {
	return original.PossibleEventSerializationTypeValues()
}
func PossibleEventsOutOfOrderPolicyValues() []EventsOutOfOrderPolicy {
	return original.PossibleEventsOutOfOrderPolicyValues()
}
func PossibleJSONOutputSerializationFormatValues() []JSONOutputSerializationFormat {
	return original.PossibleJSONOutputSerializationFormatValues()
}
func PossibleJobStateValues() []JobState {
	return original.PossibleJobStateValues()
}
func PossibleJobTypeValues() []JobType {
	return original.PossibleJobTypeValues()
}
func PossibleOutputErrorPolicyValues() []OutputErrorPolicy {
	return original.PossibleOutputErrorPolicyValues()
}
func PossibleOutputStartModeValues() []OutputStartMode {
	return original.PossibleOutputStartModeValues()
}
func PossibleStreamingJobSkuNameValues() []StreamingJobSkuName {
	return original.PossibleStreamingJobSkuNameValues()
}
func PossibleTypeBasicFunctionPropertiesValues() []TypeBasicFunctionProperties {
	return original.PossibleTypeBasicFunctionPropertiesValues()
}
func PossibleTypeBasicInputPropertiesValues() []TypeBasicInputProperties {
	return original.PossibleTypeBasicInputPropertiesValues()
}
func PossibleTypeBasicOutputDataSourceValues() []TypeBasicOutputDataSource {
	return original.PossibleTypeBasicOutputDataSourceValues()
}
func PossibleTypeBasicReferenceInputDataSourceValues() []TypeBasicReferenceInputDataSource {
	return original.PossibleTypeBasicReferenceInputDataSourceValues()
}
func PossibleTypeBasicSerializationValues() []TypeBasicSerialization {
	return original.PossibleTypeBasicSerializationValues()
}
func PossibleTypeBasicStreamInputDataSourceValues() []TypeBasicStreamInputDataSource {
	return original.PossibleTypeBasicStreamInputDataSourceValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUdfTypeValues() []UdfType {
	return original.PossibleUdfTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}