/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsAliasConflictPair


// NsAliasConflictPair :
type NsAliasConflictPair struct {
   // InitiatorWwpn
   InitiatorWwpn string `json:"initiator_wwpn,omitempty"`
   // DstAliasName
   DstAliasName string `json:"dst_alias_name,omitempty"`
   // SrcAliasName
   SrcAliasName string `json:"src_alias_name,omitempty"`
}
