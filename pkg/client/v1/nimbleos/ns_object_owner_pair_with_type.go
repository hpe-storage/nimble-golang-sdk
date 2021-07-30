// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectOwnerPairWithType - List of objects of a given type along with their owners.

// Export NsObjectOwnerPairWithTypeFields provides field names to use in filter parameters, for example.
var NsObjectOwnerPairWithTypeFields *NsObjectOwnerPairWithTypeStringFields

func init() {
	fieldObjType := "obj_type"
	fieldObjOwnerPairList := "obj_owner_pair_list"

	NsObjectOwnerPairWithTypeFields = &NsObjectOwnerPairWithTypeStringFields{
		ObjType:          &fieldObjType,
		ObjOwnerPairList: &fieldObjOwnerPairList,
	}
}

type NsObjectOwnerPairWithType struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjOwnerPairList - List of object names and owners.
	ObjOwnerPairList []*NsObjectOwnerPair `json:"obj_owner_pair_list,omitempty"`
}

// Struct for NsObjectOwnerPairWithTypeFields
type NsObjectOwnerPairWithTypeStringFields struct {
	ObjType          *string
	ObjOwnerPairList *string
}
