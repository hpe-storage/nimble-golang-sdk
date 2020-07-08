// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsUserLockStatus - User account lock status.
// Export NsUserLockStatusFields for advance operations like search filter etc.
var NsUserLockStatusFields *NsUserLockStatus

func init() {

	NsUserLockStatusFields = &NsUserLockStatus{
		ID: "id",
	}
}

type NsUserLockStatus struct {
	// ID - Unique Object ID.
	ID string `json:"id,omitempty"`
	// AuthLock - User was locked due to failed logins.
	AuthLock *bool `json:"auth_lock,omitempty"`
}

// sdk internal struct
type nsUserLockStatus struct {
	// ID - Unique Object ID.
	ID *string `json:"id,omitempty"`
	// AuthLock - User was locked due to failed logins.
	AuthLock *bool `json:"auth_lock,omitempty"`
}

// EncodeNsUserLockStatus - Transform NsUserLockStatus to nsUserLockStatus type
func EncodeNsUserLockStatus(request interface{}) (*nsUserLockStatus, error) {
	reqNsUserLockStatus := request.(*NsUserLockStatus)
	byte, err := json.Marshal(reqNsUserLockStatus)
	objPtr := &nsUserLockStatus{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsUserLockStatus Transform nsUserLockStatus to NsUserLockStatus type
func DecodeNsUserLockStatus(request interface{}) (*NsUserLockStatus, error) {
	reqNsUserLockStatus := request.(*nsUserLockStatus)
	byte, err := json.Marshal(reqNsUserLockStatus)
	obj := &NsUserLockStatus{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
