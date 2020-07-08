// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsCtrlrFcConfig - Controller Fibre Channel configuration.
// Export NsCtrlrFcConfigFields for advance operations like search filter etc.
var NsCtrlrFcConfigFields *NsCtrlrFcConfig

func init() {

	NsCtrlrFcConfigFields = &NsCtrlrFcConfig{}
}

type NsCtrlrFcConfig struct {
	// FcPortList - List of Fibre Channel ports.
	FcPortList []*NsFcPortInfo `json:"fc_port_list,omitempty"`
	// FcInterfaceList - List of Fibre Channel interfaces.
	FcInterfaceList []*NsFcInterfaceInfo `json:"fc_interface_list,omitempty"`
}

// sdk internal struct
type nsCtrlrFcConfig struct {
	// FcPortList - List of Fibre Channel ports.
	FcPortList []*NsFcPortInfo `json:"fc_port_list,omitempty"`
	// FcInterfaceList - List of Fibre Channel interfaces.
	FcInterfaceList []*NsFcInterfaceInfo `json:"fc_interface_list,omitempty"`
}

// EncodeNsCtrlrFcConfig - Transform NsCtrlrFcConfig to nsCtrlrFcConfig type
func EncodeNsCtrlrFcConfig(request interface{}) (*nsCtrlrFcConfig, error) {
	reqNsCtrlrFcConfig := request.(*NsCtrlrFcConfig)
	byte, err := json.Marshal(reqNsCtrlrFcConfig)
	objPtr := &nsCtrlrFcConfig{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsCtrlrFcConfig Transform nsCtrlrFcConfig to NsCtrlrFcConfig type
func DecodeNsCtrlrFcConfig(request interface{}) (*NsCtrlrFcConfig, error) {
	reqNsCtrlrFcConfig := request.(*nsCtrlrFcConfig)
	byte, err := json.Marshal(reqNsCtrlrFcConfig)
	obj := &NsCtrlrFcConfig{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
