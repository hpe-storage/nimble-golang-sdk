/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsObjectLimit


// NsObjectLimit :
type NsObjectLimit struct {
   // ObjLimit
   ObjLimit float64 `json:"obj_limit,omitempty"`
   // ObjNum
   ObjNum float64 `json:"obj_num,omitempty"`
}
