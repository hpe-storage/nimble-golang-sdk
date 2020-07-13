// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsISCSIIQN - ISCSI IQN.
// Export NsISCSIIQNFields for advance operations like search filter etc.
var NsISCSIIQNFields *NsISCSIIQN

func init() {

	NsISCSIIQNFields = &NsISCSIIQN{
		Name: "name",
	}
}

type NsISCSIIQN struct {
	// Name - IQN name of the iSCSI initiator.
	Name string `json:"name,omitempty"`
}

// sdk internal struct
type nsISCSIIQN struct {
	// Name - IQN name of the iSCSI initiator.
	Name *string `json:"name,omitempty"`
}

// EncodeNsISCSIIQN - Transform NsISCSIIQN to nsISCSIIQN type
func EncodeNsISCSIIQN(request interface{}) (*nsISCSIIQN, error) {
	reqNsISCSIIQN := request.(*NsISCSIIQN)
	byte, err := json.Marshal(reqNsISCSIIQN)
	objPtr := &nsISCSIIQN{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsISCSIIQN Transform nsISCSIIQN to NsISCSIIQN type
func DecodeNsISCSIIQN(request interface{}) (*NsISCSIIQN, error) {
	reqNsISCSIIQN := request.(*nsISCSIIQN)
	byte, err := json.Marshal(reqNsISCSIIQN)
	obj := &NsISCSIIQN{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
