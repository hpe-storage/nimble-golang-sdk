// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFCSession - Fibre Channel initiator session information.
// Export NsFCSessionFields for advance operations like search filter etc.
var NsFCSessionFields *NsFCSession

func init() {

	NsFCSessionFields = &NsFCSession{
		ID:                        "id",
		SessionId:                 "session_id",
		InitiatorAlias:            "initiator_alias",
		InitiatorWwpn:             "initiator_wwpn",
		InitiatorWwnn:             "initiator_wwnn",
		InitiatorSwitchName:       "initiator_switch_name",
		InitiatorSwitchPort:       "initiator_switch_port",
		InitiatorSymbolicPortname: "initiator_symbolic_portname",
		InitiatorSymbolicNodename: "initiator_symbolic_nodename",
		TargetPortArrayName:       "target_port_array_name",
		TargetPortInterfaceName:   "target_port_interface_name",
		TargetWwnn:                "target_wwnn",
		TargetWwpn:                "target_wwpn",
	}
}

type NsFCSession struct {
	// ID - Unique identifier of the Fibre Channel session.
	ID string `json:"id,omitempty"`
	// SessionId - Unique identifier of the Fibre Channel session.
	SessionId string `json:"session_id,omitempty"`
	// Alua - ALUA (Asymmetric Logical Unit Access) value of the Fibre Channel session.
	Alua *NsALUA `json:"alua,omitempty"`
	// PrKey - Registered persistent reservation key for this session on associated logical unit (if there is no registration, its value will be zero).
	PrKey int64 `json:"pr_key,omitempty"`
	// InitiatorAlias - Alias of the Fibre Channel initiator.
	InitiatorAlias string `json:"initiator_alias,omitempty"`
	// InitiatorWwpn - WWPN (World Wide Port Name) of the Fibre Channel initiator.
	InitiatorWwpn string `json:"initiator_wwpn,omitempty"`
	// InitiatorWwnn - WWNN (World Wide Node Name) of the Fibre Channel initiator.
	InitiatorWwnn string `json:"initiator_wwnn,omitempty"`
	// InitiatorSwitchName - Name of the switch used by the Fibre Channel initiator.
	InitiatorSwitchName string `json:"initiator_switch_name,omitempty"`
	// InitiatorSwitchPort - Port on the switch used by the Fibre Channel initiator.
	InitiatorSwitchPort string `json:"initiator_switch_port,omitempty"`
	// InitiatorSymbolicPortname - Symbolic port name associated with the Fibre Channel initiator.
	InitiatorSymbolicPortname string `json:"initiator_symbolic_portname,omitempty"`
	// InitiatorSymbolicNodename - Symbolic node name associated with the Fibre Channel initiator.
	InitiatorSymbolicNodename string `json:"initiator_symbolic_nodename,omitempty"`
	// InitiatorFcid - FCID assigned to the Fibre Channel initiator.
	InitiatorFcid int64 `json:"initiator_fcid,omitempty"`
	// TargetPortArrayName - Name of the array hosting the Fibre Channel target port.
	TargetPortArrayName string `json:"target_port_array_name,omitempty"`
	// TargetPortCtrlrId - Identifier of the controller hosting the Fibre Channel target port.
	TargetPortCtrlrId int64 `json:"target_port_ctrlr_id,omitempty"`
	// TargetPortInterfaceName - Name of the interface hosted on the Fibre Channel target port.
	TargetPortInterfaceName string `json:"target_port_interface_name,omitempty"`
	// TargetWwnn - WWNN (World Wide Node Name) of the Fibre Channel target port.
	TargetWwnn string `json:"target_wwnn,omitempty"`
	// TargetWwpn - WWPN (World Wide Port Name) of the Fibre Channel target port.
	TargetWwpn string `json:"target_wwpn,omitempty"`
	// TargetFcid - FCID assigned to the Fibre Channel target port.
	TargetFcid int64 `json:"target_fcid,omitempty"`
}

// sdk internal struct
type nsFCSession struct {
	// ID - Unique identifier of the Fibre Channel session.
	ID *string `json:"id,omitempty"`
	// SessionId - Unique identifier of the Fibre Channel session.
	SessionId *string `json:"session_id,omitempty"`
	// Alua - ALUA (Asymmetric Logical Unit Access) value of the Fibre Channel session.
	Alua *NsALUA `json:"alua,omitempty"`
	// PrKey - Registered persistent reservation key for this session on associated logical unit (if there is no registration, its value will be zero).
	PrKey *int64 `json:"pr_key,omitempty"`
	// InitiatorAlias - Alias of the Fibre Channel initiator.
	InitiatorAlias *string `json:"initiator_alias,omitempty"`
	// InitiatorWwpn - WWPN (World Wide Port Name) of the Fibre Channel initiator.
	InitiatorWwpn *string `json:"initiator_wwpn,omitempty"`
	// InitiatorWwnn - WWNN (World Wide Node Name) of the Fibre Channel initiator.
	InitiatorWwnn *string `json:"initiator_wwnn,omitempty"`
	// InitiatorSwitchName - Name of the switch used by the Fibre Channel initiator.
	InitiatorSwitchName *string `json:"initiator_switch_name,omitempty"`
	// InitiatorSwitchPort - Port on the switch used by the Fibre Channel initiator.
	InitiatorSwitchPort *string `json:"initiator_switch_port,omitempty"`
	// InitiatorSymbolicPortname - Symbolic port name associated with the Fibre Channel initiator.
	InitiatorSymbolicPortname *string `json:"initiator_symbolic_portname,omitempty"`
	// InitiatorSymbolicNodename - Symbolic node name associated with the Fibre Channel initiator.
	InitiatorSymbolicNodename *string `json:"initiator_symbolic_nodename,omitempty"`
	// InitiatorFcid - FCID assigned to the Fibre Channel initiator.
	InitiatorFcid *int64 `json:"initiator_fcid,omitempty"`
	// TargetPortArrayName - Name of the array hosting the Fibre Channel target port.
	TargetPortArrayName *string `json:"target_port_array_name,omitempty"`
	// TargetPortCtrlrId - Identifier of the controller hosting the Fibre Channel target port.
	TargetPortCtrlrId *int64 `json:"target_port_ctrlr_id,omitempty"`
	// TargetPortInterfaceName - Name of the interface hosted on the Fibre Channel target port.
	TargetPortInterfaceName *string `json:"target_port_interface_name,omitempty"`
	// TargetWwnn - WWNN (World Wide Node Name) of the Fibre Channel target port.
	TargetWwnn *string `json:"target_wwnn,omitempty"`
	// TargetWwpn - WWPN (World Wide Port Name) of the Fibre Channel target port.
	TargetWwpn *string `json:"target_wwpn,omitempty"`
	// TargetFcid - FCID assigned to the Fibre Channel target port.
	TargetFcid *int64 `json:"target_fcid,omitempty"`
}

// EncodeNsFCSession - Transform NsFCSession to nsFCSession type
func EncodeNsFCSession(request interface{}) (*nsFCSession, error) {
	reqNsFCSession := request.(*NsFCSession)
	byte, err := json.Marshal(reqNsFCSession)
	objPtr := &nsFCSession{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFCSession Transform nsFCSession to NsFCSession type
func DecodeNsFCSession(request interface{}) (*NsFCSession, error) {
	reqNsFCSession := request.(*nsFCSession)
	byte, err := json.Marshal(reqNsFCSession)
	obj := &NsFCSession{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
