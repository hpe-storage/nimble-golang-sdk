// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsLunConflictPair - LUN number conflict.

// Export NsLunConflictPairFields provides field names to use in filter parameters, for example.
var NsLunConflictPairFields *NsLunConflictPairStringFields

func init() {
	fieldInitiatorWwpn := "initiator_wwpn"
	fieldInitiatorAlias := "initiator_alias"
	fieldDstIgrpName := "dst_igrp_name"
	fieldDstVolName := "dst_vol_name"
	fieldDstSnapName := "dst_snap_name"
	fieldDstPeName := "dst_pe_name"
	fieldDstLun := "dst_lun"
	fieldSrcIgrpName := "src_igrp_name"
	fieldSrcVolName := "src_vol_name"
	fieldSrcSnapName := "src_snap_name"
	fieldSrcLun := "src_lun"

	NsLunConflictPairFields = &NsLunConflictPairStringFields{
		InitiatorWwpn:  &fieldInitiatorWwpn,
		InitiatorAlias: &fieldInitiatorAlias,
		DstIgrpName:    &fieldDstIgrpName,
		DstVolName:     &fieldDstVolName,
		DstSnapName:    &fieldDstSnapName,
		DstPeName:      &fieldDstPeName,
		DstLun:         &fieldDstLun,
		SrcIgrpName:    &fieldSrcIgrpName,
		SrcVolName:     &fieldSrcVolName,
		SrcSnapName:    &fieldSrcSnapName,
		SrcLun:         &fieldSrcLun,
	}
}

type NsLunConflictPair struct {
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

// Struct for NsLunConflictPairFields
type NsLunConflictPairStringFields struct {
	InitiatorWwpn  *string
	InitiatorAlias *string
	DstIgrpName    *string
	DstVolName     *string
	DstSnapName    *string
	DstPeName      *string
	DstLun         *string
	SrcIgrpName    *string
	SrcVolName     *string
	SrcSnapName    *string
	SrcLun         *string
}
