package storagesync

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CloudTiering enumerates the values for cloud tiering.
type CloudTiering string

const (
	// Off ...
	Off CloudTiering = "off"
	// On ...
	On CloudTiering = "on"
)

// PossibleCloudTieringValues returns an array of possible values for the CloudTiering const type.
func PossibleCloudTieringValues() []CloudTiering {
	return []CloudTiering{Off, On}
}

// CloudTiering1 enumerates the values for cloud tiering 1.
type CloudTiering1 string

const (
	// CloudTiering1Off ...
	CloudTiering1Off CloudTiering1 = "off"
	// CloudTiering1On ...
	CloudTiering1On CloudTiering1 = "on"
)

// PossibleCloudTiering1Values returns an array of possible values for the CloudTiering1 const type.
func PossibleCloudTiering1Values() []CloudTiering1 {
	return []CloudTiering1{CloudTiering1Off, CloudTiering1On}
}

// CloudTiering2 enumerates the values for cloud tiering 2.
type CloudTiering2 string

const (
	// CloudTiering2Off ...
	CloudTiering2Off CloudTiering2 = "off"
	// CloudTiering2On ...
	CloudTiering2On CloudTiering2 = "on"
)

// PossibleCloudTiering2Values returns an array of possible values for the CloudTiering2 const type.
func PossibleCloudTiering2Values() []CloudTiering2 {
	return []CloudTiering2{CloudTiering2Off, CloudTiering2On}
}

// CombinedHealth enumerates the values for combined health.
type CombinedHealth string

const (
	// CombinedHealthError ...
	CombinedHealthError CombinedHealth = "Error"
	// CombinedHealthHealthy ...
	CombinedHealthHealthy CombinedHealth = "Healthy"
	// CombinedHealthNoActivity ...
	CombinedHealthNoActivity CombinedHealth = "NoActivity"
	// CombinedHealthSyncBlockedForChangeDetectionPostRestore ...
	CombinedHealthSyncBlockedForChangeDetectionPostRestore CombinedHealth = "SyncBlockedForChangeDetectionPostRestore"
	// CombinedHealthSyncBlockedForRestore ...
	CombinedHealthSyncBlockedForRestore CombinedHealth = "SyncBlockedForRestore"
)

// PossibleCombinedHealthValues returns an array of possible values for the CombinedHealth const type.
func PossibleCombinedHealthValues() []CombinedHealth {
	return []CombinedHealth{CombinedHealthError, CombinedHealthHealthy, CombinedHealthNoActivity, CombinedHealthSyncBlockedForChangeDetectionPostRestore, CombinedHealthSyncBlockedForRestore}
}

// DownloadHealth enumerates the values for download health.
type DownloadHealth string

const (
	// DownloadHealthError ...
	DownloadHealthError DownloadHealth = "Error"
	// DownloadHealthHealthy ...
	DownloadHealthHealthy DownloadHealth = "Healthy"
	// DownloadHealthNoActivity ...
	DownloadHealthNoActivity DownloadHealth = "NoActivity"
	// DownloadHealthSyncBlockedForChangeDetectionPostRestore ...
	DownloadHealthSyncBlockedForChangeDetectionPostRestore DownloadHealth = "SyncBlockedForChangeDetectionPostRestore"
	// DownloadHealthSyncBlockedForRestore ...
	DownloadHealthSyncBlockedForRestore DownloadHealth = "SyncBlockedForRestore"
)

// PossibleDownloadHealthValues returns an array of possible values for the DownloadHealth const type.
func PossibleDownloadHealthValues() []DownloadHealth {
	return []DownloadHealth{DownloadHealthError, DownloadHealthHealthy, DownloadHealthNoActivity, DownloadHealthSyncBlockedForChangeDetectionPostRestore, DownloadHealthSyncBlockedForRestore}
}

// NameAvailabilityReason enumerates the values for name availability reason.
type NameAvailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists NameAvailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid NameAvailabilityReason = "Invalid"
)

// PossibleNameAvailabilityReasonValues returns an array of possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{AlreadyExists, Invalid}
}

// OfflineDataTransfer enumerates the values for offline data transfer.
type OfflineDataTransfer string

const (
	// OfflineDataTransferOff ...
	OfflineDataTransferOff OfflineDataTransfer = "off"
	// OfflineDataTransferOn ...
	OfflineDataTransferOn OfflineDataTransfer = "on"
)

// PossibleOfflineDataTransferValues returns an array of possible values for the OfflineDataTransfer const type.
func PossibleOfflineDataTransferValues() []OfflineDataTransfer {
	return []OfflineDataTransfer{OfflineDataTransferOff, OfflineDataTransferOn}
}

// OfflineDataTransfer1 enumerates the values for offline data transfer 1.
type OfflineDataTransfer1 string

