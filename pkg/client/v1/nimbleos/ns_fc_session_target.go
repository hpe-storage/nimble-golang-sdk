// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFcSessionTargetFields provides field names to use in filter parameters, for example.
var NsFcSessionTargetFields *NsFcSessionTargetFieldHandles

func init() {
	fieldTargetPortArrayName := "target_port_array_name"
	fieldTargetPortCtrlrName := "target_port_ctrlr_name"
	fieldTargetPortInterfaceName := "target_port_interface_name"
	fieldTargetWwnn := "target_wwnn"
	fieldTargetWwpn := "target_wwpn"
	fieldTargetFcid := "target_fcid"

	NsFcSessionTargetFields = &NsFcSessionTargetFieldHandles{
		TargetPortArrayName:     &fieldTargetPortArrayName,
		TargetPortCtrlrName:     &fieldTargetPortCtrlrName,
		TargetPortInterfaceName: &fieldTargetPortInterfaceName,
		TargetWwnn:              &fieldTargetWwnn,
		TargetWwpn:              &fieldTargetWwpn,
		TargetFcid:              &fieldTargetFcid,
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

// NsFcSessionTargetFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcSessionTargetFieldHandles struct {
	TargetPortArrayName     *string
	TargetPortCtrlrName     *string
	TargetPortInterfaceName *string
	TargetWwnn              *string
	TargetWwpn              *string
	TargetFcid              *string
}
