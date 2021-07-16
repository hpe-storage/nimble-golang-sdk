// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectOwnerPairWithType - List of objects of a given type along with their owners.
// Export NsObjectOwnerPairWithTypeFields for advance operations like search filter etc.
var NsObjectOwnerPairWithTypeFields *NsObjectOwnerPairWithType

func init() {

	NsObjectOwnerPairWithTypeFields = &NsObjectOwnerPairWithType{}
}

type NsObjectOwnerPairWithType struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjOwnerPairList - List of object names and owners.
	ObjOwnerPairList []*NsObjectOwnerPair `json:"obj_owner_pair_list,omitempty"`
}
