// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

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
	// ID - ID of pool.
	ID *string `json:"id,omitempty"`
	// Name - Name of pool.
	Name *string `json:"name,omitempty"`
}

// EncodeNsPoolSummary - Transform NsPoolSummary to nsPoolSummary type
func EncodeNsPoolSummary(request interface{}) (*nsPoolSummary, error) {
	reqNsPoolSummary := request.(*NsPoolSummary)
	byte, err := json.Marshal(reqNsPoolSummary)
	objPtr := &nsPoolSummary{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsPoolSummary Transform nsPoolSummary to NsPoolSummary type
func DecodeNsPoolSummary(request interface{}) (*NsPoolSummary, error) {
	reqNsPoolSummary := request.(*nsPoolSummary)
	byte, err := json.Marshal(reqNsPoolSummary)
	obj := &NsPoolSummary{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
