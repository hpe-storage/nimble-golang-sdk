// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsObjectOwnerPair - Objects and their owners.
// Export NsObjectOwnerPairFields for advance operations like search filter etc.
var NsObjectOwnerPairFields *NsObjectOwnerPair

func init(){
	ObjNamefield:= "obj_name"
	SrcOwnerfield:= "src_owner"
	DstOwnerfield:= "dst_owner"
		
	NsObjectOwnerPairFields= &NsObjectOwnerPair{
		ObjName:  &ObjNamefield,
		SrcOwner: &SrcOwnerfield,
		DstOwner: &DstOwnerfield,
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
