// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPrivilegeFields provides field names to use in filter parameters, for example.
var NsPrivilegeFields *NsPrivilegeFieldHandles

func init() {
	NsPrivilegeFields = &NsPrivilegeFieldHandles{
		ObjectType: "object_type",
		Operations: "operations",
	}
}

// NsPrivilege - Privilege info.
type NsPrivilege struct {
	// ObjectType - Object type name associated with this privilege.
	ObjectType *string `json:"object_type,omitempty"`
	// Operations - List of operations associated with the above object for this privilege.
	Operations []*string `json:"operations,omitempty"`
}

// NsPrivilegeFieldHandles provides a string representation for each NsPrivilege field.
type NsPrivilegeFieldHandles struct {
	ObjectType string
	Operations string
}
