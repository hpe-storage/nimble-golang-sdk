// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectOwnerPairWithTypeFields provides field names to use in filter parameters, for example.
var NsObjectOwnerPairWithTypeFields *NsObjectOwnerPairWithTypeFieldHandles

func init() {
	NsObjectOwnerPairWithTypeFields = &NsObjectOwnerPairWithTypeFieldHandles{
		ObjType:          "obj_type",
		ObjOwnerPairList: "obj_owner_pair_list",
	}
}

// NsObjectOwnerPairWithType - List of objects of a given type along with their owners.
type NsObjectOwnerPairWithType struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjOwnerPairList - List of object names and owners.
	ObjOwnerPairList []*NsObjectOwnerPair `json:"obj_owner_pair_list,omitempty"`
}

// NsObjectOwnerPairWithTypeFieldHandles provides a string representation for each NsObjectOwnerPairWithType field.
type NsObjectOwnerPairWithTypeFieldHandles struct {
	ObjType          string
	ObjOwnerPairList string
}
