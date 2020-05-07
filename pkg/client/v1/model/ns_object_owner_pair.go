/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsObjectOwnerPair


// NsObjectOwnerPair :
type NsObjectOwnerPair struct {
   // ObjName
   ObjName string `json:"obj_name,omitempty"`
   // SrcOwner
   SrcOwner string `json:"src_owner,omitempty"`
   // DstOwner
   DstOwner string `json:"dst_owner,omitempty"`
}
