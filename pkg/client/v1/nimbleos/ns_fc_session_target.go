// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcSessionTargetFields provides field names to use in filter parameters, for example.
var NsFcSessionTargetFields *NsFcSessionTargetFieldHandles

func init() {
	NsFcSessionTargetFields = &NsFcSessionTargetFieldHandles{
		TargetPortArrayName:     "target_port_array_name",
		TargetPortCtrlrName:     "target_port_ctrlr_name",
		TargetPortInterfaceName: "target_port_interface_name",
		TargetWwnn:              "target_wwnn",
		TargetWwpn:              "target_wwpn",
		TargetFcid:              "target_fcid",
	}
}

// NsFcSessionTarget - Information of the Fibre Channel session target.
type NsFcSessionTarget struct {
	// TargetPortArrayName - Name of the array hosting the Fibre Channel target port.
	TargetPortArrayName *string `json:"target_port_array_name,omitempty"`
	// TargetPortCtrlrName - Name (A or B) of the controller to which the port belongs.
	TargetPortCtrlrName *string `json:"target_port_ctrlr_name,omitempty"`
	// TargetPortInterfaceName - Name of the interface hosted on the Fibre Channel target port.
	TargetPortInterfaceName *string `json:"target_port_interface_name,omitempty"`
	// TargetWwnn - WWNN (World Wide Node Name) of the Fibre Channel target port.
	TargetWwnn *string `json:"target_wwnn,omitempty"`
	// TargetWwpn - WWPN (World Wide Port Name) of the Fibre Channel target port.
	TargetWwpn *string `json:"target_wwpn,omitempty"`
	// TargetFcid - FCID assigned to the Fibre Channel target port.
	TargetFcid *string `json:"target_fcid,omitempty"`
}

// NsFcSessionTargetFieldHandles provides a string representation for each NsFcSessionTarget field.
type NsFcSessionTargetFieldHandles struct {
	TargetPortArrayName     string
	TargetPortCtrlrName     string
	TargetPortInterfaceName string
	TargetWwnn              string
	TargetWwpn              string
	TargetFcid              string
}
