// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsExtendedPrivilege - Extended privilege info.

// Export NsExtendedPrivilegeFields provides field names to use in filter parameters, for example.
var NsExtendedPrivilegeFields *NsExtendedPrivilegeStringFields

func init() {
	fieldObjectType := "object_type"
	fieldOperation := "operation"
	fieldAllow := "allow"

	NsExtendedPrivilegeFields = &NsExtendedPrivilegeStringFields{
		ObjectType: &fieldObjectType,
		Operation:  &fieldOperation,
		Allow:      &fieldAllow,
	}
}

type NsExtendedPrivilege struct {
	// ObjectType - Object type name associated with this privilege.
	ObjectType *string `json:"object_type,omitempty"`
	// Operation - Operation associated with the above object for this privilege.
	Operation *string `json:"operation,omitempty"`
	// Allow - Indicate whether the above operation is allowed for this privilege.
	Allow *bool `json:"allow,omitempty"`
}

// Struct for NsExtendedPrivilegeFields
type NsExtendedPrivilegeStringFields struct {
	ObjectType *string
	Operation  *string
	Allow      *string
}
