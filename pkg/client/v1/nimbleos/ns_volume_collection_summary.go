// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolumeCollectionSummary - Select fields of volume collection info.
// Export NsVolumeCollectionSummaryFields for advance operations like search filter etc.
var NsVolumeCollectionSummaryFields *NsVolumeCollectionSummary

func init() {

	NsVolumeCollectionSummaryFields = &NsVolumeCollectionSummary{
		ID:   "id",
		Name: "name",
	}
}

type NsVolumeCollectionSummary struct {
	// ID - Identifier of volume collection.
	ID string `json:"id,omitempty"`
	// Name - Name of volume collection.
	Name string `json:"name,omitempty"`
}

// sdk internal struct
type nsVolumeCollectionSummary struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// EncodeNsVolumeCollectionSummary - Transform NsVolumeCollectionSummary to nsVolumeCollectionSummary type
func EncodeNsVolumeCollectionSummary(request interface{}) (*nsVolumeCollectionSummary, error) {
	reqNsVolumeCollectionSummary := request.(*NsVolumeCollectionSummary)
	byte, err := json.Marshal(reqNsVolumeCollectionSummary)
	if err != nil {
		return nil, err
	}
	respNsVolumeCollectionSummaryPtr := &nsVolumeCollectionSummary{}
	err = json.Unmarshal(byte, respNsVolumeCollectionSummaryPtr)
	return respNsVolumeCollectionSummaryPtr, err
}

// DecodeNsVolumeCollectionSummary Transform nsVolumeCollectionSummary to NsVolumeCollectionSummary type
func DecodeNsVolumeCollectionSummary(request interface{}) (*NsVolumeCollectionSummary, error) {
	reqNsVolumeCollectionSummary := request.(*nsVolumeCollectionSummary)
	byte, err := json.Marshal(reqNsVolumeCollectionSummary)
	if err != nil {
		return nil, err
	}
	respNsVolumeCollectionSummary := &NsVolumeCollectionSummary{}
	err = json.Unmarshal(byte, respNsVolumeCollectionSummary)
	return respNsVolumeCollectionSummary, err
}
