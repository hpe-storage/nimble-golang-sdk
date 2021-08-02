// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectOwnerPairFields provides field names to use in filter parameters, for example.
var NsObjectOwnerPairFields *NsObjectOwnerPairFieldHandles

func init() {
	NsObjectOwnerPairFields = &NsObjectOwnerPairFieldHandles{
		ObjName:  "obj_name",
		SrcOwner: "src_owner",
		DstOwner: "dst_owner",
	}
}

// NsObjectOwnerPair - Objects and their owners.
type NsObjectOwnerPair struct {
	// ObjName - Object name. Same on source and destination.
	ObjName *string `json:"obj_name,omitempty"`
	// SrcOwner - Name of the owner on the source group.
	SrcOwner *string `json:"src_owner,omitempty"`
	// DstOwner - Name of the owner on the destination group.
	DstOwner *string `json:"dst_owner,omitempty"`
}

// NsObjectOwnerPairFieldHandles provides a string representation for each NsObjectOwnerPair field.
type NsObjectOwnerPairFieldHandles struct {
	ObjName  string
	SrcOwner string
	DstOwner string
}
