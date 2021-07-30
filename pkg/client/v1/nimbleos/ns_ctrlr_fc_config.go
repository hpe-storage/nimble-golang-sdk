// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrFcConfig - Controller Fibre Channel configuration.

// Export NsCtrlrFcConfigFields provides field names to use in filter parameters, for example.
var NsCtrlrFcConfigFields *NsCtrlrFcConfigStringFields

func init() {
	fieldFcPortList := "fc_port_list"
	fieldFcInterfaceList := "fc_interface_list"

	NsCtrlrFcConfigFields = &NsCtrlrFcConfigStringFields{
		FcPortList:      &fieldFcPortList,
		FcInterfaceList: &fieldFcInterfaceList,
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