const (
	// OfflineDataTransfer1Off ...
	OfflineDataTransfer1Off OfflineDataTransfer1 = "off"
	// OfflineDataTransfer1On ...
	OfflineDataTransfer1On OfflineDataTransfer1 = "on"
)

// PossibleOfflineDataTransfer1Values returns an array of possible values for the OfflineDataTransfer1 const type.
func PossibleOfflineDataTransfer1Values() []OfflineDataTransfer1 {
	return []OfflineDataTransfer1{OfflineDataTransfer1Off, OfflineDataTransfer1On}
}

// OfflineDataTransfer2 enumerates the values for offline data transfer 2.
type OfflineDataTransfer2 string

const (
	// OfflineDataTransfer2Off ...
	OfflineDataTransfer2Off OfflineDataTransfer2 = "off"
	// OfflineDataTransfer2On ...
	OfflineDataTransfer2On OfflineDataTransfer2 = "on"
)

// PossibleOfflineDataTransfer2Values returns an array of possible values for the OfflineDataTransfer2 const type.
func PossibleOfflineDataTransfer2Values() []OfflineDataTransfer2 {
	return []OfflineDataTransfer2{OfflineDataTransfer2Off, OfflineDataTransfer2On}
}

// OfflineDataTransferStatus enumerates the values for offline data transfer status.
type OfflineDataTransferStatus string

const (
	// Complete ...
	Complete OfflineDataTransferStatus = "Complete"
	// InProgress ...
	InProgress OfflineDataTransferStatus = "InProgress"
	// NotRunning ...
	NotRunning OfflineDataTransferStatus = "NotRunning"
	// Stopping ...
	Stopping OfflineDataTransferStatus = "Stopping"
)

// PossibleOfflineDataTransferStatusValues returns an array of possible values for the OfflineDataTransferStatus const type.
func PossibleOfflineDataTransferStatusValues() []OfflineDataTransferStatus {
	return []OfflineDataTransferStatus{Complete, InProgress, NotRunning, Stopping}
}

// Operation enumerates the values for operation.
type Operation string

const (
	// Cancel ...
	Cancel Operation = "cancel"
	// Do ...
	Do Operation = "do"
	// Undo ...
	Undo Operation = "undo"
)

// PossibleOperationValues returns an array of possible values for the Operation const type.
func PossibleOperationValues() []Operation {
	return []Operation{Cancel, Do, Undo}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// Deleted ...
	Deleted Reason = "Deleted"
	// Registered ...
	Registered Reason = "Registered"
	// Suspended ...
	Suspended Reason = "Suspended"
	// Unregistered ...
	Unregistered Reason = "Unregistered"
	// Warned ...
	Warned Reason = "Warned"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{Deleted, Registered, Suspended, Unregistered, Warned}
}

// Status enumerates the values for status.
type Status string

const (
	// Aborted ...
	Aborted Status = "aborted"
	// Active ...
	Active Status = "active"
	// Expired ...
	Expired Status = "expired"
	// Failed ...
	Failed Status = "failed"
	// Succeeded ...
	Succeeded Status = "succeeded"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Aborted, Active, Expired, Failed, Succeeded}
}

// SyncActivity enumerates the values for sync activity.
type SyncActivity string

const (
	// Download ...
	Download SyncActivity = "Download"
	// Upload ...
	Upload SyncActivity = "Upload"
	// UploadAndDownload ...
	UploadAndDownload SyncActivity = "UploadAndDownload"
)

// PossibleSyncActivityValues returns an array of possible values for the SyncActivity const type.
func PossibleSyncActivityValues() []SyncActivity {
	return []SyncActivity{Download, Upload, UploadAndDownload}
}

// UploadHealth enumerates the values for upload health.
type UploadHealth string

const (
	// UploadHealthError ...
	UploadHealthError UploadHealth = "Error"
	// UploadHealthHealthy ...
	UploadHealthHealthy UploadHealth = "Healthy"
	// UploadHealthNoActivity ...
	UploadHealthNoActivity UploadHealth = "NoActivity"
	// UploadHealthSyncBlockedForChangeDetectionPostRestore ...
	UploadHealthSyncBlockedForChangeDetectionPostRestore UploadHealth = "SyncBlockedForChangeDetectionPostRestore"
	// UploadHealthSyncBlockedForRestore ...
	UploadHealthSyncBlockedForRestore UploadHealth = "SyncBlockedForRestore"
)

// PossibleUploadHealthValues returns an array of possible values for the UploadHealth const type.
func PossibleUploadHealthValues() []UploadHealth {
	return []UploadHealth{UploadHealthError, UploadHealthHealthy, UploadHealthNoActivity, UploadHealthSyncBlockedForChangeDetectionPostRestore, UploadHealthSyncBlockedForRestore}
}