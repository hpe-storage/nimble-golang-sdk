// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrFcConfig - Controller Fibre Channel configuration.
// Export NsCtrlrFcConfigFields for advance operations like search filter etc.
var NsCtrlrFcConfigFields *NsCtrlrFcConfigStringFields

func init() {
	FcPortListfield := "fc_port_list"
	FcInterfaceListfield := "fc_interface_list"

	NsCtrlrFcConfigFields = &NsCtrlrFcConfigStringFields{
		FcPortList:      &FcPortListfield,
		FcInterfaceList: &FcInterfaceListfield,
	}
}

type NsCtrlrFcConfig struct {
	// FcPortList - List of Fibre Channel ports.
	FcPortList []*NsFcPortInfo `json:"fc_port_list,omitempty"`
	// FcInterfaceList - List of Fibre Channel interfaces.
	FcInterfaceList []*NsFcInterfaceInfo `json:"fc_interface_list,omitempty"`
}

// Struct for NsCtrlrFcConfigFields
type NsCtrlrFcConfigStringFields struct {
	FcPortList      *string
	FcInterfaceList *string
}
