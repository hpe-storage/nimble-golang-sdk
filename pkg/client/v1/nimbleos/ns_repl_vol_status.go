// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReplVolStatus - The replication status of a volume undergoing replication.
// Export NsReplVolStatusFields for advance operations like search filter etc.
var NsReplVolStatusFields *NsReplVolStatusStringFields

func init() {
	Namefield := "name"
	SnapNamefield := "snap_name"
	Statusfield := "status"
	InternalStatusfield := "internal_status"
	ReplBytesDonefield := "repl_bytes_done"
	ReplBytesTotalfield := "repl_bytes_total"

	NsReplVolStatusFields = &NsReplVolStatusStringFields{
		Name:           &Namefield,
		SnapName:       &SnapNamefield,
		Status:         &Statusfield,
		InternalStatus: &InternalStatusfield,
		ReplBytesDone:  &ReplBytesDonefield,
		ReplBytesTotal: &ReplBytesTotalfield,
	}
}

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

// Struct for NsReplVolStatusFields
type NsReplVolStatusStringFields struct {
	Name           *string
	SnapName       *string
	Status         *string
	InternalStatus *string
	ReplBytesDone  *string
	ReplBytesTotal *string
}
