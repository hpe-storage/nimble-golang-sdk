// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsPoolSummary - Object containing pool ID and name.
// Export NsPoolSummaryFields for advance operations like search filter etc.
var NsPoolSummaryFields *NsPoolSummary

func init() {

	NsPoolSummaryFields = &NsPoolSummary{
		ID:   "id",
		Name: "name",
	}
}

type NsPoolSummary struct {
	// ID - ID of pool.
	ID string `json:"id,omitempty"`
	// Name - Name of pool.
	Name string `json:"name,omitempty"`
}

// sdk internal struct
type nsPoolSummary struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// EncodeNsPoolSummary - Transform NsPoolSummary to nsPoolSummary type
func EncodeNsPoolSummary(request interface{}) (*nsPoolSummary, error) {
	reqNsPoolSummary := request.(*NsPoolSummary)
	byte, err := json.Marshal(reqNsPoolSummary)
	if err != nil {
		return nil, err
	}
	respNsPoolSummaryPtr := &nsPoolSummary{}
	err = json.Unmarshal(byte, respNsPoolSummaryPtr)
	return respNsPoolSummaryPtr, err
}

// DecodeNsPoolSummary Transform nsPoolSummary to NsPoolSummary type
func DecodeNsPoolSummary(request interface{}) (*NsPoolSummary, error) {
	reqNsPoolSummary := request.(*nsPoolSummary)
	byte, err := json.Marshal(reqNsPoolSummary)
	if err != nil {
		return nil, err
	}
	respNsPoolSummary := &NsPoolSummary{}
	err = json.Unmarshal(byte, respNsPoolSummary)
	return respNsPoolSummary, err
}
