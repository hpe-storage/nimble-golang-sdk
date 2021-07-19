// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArraySummary - Array summary information.
// Export NsArraySummaryFields for advance operations like search filter etc.
var NsArraySummaryFields *NsArraySummary

func init() {
	IDfield := "id"
	ArrayIdfield := "array_id"
	Namefield := "name"
	ArrayNamefield := "array_name"

	NsArraySummaryFields = &NsArraySummary{
		ID:        &IDfield,
		ArrayId:   &ArrayIdfield,
		Name:      &Namefield,
		ArrayName: &ArrayNamefield,
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
