// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsFcInterfaceUpdateInfo - Fibre Channel interface information to update.
// Export NsFcInterfaceUpdateInfoFields for advance operations like search filter etc.
var NsFcInterfaceUpdateInfoFields *NsFcInterfaceUpdateInfo

func init() {

	NsFcInterfaceUpdateInfoFields = &NsFcInterfaceUpdateInfo{
		ID: "id",
	}
}

type NsFcInterfaceUpdateInfo struct {
	// ID - ID of Fibre Channel interface.
	ID string `json:"id,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
}

// sdk internal struct
type nsFcInterfaceUpdateInfo struct {
	// ID - ID of Fibre Channel interface.
	ID *string `json:"id,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
}

// EncodeNsFcInterfaceUpdateInfo - Transform NsFcInterfaceUpdateInfo to nsFcInterfaceUpdateInfo type
func EncodeNsFcInterfaceUpdateInfo(request interface{}) (*nsFcInterfaceUpdateInfo, error) {
	reqNsFcInterfaceUpdateInfo := request.(*NsFcInterfaceUpdateInfo)
	byte, err := json.Marshal(reqNsFcInterfaceUpdateInfo)
	objPtr := &nsFcInterfaceUpdateInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFcInterfaceUpdateInfo Transform nsFcInterfaceUpdateInfo to NsFcInterfaceUpdateInfo type
func DecodeNsFcInterfaceUpdateInfo(request interface{}) (*NsFcInterfaceUpdateInfo, error) {
	reqNsFcInterfaceUpdateInfo := request.(*nsFcInterfaceUpdateInfo)
	byte, err := json.Marshal(reqNsFcInterfaceUpdateInfo)
	obj := &NsFcInterfaceUpdateInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
