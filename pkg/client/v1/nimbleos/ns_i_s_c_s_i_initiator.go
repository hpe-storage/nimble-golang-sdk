// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsISCSIInitiator - ISCSI initiator.
// Export NsISCSIInitiatorFields for advance operations like search filter etc.
var NsISCSIInitiatorFields *NsISCSIInitiator

func init() {

	NsISCSIInitiatorFields = &NsISCSIInitiator{
		ID:          "id",
		InitiatorId: "initiator_id",
		Label:       "label",
		Iqn:         "iqn",
		IpAddress:   "ip_address",
	}
}

type NsISCSIInitiator struct {
	// ID - Unique identifier of the iSCSI initiator.
	ID string `json:"id,omitempty"`
	// InitiatorId - Unique identifier of the iSCSI initiator.
	InitiatorId string `json:"initiator_id,omitempty"`
	// Label - Unique label of the iSCSI initiator.
	Label string `json:"label,omitempty"`
	// Iqn - IQN name of the iSCSI initiator.
	Iqn string `json:"iqn,omitempty"`
	// IpAddress - IP address of the iSCSI initiator.
	IpAddress string `json:"ip_address,omitempty"`
}

// sdk internal struct
type nsISCSIInitiator struct {
	ID          *string `json:"id,omitempty"`
	InitiatorId *string `json:"initiator_id,omitempty"`
	Label       *string `json:"label,omitempty"`
	Iqn         *string `json:"iqn,omitempty"`
	IpAddress   *string `json:"ip_address,omitempty"`
}

// EncodeNsISCSIInitiator - Transform NsISCSIInitiator to nsISCSIInitiator type
func EncodeNsISCSIInitiator(request interface{}) (*nsISCSIInitiator, error) {
	reqNsISCSIInitiator := request.(*NsISCSIInitiator)
	byte, err := json.Marshal(reqNsISCSIInitiator)
	if err != nil {
		return nil, err
	}
	respNsISCSIInitiatorPtr := &nsISCSIInitiator{}
	err = json.Unmarshal(byte, respNsISCSIInitiatorPtr)
	return respNsISCSIInitiatorPtr, err
}

// DecodeNsISCSIInitiator Transform nsISCSIInitiator to NsISCSIInitiator type
func DecodeNsISCSIInitiator(request interface{}) (*NsISCSIInitiator, error) {
	reqNsISCSIInitiator := request.(*nsISCSIInitiator)
	byte, err := json.Marshal(reqNsISCSIInitiator)
	if err != nil {
		return nil, err
	}
	respNsISCSIInitiator := &NsISCSIInitiator{}
	err = json.Unmarshal(byte, respNsISCSIInitiator)
	return respNsISCSIInitiator, err
}
