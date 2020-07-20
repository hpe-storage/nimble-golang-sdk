// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	FcPortList      []*NsFcPortInfo      `json:"fc_port_list,omitempty"`
	FcInterfaceList []*NsFcInterfaceInfo `json:"fc_interface_list,omitempty"`
}

// EncodeNsCtrlrFcConfig - Transform NsCtrlrFcConfig to nsCtrlrFcConfig type
func EncodeNsCtrlrFcConfig(request interface{}) (*nsCtrlrFcConfig, error) {
	reqNsCtrlrFcConfig := request.(*NsCtrlrFcConfig)
	byte, err := json.Marshal(reqNsCtrlrFcConfig)
	if err != nil {
		return nil, err
	}
	respNsCtrlrFcConfigPtr := &nsCtrlrFcConfig{}
	err = json.Unmarshal(byte, respNsCtrlrFcConfigPtr)
	return respNsCtrlrFcConfigPtr, err
}

// DecodeNsCtrlrFcConfig Transform nsCtrlrFcConfig to NsCtrlrFcConfig type
func DecodeNsCtrlrFcConfig(request interface{}) (*NsCtrlrFcConfig, error) {
	reqNsCtrlrFcConfig := request.(*nsCtrlrFcConfig)
	byte, err := json.Marshal(reqNsCtrlrFcConfig)
	if err != nil {
		return nil, err
	}
	respNsCtrlrFcConfig := &NsCtrlrFcConfig{}
	err = json.Unmarshal(byte, respNsCtrlrFcConfig)
	return respNsCtrlrFcConfig, err
}
