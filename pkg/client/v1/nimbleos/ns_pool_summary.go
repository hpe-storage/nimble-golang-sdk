// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPoolSummaryFields provides field names to use in filter parameters, for example.
var NsPoolSummaryFields *NsPoolSummaryFieldHandles

func init() {
	NsPoolSummaryFields = &NsPoolSummaryFieldHandles{
		ID:   "id",
		Name: "name",
	}
}

// NsPoolSummary - Object containing pool ID and name.
type NsPoolSummary struct {
	// ID - ID of pool.
	ID *string `json:"id,omitempty"`
	// Name - Name of pool.
	Name *string `json:"name,omitempty"`
}

// NsPoolSummaryFieldHandles provides a string representation for each AccessControlRecord field.
type NsPoolSummaryFieldHandles struct {
	ID   string
	Name string
}
