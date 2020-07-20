// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	ID       *string `json:"id,omitempty"`
	AuthLock *bool   `json:"auth_lock,omitempty"`
}

// EncodeNsUserLockStatus - Transform NsUserLockStatus to nsUserLockStatus type
func EncodeNsUserLockStatus(request interface{}) (*nsUserLockStatus, error) {
	reqNsUserLockStatus := request.(*NsUserLockStatus)
	byte, err := json.Marshal(reqNsUserLockStatus)
	if err != nil {
		return nil, err
	}
	respNsUserLockStatusPtr := &nsUserLockStatus{}
	err = json.Unmarshal(byte, respNsUserLockStatusPtr)
	return respNsUserLockStatusPtr, err
}

// DecodeNsUserLockStatus Transform nsUserLockStatus to NsUserLockStatus type
func DecodeNsUserLockStatus(request interface{}) (*NsUserLockStatus, error) {
	reqNsUserLockStatus := request.(*nsUserLockStatus)
	byte, err := json.Marshal(reqNsUserLockStatus)
	if err != nil {
		return nil, err
	}
	respNsUserLockStatus := &NsUserLockStatus{}
	err = json.Unmarshal(byte, respNsUserLockStatus)
	return respNsUserLockStatus, err
}
