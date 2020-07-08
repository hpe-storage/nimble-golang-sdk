// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsExtendedPrivilege - Extended privilege info.
// Export NsExtendedPrivilegeFields for advance operations like search filter etc.
var NsExtendedPrivilegeFields *NsExtendedPrivilege

func init() {

	NsExtendedPrivilegeFields = &NsExtendedPrivilege{
		ObjectType: "object_type",
		Operation:  "operation",
	}
}

type NsExtendedPrivilege struct {
	// ObjectType - Object type name associated with this privilege.
	ObjectType string `json:"object_type,omitempty"`
	// Operation - Operation associated with the above object for this privilege.
	Operation string `json:"operation,omitempty"`
	// Allow - Indicate whether the above operation is allowed for this privilege.
	Allow *bool `json:"allow,omitempty"`
}

// sdk internal struct
type nsExtendedPrivilege struct {
	// ObjectType - Object type name associated with this privilege.
	ObjectType *string `json:"object_type,omitempty"`
	// Operation - Operation associated with the above object for this privilege.
	Operation *string `json:"operation,omitempty"`
	// Allow - Indicate whether the above operation is allowed for this privilege.
	Allow *bool `json:"allow,omitempty"`
}

// EncodeNsExtendedPrivilege - Transform NsExtendedPrivilege to nsExtendedPrivilege type
func EncodeNsExtendedPrivilege(request interface{}) (*nsExtendedPrivilege, error) {
	reqNsExtendedPrivilege := request.(*NsExtendedPrivilege)
	byte, err := json.Marshal(reqNsExtendedPrivilege)
	objPtr := &nsExtendedPrivilege{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsExtendedPrivilege Transform nsExtendedPrivilege to NsExtendedPrivilege type
func DecodeNsExtendedPrivilege(request interface{}) (*NsExtendedPrivilege, error) {
	reqNsExtendedPrivilege := request.(*nsExtendedPrivilege)
	byte, err := json.Marshal(reqNsExtendedPrivilege)
	obj := &NsExtendedPrivilege{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
