// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsLunConflictPair - LUN number conflict.
// Export NsLunConflictPairFields for advance operations like search filter etc.
var NsLunConflictPairFields *NsLunConflictPair

func init() {

	NsLunConflictPairFields = &NsLunConflictPair{
		InitiatorWwpn:  "initiator_wwpn",
		InitiatorAlias: "initiator_alias",
		DstIgrpName:    "dst_igrp_name",
		DstVolName:     "dst_vol_name",
		DstSnapName:    "dst_snap_name",
		DstPeName:      "dst_pe_name",
		SrcIgrpName:    "src_igrp_name",
		SrcVolName:     "src_vol_name",
		SrcSnapName:    "src_snap_name",
	}
}

type NsLunConflictPair struct {
	// InitiatorWwpn - WWPN/IQN of the common initiator.
	InitiatorWwpn string `json:"initiator_wwpn,omitempty"`
	// InitiatorAlias - Alias of the common initiator (if exists).
	InitiatorAlias string `json:"initiator_alias,omitempty"`
	// DstIgrpName - Name of the initiator group on the destination group.
	DstIgrpName string `json:"dst_igrp_name,omitempty"`
	// DstVolName - Name of the volume on the destination group.
	DstVolName string `json:"dst_vol_name,omitempty"`
	// DstSnapName - Name of the snapshot on the destination group (if applicable).
	DstSnapName string `json:"dst_snap_name,omitempty"`
	// DstPeName - Name of the protocol endpoint on the destination group (if applicable).
	DstPeName string `json:"dst_pe_name,omitempty"`
	// DstLun - LUN number on the destination group.
	DstLun int64 `json:"dst_lun,omitempty"`
	// SrcIgrpName - Name of the initiator group on the source group.
	SrcIgrpName string `json:"src_igrp_name,omitempty"`
	// SrcVolName - Name of the volume on the source group.
	SrcVolName string `json:"src_vol_name,omitempty"`
	// SrcSnapName - Name of the snapshot on the source group (if applicable).
	SrcSnapName string `json:"src_snap_name,omitempty"`
	// SrcLun - LUN number on the source group.
	SrcLun int64 `json:"src_lun,omitempty"`
}

// sdk internal struct
type nsLunConflictPair struct {
	// InitiatorWwpn - WWPN/IQN of the common initiator.
	InitiatorWwpn *string `json:"initiator_wwpn,omitempty"`
	// InitiatorAlias - Alias of the common initiator (if exists).
	InitiatorAlias *string `json:"initiator_alias,omitempty"`
	// DstIgrpName - Name of the initiator group on the destination group.
	DstIgrpName *string `json:"dst_igrp_name,omitempty"`
	// DstVolName - Name of the volume on the destination group.
	DstVolName *string `json:"dst_vol_name,omitempty"`
	// DstSnapName - Name of the snapshot on the destination group (if applicable).
	DstSnapName *string `json:"dst_snap_name,omitempty"`
	// DstPeName - Name of the protocol endpoint on the destination group (if applicable).
	DstPeName *string `json:"dst_pe_name,omitempty"`
	// DstLun - LUN number on the destination group.
	DstLun *int64 `json:"dst_lun,omitempty"`
	// SrcIgrpName - Name of the initiator group on the source group.
	SrcIgrpName *string `json:"src_igrp_name,omitempty"`
	// SrcVolName - Name of the volume on the source group.
	SrcVolName *string `json:"src_vol_name,omitempty"`
	// SrcSnapName - Name of the snapshot on the source group (if applicable).
	SrcSnapName *string `json:"src_snap_name,omitempty"`
	// SrcLun - LUN number on the source group.
	SrcLun *int64 `json:"src_lun,omitempty"`
}

// EncodeNsLunConflictPair - Transform NsLunConflictPair to nsLunConflictPair type
func EncodeNsLunConflictPair(request interface{}) (*nsLunConflictPair, error) {
	reqNsLunConflictPair := request.(*NsLunConflictPair)
	byte, err := json.Marshal(reqNsLunConflictPair)
	objPtr := &nsLunConflictPair{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsLunConflictPair Transform nsLunConflictPair to NsLunConflictPair type
func DecodeNsLunConflictPair(request interface{}) (*NsLunConflictPair, error) {
	reqNsLunConflictPair := request.(*nsLunConflictPair)
	byte, err := json.Marshal(reqNsLunConflictPair)
	obj := &NsLunConflictPair{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
