// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsObjectOwnerPairFields provides field names to use in filter parameters, for example.
var NsObjectOwnerPairFields *NsObjectOwnerPairFieldHandles

func init() {
	fieldObjName := "obj_name"
	fieldSrcOwner := "src_owner"
	fieldDstOwner := "dst_owner"

	NsObjectOwnerPairFields = &NsObjectOwnerPairFieldHandles{
		ObjName:  &fieldObjName,
		SrcOwner: &fieldSrcOwner,
		DstOwner: &fieldDstOwner,
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

// NsObjectOwnerPairFieldHandles provides a string representation for each AccessControlRecord field.
type NsObjectOwnerPairFieldHandles struct {
	ObjName  *string
	SrcOwner *string
	DstOwner *string
}
