// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// FibreChannelSession - Fibre Channel session is created when Fibre Channel initiator connects to this group.
// Export FibreChannelSessionFields for advance operations like search filter etc.
var FibreChannelSessionFields *FibreChannelSession

func init() {

	FibreChannelSessionFields = &FibreChannelSession{
		ID: "id",
	}
}

type FibreChannelSession struct {
	// ID - Unique identifier of the Fibre Channel session.
	ID string `json:"id,omitempty"`
	// InitiatorInfo - Information about the Fibre Channel initiator.
	InitiatorInfo *NsFcSessionInitiator `json:"initiator_info,omitempty"`
	// TargetInfo - Information about the Fibre Channel target.
	TargetInfo *NsFcSessionTarget `json:"target_info,omitempty"`
}

// sdk internal struct
type fibreChannelSession struct {
	// ID - Unique identifier of the Fibre Channel session.
	ID *string `json:"id,omitempty"`
	// InitiatorInfo - Information about the Fibre Channel initiator.
	InitiatorInfo *NsFcSessionInitiator `json:"initiator_info,omitempty"`
	// TargetInfo - Information about the Fibre Channel target.
	TargetInfo *NsFcSessionTarget `json:"target_info,omitempty"`
}

// EncodeFibreChannelSession - Transform FibreChannelSession to fibreChannelSession type
func EncodeFibreChannelSession(request interface{}) (*fibreChannelSession, error) {
	reqFibreChannelSession := request.(*FibreChannelSession)
	byte, err := json.Marshal(reqFibreChannelSession)
	objPtr := &fibreChannelSession{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeFibreChannelSession Transform fibreChannelSession to FibreChannelSession type
func DecodeFibreChannelSession(request interface{}) (*FibreChannelSession, error) {
	reqFibreChannelSession := request.(*fibreChannelSession)
	byte, err := json.Marshal(reqFibreChannelSession)
	obj := &FibreChannelSession{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
