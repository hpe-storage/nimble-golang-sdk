// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFibreChannelInterfaceFullName - Fully qualified name information for a Fibre Channel interface (array/controller/interface).
// Export NsFibreChannelInterfaceFullNameFields for advance operations like search filter etc.
var NsFibreChannelInterfaceFullNameFields *NsFibreChannelInterfaceFullName

func init() {

	NsFibreChannelInterfaceFullNameFields = &NsFibreChannelInterfaceFullName{
		ArrayName: "array_name",
		CtrlrName: "ctrlr_name",
		IntfName:  "intf_name",
	}
}

type NsFibreChannelInterfaceFullName struct {
	// ArrayName - Array name.
	ArrayName string `json:"array_name,omitempty"`
	// CtrlrName - Controller name.
	CtrlrName string `json:"ctrlr_name,omitempty"`
	// IntfName - Fibre Channel interface name.
	IntfName string `json:"intf_name,omitempty"`
}

// sdk internal struct
type nsFibreChannelInterfaceFullName struct {
	ArrayName *string `json:"array_name,omitempty"`
	CtrlrName *string `json:"ctrlr_name,omitempty"`
	IntfName  *string `json:"intf_name,omitempty"`
}

// EncodeNsFibreChannelInterfaceFullName - Transform NsFibreChannelInterfaceFullName to nsFibreChannelInterfaceFullName type
func EncodeNsFibreChannelInterfaceFullName(request interface{}) (*nsFibreChannelInterfaceFullName, error) {
	reqNsFibreChannelInterfaceFullName := request.(*NsFibreChannelInterfaceFullName)
	byte, err := json.Marshal(reqNsFibreChannelInterfaceFullName)
	if err != nil {
		return nil, err
	}
	respNsFibreChannelInterfaceFullNamePtr := &nsFibreChannelInterfaceFullName{}
	err = json.Unmarshal(byte, respNsFibreChannelInterfaceFullNamePtr)
	return respNsFibreChannelInterfaceFullNamePtr, err
}

// DecodeNsFibreChannelInterfaceFullName Transform nsFibreChannelInterfaceFullName to NsFibreChannelInterfaceFullName type
func DecodeNsFibreChannelInterfaceFullName(request interface{}) (*NsFibreChannelInterfaceFullName, error) {
	reqNsFibreChannelInterfaceFullName := request.(*nsFibreChannelInterfaceFullName)
	byte, err := json.Marshal(reqNsFibreChannelInterfaceFullName)
	if err != nil {
		return nil, err
	}
	respNsFibreChannelInterfaceFullName := &NsFibreChannelInterfaceFullName{}
	err = json.Unmarshal(byte, respNsFibreChannelInterfaceFullName)
	return respNsFibreChannelInterfaceFullName, err
}
