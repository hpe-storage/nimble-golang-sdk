// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAliasConflictPair - Alias conflict (same initiator WWPN, but different aliases).

// Export NsAliasConflictPairFields provides field names to use in filter parameters, for example.
var NsAliasConflictPairFields *NsAliasConflictPairStringFields

func init() {
	fieldInitiatorWwpn := "initiator_wwpn"
	fieldDstAliasName := "dst_alias_name"
	fieldSrcAliasName := "src_alias_name"

	NsAliasConflictPairFields = &NsAliasConflictPairStringFields{
		InitiatorWwpn: &fieldInitiatorWwpn,
		DstAliasName:  &fieldDstAliasName,
		SrcAliasName:  &fieldSrcAliasName,
	}
}

type NsAliasConflictPair struct {
	// InitiatorWwpn - WWPN of the common initiator.
	InitiatorWwpn *string `json:"initiator_wwpn,omitempty"`
	// DstAliasName - Name of the alias on the destination group.
	DstAliasName *string `json:"dst_alias_name,omitempty"`
	// SrcAliasName - Name of the alias on the source group.
	SrcAliasName *string `json:"src_alias_name,omitempty"`
}

// Struct for NsAliasConflictPairFields
type NsAliasConflictPairStringFields struct {
	InitiatorWwpn *string
	DstAliasName  *string
	SrcAliasName  *string
}
