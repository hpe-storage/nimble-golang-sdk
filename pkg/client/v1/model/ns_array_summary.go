// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsArraySummary - Array summary information.
// Export NsArraySummaryFields for advance operations like search filter etc.
var NsArraySummaryFields *NsArraySummary

func init() {

	NsArraySummaryFields = &NsArraySummary{
		ID:        "id",
		ArrayId:   "array_id",
		Name:      "name",
		ArrayName: "array_name",
	}
}

type NsArraySummary struct {
	// ID - Array API ID.
	ID string `json:"id,omitempty"`
	// ArrayId - Array API ID.
	ArrayId string `json:"array_id,omitempty"`
	// Name - Unique name of array.
	Name string `json:"name,omitempty"`
	// ArrayName - Unique name of array.
	ArrayName string `json:"array_name,omitempty"`
}

// sdk internal struct
type nsArraySummary struct {
	// ID - Array API ID.
	ID *string `json:"id,omitempty"`
	// ArrayId - Array API ID.
	ArrayId *string `json:"array_id,omitempty"`
	// Name - Unique name of array.
	Name *string `json:"name,omitempty"`
	// ArrayName - Unique name of array.
	ArrayName *string `json:"array_name,omitempty"`
}

// EncodeNsArraySummary - Transform NsArraySummary to nsArraySummary type
func EncodeNsArraySummary(request interface{}) (*nsArraySummary, error) {
	reqNsArraySummary := request.(*NsArraySummary)
	byte, err := json.Marshal(reqNsArraySummary)
	objPtr := &nsArraySummary{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsArraySummary Transform nsArraySummary to NsArraySummary type
func DecodeNsArraySummary(request interface{}) (*NsArraySummary, error) {
	reqNsArraySummary := request.(*nsArraySummary)
	byte, err := json.Marshal(reqNsArraySummary)
	obj := &NsArraySummary{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
