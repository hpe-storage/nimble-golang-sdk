// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArrayUpgrade - Array upgrade attributes.
// Export NsArrayUpgradeFields for advance operations like search filter etc.
var NsArrayUpgradeFields *NsArrayUpgrade

func init() {

	NsArrayUpgradeFields = &NsArrayUpgrade{}
}

type NsArrayUpgrade struct {
	// Type - Array upgrade type.
	Type *NsArrayUpgradeType `json:"type,omitempty"`
	// State - Array upgrade state.
	State *NsArrayUpgradeState `json:"state,omitempty"`
	// Stage - Array upgrade stage.
	Stage *NsArrayUpgradeStage `json:"stage,omitempty"`
	// Messages - List of error or info messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// sdk internal struct
type nsArrayUpgrade struct {
	// Type - Array upgrade type.
	Type *NsArrayUpgradeType `json:"type,omitempty"`
	// State - Array upgrade state.
	State *NsArrayUpgradeState `json:"state,omitempty"`
	// Stage - Array upgrade stage.
	Stage *NsArrayUpgradeStage `json:"stage,omitempty"`
	// Messages - List of error or info messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// EncodeNsArrayUpgrade - Transform NsArrayUpgrade to nsArrayUpgrade type
func EncodeNsArrayUpgrade(request interface{}) (*nsArrayUpgrade, error) {
	reqNsArrayUpgrade := request.(*NsArrayUpgrade)
	byte, err := json.Marshal(reqNsArrayUpgrade)
	objPtr := &nsArrayUpgrade{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsArrayUpgrade Transform nsArrayUpgrade to NsArrayUpgrade type
func DecodeNsArrayUpgrade(request interface{}) (*NsArrayUpgrade, error) {
	reqNsArrayUpgrade := request.(*nsArrayUpgrade)
	byte, err := json.Marshal(reqNsArrayUpgrade)
	obj := &NsArrayUpgrade{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
