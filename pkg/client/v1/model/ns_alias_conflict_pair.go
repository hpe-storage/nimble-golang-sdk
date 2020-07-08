// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsAliasConflictPair - Alias conflict (same initiator WWPN, but different aliases).
// Export NsAliasConflictPairFields for advance operations like search filter etc.
var NsAliasConflictPairFields *NsAliasConflictPair

func init() {

	NsAliasConflictPairFields = &NsAliasConflictPair{
		InitiatorWwpn: "initiator_wwpn",
		DstAliasName:  "dst_alias_name",
		SrcAliasName:  "src_alias_name",
	}
}

type NsAliasConflictPair struct {
	// InitiatorWwpn - WWPN of the common initiator.
	InitiatorWwpn string `json:"initiator_wwpn,omitempty"`
	// DstAliasName - Name of the alias on the destination group.
	DstAliasName string `json:"dst_alias_name,omitempty"`
	// SrcAliasName - Name of the alias on the source group.
	SrcAliasName string `json:"src_alias_name,omitempty"`
}

// sdk internal struct
type nsAliasConflictPair struct {
	// InitiatorWwpn - WWPN of the common initiator.
	InitiatorWwpn *string `json:"initiator_wwpn,omitempty"`
	// DstAliasName - Name of the alias on the destination group.
	DstAliasName *string `json:"dst_alias_name,omitempty"`
	// SrcAliasName - Name of the alias on the source group.
	SrcAliasName *string `json:"src_alias_name,omitempty"`
}

// EncodeNsAliasConflictPair - Transform NsAliasConflictPair to nsAliasConflictPair type
func EncodeNsAliasConflictPair(request interface{}) (*nsAliasConflictPair, error) {
	reqNsAliasConflictPair := request.(*NsAliasConflictPair)
	byte, err := json.Marshal(reqNsAliasConflictPair)
	objPtr := &nsAliasConflictPair{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsAliasConflictPair Transform nsAliasConflictPair to NsAliasConflictPair type
func DecodeNsAliasConflictPair(request interface{}) (*NsAliasConflictPair, error) {
	reqNsAliasConflictPair := request.(*nsAliasConflictPair)
	byte, err := json.Marshal(reqNsAliasConflictPair)
	obj := &NsAliasConflictPair{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
