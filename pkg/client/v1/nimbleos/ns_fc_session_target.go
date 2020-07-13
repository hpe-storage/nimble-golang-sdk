// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFcSessionTarget - Information of the Fibre Channel session target.
// Export NsFcSessionTargetFields for advance operations like search filter etc.
var NsFcSessionTargetFields *NsFcSessionTarget

func init() {

	NsFcSessionTargetFields = &NsFcSessionTarget{
		TargetPortArrayName:     "target_port_array_name",
		TargetPortCtrlrName:     "target_port_ctrlr_name",
		TargetPortInterfaceName: "target_port_interface_name",
		TargetWwnn:              "target_wwnn",
		TargetWwpn:              "target_wwpn",
		TargetFcid:              "target_fcid",
	}
}

type NsFcSessionTarget struct {
	// TargetPortArrayName - Name of the array hosting the Fibre Channel target port.
	TargetPortArrayName string `json:"target_port_array_name,omitempty"`
	// TargetPortCtrlrName - Name (A or B) of the controller to which the port belongs.
	TargetPortCtrlrName string `json:"target_port_ctrlr_name,omitempty"`
	// TargetPortInterfaceName - Name of the interface hosted on the Fibre Channel target port.
	TargetPortInterfaceName string `json:"target_port_interface_name,omitempty"`
	// TargetWwnn - WWNN (World Wide Node Name) of the Fibre Channel target port.
	TargetWwnn string `json:"target_wwnn,omitempty"`
	// TargetWwpn - WWPN (World Wide Port Name) of the Fibre Channel target port.
	TargetWwpn string `json:"target_wwpn,omitempty"`
	// TargetFcid - FCID assigned to the Fibre Channel target port.
	TargetFcid string `json:"target_fcid,omitempty"`
}

// sdk internal struct
type nsFcSessionTarget struct {
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

// EncodeNsFcSessionTarget - Transform NsFcSessionTarget to nsFcSessionTarget type
func EncodeNsFcSessionTarget(request interface{}) (*nsFcSessionTarget, error) {
	reqNsFcSessionTarget := request.(*NsFcSessionTarget)
	byte, err := json.Marshal(reqNsFcSessionTarget)
	objPtr := &nsFcSessionTarget{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFcSessionTarget Transform nsFcSessionTarget to NsFcSessionTarget type
func DecodeNsFcSessionTarget(request interface{}) (*NsFcSessionTarget, error) {
	reqNsFcSessionTarget := request.(*nsFcSessionTarget)
	byte, err := json.Marshal(reqNsFcSessionTarget)
	obj := &NsFcSessionTarget{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
