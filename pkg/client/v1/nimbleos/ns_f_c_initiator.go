// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFCInitiator - Fibre Channel initiator.
// Export NsFCInitiatorFields for advance operations like search filter etc.
var NsFCInitiatorFields *NsFCInitiator

func init() {

	NsFCInitiatorFields = &NsFCInitiator{
		ID:          "id",
		InitiatorId: "initiator_id",
		Wwpn:        "wwpn",
		Alias:       "alias",
	}
}

type NsFCInitiator struct {
	// ID - Unique identifier of the Fibre Channel initiator.
	ID string `json:"id,omitempty"`
	// InitiatorId - Unique identifier of the Fibre Channel initiator.
	InitiatorId string `json:"initiator_id,omitempty"`
	// Wwpn - WWPN (World Wide Port Name) of the Fibre Channel initiator.
	Wwpn string `json:"wwpn,omitempty"`
	// Alias - Alias of the Fibre Channel initiator.
	Alias string `json:"alias,omitempty"`
}

// sdk internal struct
type nsFCInitiator struct {
	ID          *string `json:"id,omitempty"`
	InitiatorId *string `json:"initiator_id,omitempty"`
	Wwpn        *string `json:"wwpn,omitempty"`
	Alias       *string `json:"alias,omitempty"`
}

// EncodeNsFCInitiator - Transform NsFCInitiator to nsFCInitiator type
func EncodeNsFCInitiator(request interface{}) (*nsFCInitiator, error) {
	reqNsFCInitiator := request.(*NsFCInitiator)
	byte, err := json.Marshal(reqNsFCInitiator)
	if err != nil {
		return nil, err
	}
	respNsFCInitiatorPtr := &nsFCInitiator{}
	err = json.Unmarshal(byte, respNsFCInitiatorPtr)
	return respNsFCInitiatorPtr, err
}

// DecodeNsFCInitiator Transform nsFCInitiator to NsFCInitiator type
func DecodeNsFCInitiator(request interface{}) (*NsFCInitiator, error) {
	reqNsFCInitiator := request.(*nsFCInitiator)
	byte, err := json.Marshal(reqNsFCInitiator)
	if err != nil {
		return nil, err
	}
	respNsFCInitiator := &NsFCInitiator{}
	err = json.Unmarshal(byte, respNsFCInitiator)
	return respNsFCInitiator, err
}
