// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapSummary - Select fields containing snapshot information.
// Export NsSnapSummaryFields for advance operations like search filter etc.
var NsSnapSummaryFields *NsSnapSummary

func init() {
	SnapIdfield := "snap_id"
	SnapNamefield := "snap_name"

	NsSnapSummaryFields = &NsSnapSummary{
		SnapId:   &SnapIdfield,
		SnapName: &SnapNamefield,
	}
}

type NsSnapSummary struct {
	// SnapId - ID of snapshot.
	SnapId *string `json:"snap_id,omitempty"`
	// SnapName - Name of snapshot.
	SnapName *string `json:"snap_name,omitempty"`
	// SnapCreationTime - Creation time of snapshot.
	SnapCreationTime *int64 `json:"snap_creation_time,omitempty"`
}
