// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	ObjectType *string `json:"object_type,omitempty"`
	Operation  *string `json:"operation,omitempty"`
	Allow      *bool   `json:"allow,omitempty"`
}

// EncodeNsExtendedPrivilege - Transform NsExtendedPrivilege to nsExtendedPrivilege type
func EncodeNsExtendedPrivilege(request interface{}) (*nsExtendedPrivilege, error) {
	reqNsExtendedPrivilege := request.(*NsExtendedPrivilege)
	byte, err := json.Marshal(reqNsExtendedPrivilege)
	if err != nil {
		return nil, err
	}
	respNsExtendedPrivilegePtr := &nsExtendedPrivilege{}
	err = json.Unmarshal(byte, respNsExtendedPrivilegePtr)
	return respNsExtendedPrivilegePtr, err
}

// DecodeNsExtendedPrivilege Transform nsExtendedPrivilege to NsExtendedPrivilege type
func DecodeNsExtendedPrivilege(request interface{}) (*NsExtendedPrivilege, error) {
	reqNsExtendedPrivilege := request.(*nsExtendedPrivilege)
	byte, err := json.Marshal(reqNsExtendedPrivilege)
	if err != nil {
		return nil, err
	}
	respNsExtendedPrivilege := &NsExtendedPrivilege{}
	err = json.Unmarshal(byte, respNsExtendedPrivilege)
	return respNsExtendedPrivilege, err
}
