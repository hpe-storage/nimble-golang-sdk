// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsVolumeCollectionSummaryFields provides field names to use in filter parameters, for example.
var NsVolumeCollectionSummaryFields *NsVolumeCollectionSummaryFieldHandles

func init() {
	NsVolumeCollectionSummaryFields = &NsVolumeCollectionSummaryFieldHandles{
		ID:   "id",
		Name: "name",
	}
}

// NsVolumeCollectionSummary - Select fields of volume collection info.
type NsVolumeCollectionSummary struct {
	// ID - Identifier of volume collection.
	ID *string `json:"id,omitempty"`
	// Name - Name of volume collection.
	Name *string `json:"name,omitempty"`
}

// NsVolumeCollectionSummaryFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolumeCollectionSummaryFieldHandles struct {
	ID   string
	Name string
}
