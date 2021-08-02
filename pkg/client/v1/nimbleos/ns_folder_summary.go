// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFolderSummaryFields provides field names to use in filter parameters, for example.
var NsFolderSummaryFields *NsFolderSummaryFieldHandles

func init() {
	NsFolderSummaryFields = &NsFolderSummaryFieldHandles{
		ID:  "id",
		Fqn: "fqn",
	}
}

// NsFolderSummary - Select fields containing folder info.
type NsFolderSummary struct {
	// ID - ID of folder.
	ID *string `json:"id,omitempty"`
	// Fqn - Fully qualified name of folder.
	Fqn *string `json:"fqn,omitempty"`
}

// NsFolderSummaryFieldHandles provides a string representation for each NsFolderSummary field.
type NsFolderSummaryFieldHandles struct {
	ID  string
	Fqn string
}
