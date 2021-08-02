// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapcollSummaryFields provides field names to use in filter parameters, for example.
var NsSnapcollSummaryFields *NsSnapcollSummaryFieldHandles

func init() {
	NsSnapcollSummaryFields = &NsSnapcollSummaryFieldHandles{
		SnapcollId:           "snapcoll_id",
		SnapcollName:         "snapcoll_name",
		SnapcollCreationTime: "snapcoll_creation_time",
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
	SnapcollId           string
	SnapcollName         string
	SnapcollCreationTime string
}
