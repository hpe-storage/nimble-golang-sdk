// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSnapSummaryFields provides field names to use in filter parameters, for example.
var NsSnapSummaryFields *NsSnapSummaryFieldHandles

func init() {
	fieldSnapId := "snap_id"
	fieldSnapName := "snap_name"
	fieldSnapCreationTime := "snap_creation_time"

	NsSnapSummaryFields = &NsSnapSummaryFieldHandles{
		SnapId:           &fieldSnapId,
		SnapName:         &fieldSnapName,
		SnapCreationTime: &fieldSnapCreationTime,
	}
}

// NsSnapSummary - Select fields containing snapshot information.
type NsSnapSummary struct {
	// SnapId - ID of snapshot.
	SnapId *string `json:"snap_id,omitempty"`
	// SnapName - Name of snapshot.
	SnapName *string `json:"snap_name,omitempty"`
	// SnapCreationTime - Creation time of snapshot.
	SnapCreationTime *int64 `json:"snap_creation_time,omitempty"`
}

// NsSnapSummaryFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapSummaryFieldHandles struct {
	SnapId           *string
	SnapName         *string
	SnapCreationTime *string
}
