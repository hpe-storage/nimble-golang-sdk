// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsUserLockStatusFields provides field names to use in filter parameters, for example.
var NsUserLockStatusFields *NsUserLockStatusFieldHandles

func init() {
	fieldID := "id"
	fieldAuthLock := "auth_lock"

	NsUserLockStatusFields = &NsUserLockStatusFieldHandles{
		ID:       &fieldID,
		AuthLock: &fieldAuthLock,
	}
}

// NsUserLockStatus - User account lock status.
type NsUserLockStatus struct {
	// ID - Unique Object ID.
	ID *string `json:"id,omitempty"`
	// AuthLock - User was locked due to failed logins.
	AuthLock *bool `json:"auth_lock,omitempty"`
}

// NsUserLockStatusFieldHandles provides a string representation for each AccessControlRecord field.
type NsUserLockStatusFieldHandles struct {
	ID       *string
	AuthLock *string
}
