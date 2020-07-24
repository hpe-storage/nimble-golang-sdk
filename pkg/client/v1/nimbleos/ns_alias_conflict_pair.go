// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAliasConflictPair - Alias conflict (same initiator WWPN, but different aliases).
// Export NsAliasConflictPairFields for advance operations like search filter etc.
var NsAliasConflictPairFields *NsAliasConflictPair

func init() {
	InitiatorWwpnfield := "initiator_wwpn"
	DstAliasNamefield := "dst_alias_name"
	SrcAliasNamefield := "src_alias_name"

	NsAliasConflictPairFields = &NsAliasConflictPair{
		InitiatorWwpn: &InitiatorWwpnfield,
		DstAliasName:  &DstAliasNamefield,
		SrcAliasName:  &SrcAliasNamefield,
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
