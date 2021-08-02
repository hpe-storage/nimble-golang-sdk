// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFCSessionFields provides field names to use in filter parameters, for example.
var NsFCSessionFields *NsFCSessionFieldHandles

func init() {
	NsFCSessionFields = &NsFCSessionFieldHandles{
		ID:                        "id",
		SessionId:                 "session_id",
		Alua:                      "alua",
		PrKey:                     "pr_key",
		InitiatorAlias:            "initiator_alias",
		InitiatorWwpn:             "initiator_wwpn",
		InitiatorWwnn:             "initiator_wwnn",
		InitiatorSwitchName:       "initiator_switch_name",
		InitiatorSwitchPort:       "initiator_switch_port",
		InitiatorSymbolicPortname: "initiator_symbolic_portname",
		InitiatorSymbolicNodename: "initiator_symbolic_nodename",
		InitiatorFcid:             "initiator_fcid",
		TargetPortArrayName:       "target_port_array_name",
		TargetPortCtrlrId:         "target_port_ctrlr_id",
		TargetPortInterfaceName:   "target_port_interface_name",
		TargetWwnn:                "target_wwnn",
		TargetWwpn:                "target_wwpn",
		TargetFcid:                "target_fcid",
	}
}

// NsFCSession - Fibre Channel initiator session information.
type NsFCSession struct {
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

// NsFCSessionFieldHandles provides a string representation for each NsFCSession field.
type NsFCSessionFieldHandles struct {
	ID                        string
	SessionId                 string
	Alua                      string
	PrKey                     string
	InitiatorAlias            string
	InitiatorWwpn             string
	InitiatorWwnn             string
	InitiatorSwitchName       string
	InitiatorSwitchPort       string
	InitiatorSymbolicPortname string
	InitiatorSymbolicNodename string
	InitiatorFcid             string
	TargetPortArrayName       string
	TargetPortCtrlrId         string
	TargetPortInterfaceName   string
	TargetWwnn                string
	TargetWwpn                string
	TargetFcid                string
}
