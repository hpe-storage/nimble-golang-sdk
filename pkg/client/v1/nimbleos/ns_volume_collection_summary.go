// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeCollectionSummary - Select fields of volume collection info.

// Export NsVolumeCollectionSummaryFields provides field names to use in filter parameters, for example.
var NsVolumeCollectionSummaryFields *NsVolumeCollectionSummaryStringFields

func init() {
	fieldID := "id"
	fieldName := "name"

	NsVolumeCollectionSummaryFields = &NsVolumeCollectionSummaryStringFields{
		ID:   &fieldID,
		Name: &fieldName,
	}
}

type NsVolumeCollectionSummary struct {
	// ID - Identifier of volume collection.
	ID *string `json:"id,omitempty"`
	// Name - Name of volume collection.
	Name *string `json:"name,omitempty"`
}

// Struct for NsVolumeCollectionSummaryFields
type NsVolumeCollectionSummaryStringFields struct {
	ID   *string
	Name *string
}
