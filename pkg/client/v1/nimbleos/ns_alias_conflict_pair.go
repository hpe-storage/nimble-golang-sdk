// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsAliasConflictPairFields provides field names to use in filter parameters, for example.
var NsAliasConflictPairFields *NsAliasConflictPairFieldHandles

func init() {
	NsAliasConflictPairFields = &NsAliasConflictPairFieldHandles{
		InitiatorWwpn: "initiator_wwpn",
		DstAliasName:  "dst_alias_name",
		SrcAliasName:  "src_alias_name",
	}
}

// NsAliasConflictPair - Alias conflict (same initiator WWPN, but different aliases).
type NsAliasConflictPair struct {
	// InitiatorWwpn - WWPN of the common initiator.
	InitiatorWwpn *string `json:"initiator_wwpn,omitempty"`
	// DstAliasName - Name of the alias on the destination group.
	DstAliasName *string `json:"dst_alias_name,omitempty"`
	// SrcAliasName - Name of the alias on the source group.
	SrcAliasName *string `json:"src_alias_name,omitempty"`
}

// NsAliasConflictPairFieldHandles provides a string representation for each AccessControlRecord field.
type NsAliasConflictPairFieldHandles struct {
	InitiatorWwpn string
	DstAliasName  string
	SrcAliasName  string
}
