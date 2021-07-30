// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPoolSummary - Object containing pool ID and name.

// Export NsPoolSummaryFields provides field names to use in filter parameters, for example.
var NsPoolSummaryFields *NsPoolSummaryStringFields

func init() {
	fieldID := "id"
	fieldName := "name"

	NsPoolSummaryFields = &NsPoolSummaryStringFields{
		ID:   &fieldID,
		Name: &fieldName,
	}
}

type NsPoolSummary struct {
	// ID - ID of pool.
	ID *string `json:"id,omitempty"`
	// Name - Name of pool.
	Name *string `json:"name,omitempty"`
}

// Struct for NsPoolSummaryFields
type NsPoolSummaryStringFields struct {
	ID   *string
	Name *string
}
