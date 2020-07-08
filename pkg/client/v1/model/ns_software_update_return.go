// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSoftwareUpdateReturn - The status returned by the software update precheck and start actions.
// Export NsSoftwareUpdateReturnFields for advance operations like search filter etc.
var NsSoftwareUpdateReturnFields *NsSoftwareUpdateReturn

func init() {

	NsSoftwareUpdateReturnFields = &NsSoftwareUpdateReturn{
		Error: "error",
	}
}

type NsSoftwareUpdateReturn struct {
	// Error - Top level error.
	Error string `json:"error,omitempty"`
	// ArrayResponseList - Errors from all the arrays in the group.
	ArrayResponseList []*NsArraySoftwareUpdateStatus `json:"array_response_list,omitempty"`
}

// sdk internal struct
type nsSoftwareUpdateReturn struct {
	// Error - Top level error.
	Error *string `json:"error,omitempty"`
	// ArrayResponseList - Errors from all the arrays in the group.
	ArrayResponseList []*NsArraySoftwareUpdateStatus `json:"array_response_list,omitempty"`
}

// EncodeNsSoftwareUpdateReturn - Transform NsSoftwareUpdateReturn to nsSoftwareUpdateReturn type
func EncodeNsSoftwareUpdateReturn(request interface{}) (*nsSoftwareUpdateReturn, error) {
	reqNsSoftwareUpdateReturn := request.(*NsSoftwareUpdateReturn)
	byte, err := json.Marshal(reqNsSoftwareUpdateReturn)
	objPtr := &nsSoftwareUpdateReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSoftwareUpdateReturn Transform nsSoftwareUpdateReturn to NsSoftwareUpdateReturn type
func DecodeNsSoftwareUpdateReturn(request interface{}) (*NsSoftwareUpdateReturn, error) {
	reqNsSoftwareUpdateReturn := request.(*nsSoftwareUpdateReturn)
	byte, err := json.Marshal(reqNsSoftwareUpdateReturn)
	obj := &NsSoftwareUpdateReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
