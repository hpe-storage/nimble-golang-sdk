/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsUserLockStatus


// NsUserLockStatus :
type NsUserLockStatus struct {
   // ID
   ID string `json:"id,omitempty"`
   // AuthLock
   AuthLock bool `json:"auth_lock,omitempty"`
}
