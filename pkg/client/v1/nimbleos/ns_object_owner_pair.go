// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectOwnerPair - Objects and their owners.

// Export NsObjectOwnerPairFields provides field names to use in filter parameters, for example.
var NsObjectOwnerPairFields *NsObjectOwnerPairStringFields

func init() {
	fieldObjName := "obj_name"
	fieldSrcOwner := "src_owner"
	fieldDstOwner := "dst_owner"

	NsObjectOwnerPairFields = &NsObjectOwnerPairStringFields{
		ObjName:  &fieldObjName,
		SrcOwner: &fieldSrcOwner,
		DstOwner: &fieldDstOwner,
	}
}

type NsObjectOwnerPair struct {
	// ObjName - Object name. Same on source and destination.
	ObjName *string `json:"obj_name,omitempty"`
	// SrcOwner - Name of the owner on the source group.
	SrcOwner *string `json:"src_owner,omitempty"`
	// DstOwner - Name of the owner on the destination group.
	DstOwner *string `json:"dst_owner,omitempty"`
}

// Struct for NsObjectOwnerPairFields
type NsObjectOwnerPairStringFields struct {
	ObjName  *string
	SrcOwner *string
	DstOwner *string
}
