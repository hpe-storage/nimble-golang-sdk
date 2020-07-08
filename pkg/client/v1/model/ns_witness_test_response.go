// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsWitnessTestResponse - Results of witness connection test.
// Export NsWitnessTestResponseFields for advance operations like search filter etc.
var NsWitnessTestResponseFields *NsWitnessTestResponse

func init() {

	NsWitnessTestResponseFields = &NsWitnessTestResponse{
		ArrayName:                  "array_name",
		WitnessConnectivityState:   "witness_connectivity_state",
		WitnessConnectivityMessage: "witness_connectivity_message",
	}
}

type NsWitnessTestResponse struct {
	// ArrayName - Name of an array.
	ArrayName string `json:"array_name,omitempty"`
	// Role - Role of an array in the group.
	Role *NsArrayRole `json:"role,omitempty"`
	// WitnessConnectivityState - Reachability status of the witness.
	WitnessConnectivityState string `json:"witness_connectivity_state,omitempty"`
	// WitnessConnectivityMessage - Reachability message of the witness.
	WitnessConnectivityMessage string `json:"witness_connectivity_message,omitempty"`
}

// sdk internal struct
type nsWitnessTestResponse struct {
	// ArrayName - Name of an array.
	ArrayName *string `json:"array_name,omitempty"`
	// Role - Role of an array in the group.
	Role *NsArrayRole `json:"role,omitempty"`
	// WitnessConnectivityState - Reachability status of the witness.
	WitnessConnectivityState *string `json:"witness_connectivity_state,omitempty"`
	// WitnessConnectivityMessage - Reachability message of the witness.
	WitnessConnectivityMessage *string `json:"witness_connectivity_message,omitempty"`
}

// EncodeNsWitnessTestResponse - Transform NsWitnessTestResponse to nsWitnessTestResponse type
func EncodeNsWitnessTestResponse(request interface{}) (*nsWitnessTestResponse, error) {
	reqNsWitnessTestResponse := request.(*NsWitnessTestResponse)
	byte, err := json.Marshal(reqNsWitnessTestResponse)
	objPtr := &nsWitnessTestResponse{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsWitnessTestResponse Transform nsWitnessTestResponse to NsWitnessTestResponse type
func DecodeNsWitnessTestResponse(request interface{}) (*NsWitnessTestResponse, error) {
	reqNsWitnessTestResponse := request.(*nsWitnessTestResponse)
	byte, err := json.Marshal(reqNsWitnessTestResponse)
	obj := &NsWitnessTestResponse{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
