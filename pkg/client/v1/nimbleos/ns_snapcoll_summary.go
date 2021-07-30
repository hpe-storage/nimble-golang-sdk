// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapcollSummary - Select fields containing snapshot collection information.
// Export NsSnapcollSummaryFields for advance operations like search filter etc.
var NsSnapcollSummaryFields *NsSnapcollSummaryStringFields

func init() {
	SnapcollIdfield := "snapcoll_id"
	SnapcollNamefield := "snapcoll_name"
	SnapcollCreationTimefield := "snapcoll_creation_time"

	NsSnapcollSummaryFields = &NsSnapcollSummaryStringFields{
		SnapcollId:           &SnapcollIdfield,
		SnapcollName:         &SnapcollNamefield,
		SnapcollCreationTime: &SnapcollCreationTimefield,
	}
}

type NsSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
	SnapcollId *string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
	SnapcollName *string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
}

// Struct for NsSnapcollSummaryFields
type NsSnapcollSummaryStringFields struct {
	SnapcollId           *string
	SnapcollName         *string
	SnapcollCreationTime *string
}
