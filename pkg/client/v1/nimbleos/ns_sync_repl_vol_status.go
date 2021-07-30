// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSyncReplVolStatus - The sync replication status of a volume in volume collection.

// Export NsSyncReplVolStatusFields provides field names to use in filter parameters, for example.
var NsSyncReplVolStatusFields *NsSyncReplVolStatusStringFields

func init() {
	fieldResyncActive := "resync_active"
	fieldResyncBytesDone := "resync_bytes_done"
	fieldResyncBytesTotal := "resync_bytes_total"

	NsSyncReplVolStatusFields = &NsSyncReplVolStatusStringFields{
		ResyncActive:     &fieldResyncActive,
		ResyncBytesDone:  &fieldResyncBytesDone,
		ResyncBytesTotal: &fieldResyncBytesTotal,
	}
}

type NsSyncReplVolStatus struct {
	// ResyncActive - Sync replication active status.
	ResyncActive *bool `json:"resync_active,omitempty"`
	// ResyncBytesDone - Transferred bytes.
	ResyncBytesDone *int64 `json:"resync_bytes_done,omitempty"`
	// ResyncBytesTotal - Total number of bytes to be transferred.
	ResyncBytesTotal *int64 `json:"resync_bytes_total,omitempty"`
}

// Struct for NsSyncReplVolStatusFields
type NsSyncReplVolStatusStringFields struct {
	ResyncActive     *string
	ResyncBytesDone  *string
	ResyncBytesTotal *string
}
