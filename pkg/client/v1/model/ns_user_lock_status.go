// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsUserLockStatus - User account lock status.
// Export NsUserLockStatusFields for advance operations like search filter etc.
var NsUserLockStatusFields *NsUserLockStatus

func init(){
	IDfield:= "id"
		
	NsUserLockStatusFields= &NsUserLockStatus{
		ID:       &IDfield,
	}
}

type NsUserLockStatus struct {
	// ID - Unique Object ID.
 	ID *string `json:"id,omitempty"`
	// AuthLock - User was locked due to failed logins.
    AuthLock *bool `json:"auth_lock,omitempty"`
}
