// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPrivilege - Privilege info.
// Export NsPrivilegeFields for advance operations like search filter etc.
var NsPrivilegeFields *NsPrivilegeStringFields

func init() {
	ObjectTypefield := "object_type"
	Operationsfield := "operations"

	NsPrivilegeFields = &NsPrivilegeStringFields{
		ObjectType: &ObjectTypefield,
		Operations: &Operationsfield,
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
