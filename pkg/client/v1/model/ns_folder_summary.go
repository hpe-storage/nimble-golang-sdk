// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsFolderSummary - Select fields containing folder info.
// Export NsFolderSummaryFields for advance operations like search filter etc.
var NsFolderSummaryFields *NsFolderSummary

func init() {

	NsFolderSummaryFields = &NsFolderSummary{
		ID:  "id",
		Fqn: "fqn",
	}
}

type NsFolderSummary struct {
	// ID - ID of folder.
	ID string `json:"id,omitempty"`
	// Fqn - Fully qualified name of folder.
	Fqn string `json:"fqn,omitempty"`
}

// sdk internal struct
type nsFolderSummary struct {
	// ID - ID of folder.
	ID *string `json:"id,omitempty"`
	// Fqn - Fully qualified name of folder.
	Fqn *string `json:"fqn,omitempty"`
}

// EncodeNsFolderSummary - Transform NsFolderSummary to nsFolderSummary type
func EncodeNsFolderSummary(request interface{}) (*nsFolderSummary, error) {
	reqNsFolderSummary := request.(*NsFolderSummary)
	byte, err := json.Marshal(reqNsFolderSummary)
	objPtr := &nsFolderSummary{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFolderSummary Transform nsFolderSummary to NsFolderSummary type
func DecodeNsFolderSummary(request interface{}) (*NsFolderSummary, error) {
	reqNsFolderSummary := request.(*nsFolderSummary)
	byte, err := json.Marshal(reqNsFolderSummary)
	obj := &NsFolderSummary{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
