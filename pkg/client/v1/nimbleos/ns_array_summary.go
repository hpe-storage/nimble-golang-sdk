// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArraySummary - Array summary information.

// Export NsArraySummaryFields provides field names to use in filter parameters, for example.
var NsArraySummaryFields *NsArraySummaryStringFields

func init() {
	fieldID := "id"
	fieldArrayId := "array_id"
	fieldName := "name"
	fieldArrayName := "array_name"

	NsArraySummaryFields = &NsArraySummaryStringFields{
		ID:        &fieldID,
		ArrayId:   &fieldArrayId,
		Name:      &fieldName,
		ArrayName: &fieldArrayName,
	}
}

type NsArraySummary struct {
	// ID - Array API ID.
	ID *string `json:"id,omitempty"`
	// ArrayId - Array API ID.
	ArrayId *string `json:"array_id,omitempty"`
	// Name - Unique name of array.
	Name *string `json:"name,omitempty"`
	// ArrayName - Unique name of array.
	ArrayName *string `json:"array_name,omitempty"`
}

// Struct for NsArraySummaryFields
type NsArraySummaryStringFields struct {
	ID        *string
	ArrayId   *string
	Name      *string
	ArrayName *string
}
