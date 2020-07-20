// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsPerfPolicySummary - Select fields containing performance policy.
// Export NsPerfPolicySummaryFields for advance operations like search filter etc.
var NsPerfPolicySummaryFields *NsPerfPolicySummary

func init() {

	NsPerfPolicySummaryFields = &NsPerfPolicySummary{
		Name: "name",
	}
}

type NsPerfPolicySummary struct {
	// Name - Name of performance policy.
	Name string `json:"name,omitempty"`
}

// sdk internal struct
type nsPerfPolicySummary struct {
	Name *string `json:"name,omitempty"`
}

// EncodeNsPerfPolicySummary - Transform NsPerfPolicySummary to nsPerfPolicySummary type
func EncodeNsPerfPolicySummary(request interface{}) (*nsPerfPolicySummary, error) {
	reqNsPerfPolicySummary := request.(*NsPerfPolicySummary)
	byte, err := json.Marshal(reqNsPerfPolicySummary)
	if err != nil {
		return nil, err
	}
	respNsPerfPolicySummaryPtr := &nsPerfPolicySummary{}
	err = json.Unmarshal(byte, respNsPerfPolicySummaryPtr)
	return respNsPerfPolicySummaryPtr, err
}

// DecodeNsPerfPolicySummary Transform nsPerfPolicySummary to NsPerfPolicySummary type
func DecodeNsPerfPolicySummary(request interface{}) (*NsPerfPolicySummary, error) {
	reqNsPerfPolicySummary := request.(*nsPerfPolicySummary)
	byte, err := json.Marshal(reqNsPerfPolicySummary)
	if err != nil {
		return nil, err
	}
	respNsPerfPolicySummary := &NsPerfPolicySummary{}
	err = json.Unmarshal(byte, respNsPerfPolicySummary)
	return respNsPerfPolicySummary, err
}
