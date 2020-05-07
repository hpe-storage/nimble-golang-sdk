/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsLunConflictPair


// NsLunConflictPair :
type NsLunConflictPair struct {
   // InitiatorWwpn
   InitiatorWwpn string `json:"initiator_wwpn,omitempty"`
   // InitiatorAlias
   InitiatorAlias string `json:"initiator_alias,omitempty"`
   // DstIgrpName
   DstIgrpName string `json:"dst_igrp_name,omitempty"`
   // DstVolName
   DstVolName string `json:"dst_vol_name,omitempty"`
   // DstSnapName
   DstSnapName string `json:"dst_snap_name,omitempty"`
   // DstPeName
   DstPeName string `json:"dst_pe_name,omitempty"`
   // DstLun
   DstLun float64 `json:"dst_lun,omitempty"`
   // SrcIgrpName
   SrcIgrpName string `json:"src_igrp_name,omitempty"`
   // SrcVolName
   SrcVolName string `json:"src_vol_name,omitempty"`
   // SrcSnapName
   SrcSnapName string `json:"src_snap_name,omitempty"`
   // SrcLun
   SrcLun float64 `json:"src_lun,omitempty"`
}
