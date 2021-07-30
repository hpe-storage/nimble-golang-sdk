// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPrivilege - Privilege info.

// Export NsPrivilegeFields provides field names to use in filter parameters, for example.
var NsPrivilegeFields *NsPrivilegeStringFields

func init() {
	fieldObjectType := "object_type"
	fieldOperations := "operations"

	NsPrivilegeFields = &NsPrivilegeStringFields{
		ObjectType: &fieldObjectType,
		Operations: &fieldOperations,
	}
}

type NsPrivilege struct {
	// ObjectType - Object type name associated with this privilege.
	ObjectType *string `json:"object_type,omitempty"`
	// Operations - List of operations associated with the above object for this privilege.
	Operations []*string `json:"operations,omitempty"`
}

// Struct for NsPrivilegeFields
type NsPrivilegeStringFields struct {
	ObjectType *string
	Operations *string
}
