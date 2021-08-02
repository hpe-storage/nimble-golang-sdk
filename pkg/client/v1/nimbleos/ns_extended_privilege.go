// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsExtendedPrivilegeFields provides field names to use in filter parameters, for example.
var NsExtendedPrivilegeFields *NsExtendedPrivilegeFieldHandles

func init() {
	NsExtendedPrivilegeFields = &NsExtendedPrivilegeFieldHandles{
		ObjectType: "object_type",
		Operation:  "operation",
		Allow:      "allow",
	}
}

// NsExtendedPrivilege - Extended privilege info.
type NsExtendedPrivilege struct {
	// ObjectType - Object type name associated with this privilege.
	ObjectType *string `json:"object_type,omitempty"`
	// Operation - Operation associated with the above object for this privilege.
	Operation *string `json:"operation,omitempty"`
	// Allow - Indicate whether the above operation is allowed for this privilege.
	Allow *bool `json:"allow,omitempty"`
}

// NsExtendedPrivilegeFieldHandles provides a string representation for each NsExtendedPrivilege field.
type NsExtendedPrivilegeFieldHandles struct {
	ObjectType string
	Operation  string
	Allow      string
}
