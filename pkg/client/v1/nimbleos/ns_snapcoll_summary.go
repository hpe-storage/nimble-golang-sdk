// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSnapcollSummaryFields provides field names to use in filter parameters, for example.
var NsSnapcollSummaryFields *NsSnapcollSummaryFieldHandles

func init() {
	fieldSnapcollId := "snapcoll_id"
	fieldSnapcollName := "snapcoll_name"
	fieldSnapcollCreationTime := "snapcoll_creation_time"

	NsSnapcollSummaryFields = &NsSnapcollSummaryFieldHandles{
		SnapcollId:           &fieldSnapcollId,
		SnapcollName:         &fieldSnapcollName,
		SnapcollCreationTime: &fieldSnapcollCreationTime,
	}
}

// NsSnapcollSummary - Select fields containing snapshot collection information.
type NsSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
	SnapcollId *string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
	SnapcollName *string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
}

// NsSnapcollSummaryFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapcollSummaryFieldHandles struct {
	SnapcollId           *string
	SnapcollName         *string
	SnapcollCreationTime *string
}
