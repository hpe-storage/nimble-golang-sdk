// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsReplVolStatusFields provides field names to use in filter parameters, for example.
var NsReplVolStatusFields *NsReplVolStatusFieldHandles

func init() {
	fieldName := "name"
	fieldSnapName := "snap_name"
	fieldStatus := "status"
	fieldInternalStatus := "internal_status"
	fieldReplBytesDone := "repl_bytes_done"
	fieldReplBytesTotal := "repl_bytes_total"

	NsReplVolStatusFields = &NsReplVolStatusFieldHandles{
		Name:           &fieldName,
		SnapName:       &fieldSnapName,
		Status:         &fieldStatus,
		InternalStatus: &fieldInternalStatus,
		ReplBytesDone:  &fieldReplBytesDone,
		ReplBytesTotal: &fieldReplBytesTotal,
	}
}

// NsReplVolStatus - The replication status of a volume undergoing replication.
type NsReplVolStatus struct {
	// Name - Name of the volume being replicated.
	Name *string `json:"name,omitempty"`
	// SnapName - Name of the snapshot being replicated.
	SnapName *string `json:"snap_name,omitempty"`
	// Status - Replication status of the volume.
	Status *NsReplVolPartnerStatus `json:"status,omitempty"`
	// InternalStatus - Internal replication status of the volume.
	InternalStatus *NsReplVolPartnerStatus `json:"internal_status,omitempty"`
	// ReplBytesDone - Transferred bytes.
	ReplBytesDone *int64 `json:"repl_bytes_done,omitempty"`
	// ReplBytesTotal - Total number of bytes to be transferred.
	ReplBytesTotal *int64 `json:"repl_bytes_total,omitempty"`
}

// NsReplVolStatusFieldHandles provides a string representation for each AccessControlRecord field.
type NsReplVolStatusFieldHandles struct {
	Name           *string
	SnapName       *string
	Status         *string
	InternalStatus *string
	ReplBytesDone  *string
	ReplBytesTotal *string
}
