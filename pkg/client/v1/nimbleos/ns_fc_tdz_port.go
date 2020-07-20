// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFcTdzPort - Fibre Channel Target port.
// Export NsFcTdzPortFields for advance operations like search filter etc.
var NsFcTdzPortFields *NsFcTdzPort

func init() {

	NsFcTdzPortFields = &NsFcTdzPort{
		ArrayName: "array_name",
		FcName:    "fc_name",
	}
}

type NsFcTdzPort struct {
	// ArrayName - Unique name of the array.
	ArrayName string `json:"array_name,omitempty"`
	// FcName - Target port interface name.
	FcName string `json:"fc_name,omitempty"`
}

// sdk internal struct
type nsFcTdzPort struct {
	ArrayName *string `json:"array_name,omitempty"`
	FcName    *string `json:"fc_name,omitempty"`
}

// EncodeNsFcTdzPort - Transform NsFcTdzPort to nsFcTdzPort type
func EncodeNsFcTdzPort(request interface{}) (*nsFcTdzPort, error) {
	reqNsFcTdzPort := request.(*NsFcTdzPort)
	byte, err := json.Marshal(reqNsFcTdzPort)
	if err != nil {
		return nil, err
	}
	respNsFcTdzPortPtr := &nsFcTdzPort{}
	err = json.Unmarshal(byte, respNsFcTdzPortPtr)
	return respNsFcTdzPortPtr, err
}

// DecodeNsFcTdzPort Transform nsFcTdzPort to NsFcTdzPort type
func DecodeNsFcTdzPort(request interface{}) (*NsFcTdzPort, error) {
	reqNsFcTdzPort := request.(*nsFcTdzPort)
	byte, err := json.Marshal(reqNsFcTdzPort)
	if err != nil {
		return nil, err
	}
	respNsFcTdzPort := &NsFcTdzPort{}
	err = json.Unmarshal(byte, respNsFcTdzPort)
	return respNsFcTdzPort, err
}
