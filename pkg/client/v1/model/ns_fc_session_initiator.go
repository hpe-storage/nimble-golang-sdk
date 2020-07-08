// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsFcSessionInitiator - Information of the Fibre Channel Session Initiator.
// Export NsFcSessionInitiatorFields for advance operations like search filter etc.
var NsFcSessionInitiatorFields *NsFcSessionInitiator

func init() {

	NsFcSessionInitiatorFields = &NsFcSessionInitiator{
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

type NsFcSessionInitiator struct {
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
	InitiatorFcid string `json:"initiator_fcid,omitempty"`
}

// sdk internal struct
type nsFcSessionInitiator struct {
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

// EncodeNsFcSessionInitiator - Transform NsFcSessionInitiator to nsFcSessionInitiator type
func EncodeNsFcSessionInitiator(request interface{}) (*nsFcSessionInitiator, error) {
	reqNsFcSessionInitiator := request.(*NsFcSessionInitiator)
	byte, err := json.Marshal(reqNsFcSessionInitiator)
	objPtr := &nsFcSessionInitiator{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFcSessionInitiator Transform nsFcSessionInitiator to NsFcSessionInitiator type
func DecodeNsFcSessionInitiator(request interface{}) (*NsFcSessionInitiator, error) {
	reqNsFcSessionInitiator := request.(*nsFcSessionInitiator)
	byte, err := json.Marshal(reqNsFcSessionInitiator)
	obj := &NsFcSessionInitiator{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
