// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeCollectionSummary - Select fields of volume collection info.
// Export NsVolumeCollectionSummaryFields for advance operations like search filter etc.
var NsVolumeCollectionSummaryFields *NsVolumeCollectionSummary

func init() {
	IDfield := "id"
	Namefield := "name"

	NsVolumeCollectionSummaryFields = &NsVolumeCollectionSummary{
		ID:   &IDfield,
		Name: &Namefield,
	}
}

type NsVolumeCollectionSummary struct {
	// ID - Identifier of volume collection.
	ID *string `json:"id,omitempty"`
	// Name - Name of volume collection.
	Name *string `json:"name,omitempty"`
}
