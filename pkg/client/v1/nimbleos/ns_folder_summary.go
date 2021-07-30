// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFolderSummary - Select fields containing folder info.

// Export NsFolderSummaryFields provides field names to use in filter parameters, for example.
var NsFolderSummaryFields *NsFolderSummaryStringFields

func init() {
	fieldID := "id"
	fieldFqn := "fqn"

	NsFolderSummaryFields = &NsFolderSummaryStringFields{
		ID:  &fieldID,
		Fqn: &fieldFqn,
	}
}

type NsFolderSummary struct {
	// ID - ID of folder.
	ID *string `json:"id,omitempty"`
	// Fqn - Fully qualified name of folder.
	Fqn *string `json:"fqn,omitempty"`
}

// Struct for NsFolderSummaryFields
type NsFolderSummaryStringFields struct {
	ID  *string
	Fqn *string
}
