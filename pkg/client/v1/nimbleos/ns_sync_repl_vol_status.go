// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSyncReplVolStatus - The sync replication status of a volume in volume collection.
// Export NsSyncReplVolStatusFields for advance operations like search filter etc.
var NsSyncReplVolStatusFields *NsSyncReplVolStatus

func init() {

	NsSyncReplVolStatusFields = &NsSyncReplVolStatus{}
}

type NsSyncReplVolStatus struct {
	// ResyncActive - Sync replication active status.
	ResyncActive *bool `json:"resync_active,omitempty"`
	// ResyncBytesDone - Transferred bytes.
	ResyncBytesDone *int64 `json:"resync_bytes_done,omitempty"`
	// ResyncBytesTotal - Total number of bytes to be transferred.
	ResyncBytesTotal *int64 `json:"resync_bytes_total,omitempty"`
}
