// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsFcSessionInitiatorFields provides field names to use in filter parameters, for example.
var NsFcSessionInitiatorFields *NsFcSessionInitiatorFieldHandles

func init() {
	fieldInitiatorAlias := "initiator_alias"
	fieldInitiatorWwpn := "initiator_wwpn"
	fieldInitiatorWwnn := "initiator_wwnn"
	fieldInitiatorSwitchName := "initiator_switch_name"
	fieldInitiatorSwitchPort := "initiator_switch_port"
	fieldInitiatorSymbolicPortname := "initiator_symbolic_portname"
	fieldInitiatorSymbolicNodename := "initiator_symbolic_nodename"
	fieldInitiatorFcid := "initiator_fcid"

	NsFcSessionInitiatorFields = &NsFcSessionInitiatorFieldHandles{
		InitiatorAlias:            &fieldInitiatorAlias,
		InitiatorWwpn:             &fieldInitiatorWwpn,
		InitiatorWwnn:             &fieldInitiatorWwnn,
		InitiatorSwitchName:       &fieldInitiatorSwitchName,
		InitiatorSwitchPort:       &fieldInitiatorSwitchPort,
		InitiatorSymbolicPortname: &fieldInitiatorSymbolicPortname,
		InitiatorSymbolicNodename: &fieldInitiatorSymbolicNodename,
		InitiatorFcid:             &fieldInitiatorFcid,
	}
}

// NsFcSessionInitiator - Information of the Fibre Channel Session Initiator.
type NsFcSessionInitiator struct {
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
	InitiatorFcid *string `json:"initiator_fcid,omitempty"`
}

// NsFcSessionInitiatorFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcSessionInitiatorFieldHandles struct {
	InitiatorAlias            *string
	InitiatorWwpn             *string
	InitiatorWwnn             *string
	InitiatorSwitchName       *string
	InitiatorSwitchPort       *string
	InitiatorSymbolicPortname *string
	InitiatorSymbolicNodename *string
	InitiatorFcid             *string
}
