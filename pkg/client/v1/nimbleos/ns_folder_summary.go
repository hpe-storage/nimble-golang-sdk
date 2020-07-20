// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	ID  *string `json:"id,omitempty"`
	Fqn *string `json:"fqn,omitempty"`
}

// EncodeNsFolderSummary - Transform NsFolderSummary to nsFolderSummary type
func EncodeNsFolderSummary(request interface{}) (*nsFolderSummary, error) {
	reqNsFolderSummary := request.(*NsFolderSummary)
	byte, err := json.Marshal(reqNsFolderSummary)
	if err != nil {
		return nil, err
	}
	respNsFolderSummaryPtr := &nsFolderSummary{}
	err = json.Unmarshal(byte, respNsFolderSummaryPtr)
	return respNsFolderSummaryPtr, err
}

// DecodeNsFolderSummary Transform nsFolderSummary to NsFolderSummary type
func DecodeNsFolderSummary(request interface{}) (*NsFolderSummary, error) {
	reqNsFolderSummary := request.(*nsFolderSummary)
	byte, err := json.Marshal(reqNsFolderSummary)
	if err != nil {
		return nil, err
	}
	respNsFolderSummary := &NsFolderSummary{}
	err = json.Unmarshal(byte, respNsFolderSummary)
	return respNsFolderSummary, err
}
