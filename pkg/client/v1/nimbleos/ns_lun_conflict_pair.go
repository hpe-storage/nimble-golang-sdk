// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsLunConflictPair - LUN number conflict.
// Export NsLunConflictPairFields for advance operations like search filter etc.
var NsLunConflictPairFields *NsLunConflictPair

func init() {
	InitiatorWwpnfield := "initiator_wwpn"
	InitiatorAliasfield := "initiator_alias"
	DstIgrpNamefield := "dst_igrp_name"
	DstVolNamefield := "dst_vol_name"
	DstSnapNamefield := "dst_snap_name"
	DstPeNamefield := "dst_pe_name"
	SrcIgrpNamefield := "src_igrp_name"
	SrcVolNamefield := "src_vol_name"
	SrcSnapNamefield := "src_snap_name"

	NsLunConflictPairFields = &NsLunConflictPair{
		InitiatorWwpn:  &InitiatorWwpnfield,
		InitiatorAlias: &InitiatorAliasfield,
		DstIgrpName:    &DstIgrpNamefield,
		DstVolName:     &DstVolNamefield,
		DstSnapName:    &DstSnapNamefield,
		DstPeName:      &DstPeNamefield,
		SrcIgrpName:    &SrcIgrpNamefield,
		SrcVolName:     &SrcVolNamefield,
		SrcSnapName:    &SrcSnapNamefield,
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
