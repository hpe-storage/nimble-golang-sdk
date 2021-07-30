// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsUserLockStatus - User account lock status.
// Export NsUserLockStatusFields for advance operations like search filter etc.
var NsUserLockStatusFields *NsUserLockStatusStringFields

func init() {
	IDfield := "id"
	AuthLockfield := "auth_lock"

	NsUserLockStatusFields = &NsUserLockStatusStringFields{
		ID:       &IDfield,
		AuthLock: &AuthLockfield,
	}
}

type NsUserLockStatus struct {
	// ID - Unique Object ID.
	ID *string `json:"id,omitempty"`
	// AuthLock - User was locked due to failed logins.
	AuthLock *bool `json:"auth_lock,omitempty"`
}

// Struct for NsUserLockStatusFields
type NsUserLockStatusStringFields struct {
	ID       *string
	AuthLock *string
}
