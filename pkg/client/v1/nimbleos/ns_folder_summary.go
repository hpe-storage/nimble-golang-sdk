// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFolderSummary - Select fields containing folder info.
// Export NsFolderSummaryFields for advance operations like search filter etc.
var NsFolderSummaryFields *NsFolderSummaryStringFields

func init() {
	IDfield := "id"
	Fqnfield := "fqn"

	NsFolderSummaryFields = &NsFolderSummaryStringFields{
		ID:  &IDfield,
		Fqn: &Fqnfield,
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
