// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsUserLockStatusFields provides field names to use in filter parameters, for example.
var NsUserLockStatusFields *NsUserLockStatusFieldHandles

func init() {
	NsUserLockStatusFields = &NsUserLockStatusFieldHandles{
		ID:       "id",
		AuthLock: "auth_lock",
	}
}

// NsUserLockStatus - User account lock status.
type NsUserLockStatus struct {
	// ID - Unique Object ID.
	ID *string `json:"id,omitempty"`
	// AuthLock - User was locked due to failed logins.
	AuthLock *bool `json:"auth_lock,omitempty"`
}

// NsUserLockStatusFieldHandles provides a string representation for each NsUserLockStatus field.
type NsUserLockStatusFieldHandles struct {
	ID       string
	AuthLock string
}
