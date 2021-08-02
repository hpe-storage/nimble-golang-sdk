// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcSessionInitiatorFields provides field names to use in filter parameters, for example.
var NsFcSessionInitiatorFields *NsFcSessionInitiatorFieldHandles

func init() {
	NsFcSessionInitiatorFields = &NsFcSessionInitiatorFieldHandles{
		InitiatorAlias:            "initiator_alias",
		InitiatorWwpn:             "initiator_wwpn",
		InitiatorWwnn:             "initiator_wwnn",
		InitiatorSwitchName:       "initiator_switch_name",
		InitiatorSwitchPort:       "initiator_switch_port",
		InitiatorSymbolicPortname: "initiator_symbolic_portname",
		InitiatorSymbolicNodename: "initiator_symbolic_nodename",
		InitiatorFcid:             "initiator_fcid",
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
	InitiatorAlias            string
	InitiatorWwpn             string
	InitiatorWwnn             string
	InitiatorSwitchName       string
	InitiatorSwitchPort       string
	InitiatorSymbolicPortname string
	InitiatorSymbolicNodename string
	InitiatorFcid             string
}
