// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	ID     *string `json:"id,omitempty"`
	Online *bool   `json:"online,omitempty"`
}

// EncodeNsFcInterfaceUpdateInfo - Transform NsFcInterfaceUpdateInfo to nsFcInterfaceUpdateInfo type
func EncodeNsFcInterfaceUpdateInfo(request interface{}) (*nsFcInterfaceUpdateInfo, error) {
	reqNsFcInterfaceUpdateInfo := request.(*NsFcInterfaceUpdateInfo)
	byte, err := json.Marshal(reqNsFcInterfaceUpdateInfo)
	if err != nil {
		return nil, err
	}
	respNsFcInterfaceUpdateInfoPtr := &nsFcInterfaceUpdateInfo{}
	err = json.Unmarshal(byte, respNsFcInterfaceUpdateInfoPtr)
	return respNsFcInterfaceUpdateInfoPtr, err
}

// DecodeNsFcInterfaceUpdateInfo Transform nsFcInterfaceUpdateInfo to NsFcInterfaceUpdateInfo type
func DecodeNsFcInterfaceUpdateInfo(request interface{}) (*NsFcInterfaceUpdateInfo, error) {
	reqNsFcInterfaceUpdateInfo := request.(*nsFcInterfaceUpdateInfo)
	byte, err := json.Marshal(reqNsFcInterfaceUpdateInfo)
	if err != nil {
		return nil, err
	}
	respNsFcInterfaceUpdateInfo := &NsFcInterfaceUpdateInfo{}
	err = json.Unmarshal(byte, respNsFcInterfaceUpdateInfo)
	return respNsFcInterfaceUpdateInfo, err
}
