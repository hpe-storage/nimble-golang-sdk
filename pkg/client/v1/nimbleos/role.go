// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export RoleFields provides field names to use in filter parameters, for example.
var RoleFields *RoleFieldHandles

func init() {
	RoleFields = &RoleFieldHandles{
		ID:                    "id",
		Name:                  "name",
		FullName:              "full_name",
		Description:           "description",
		PrivilegeList:         "privilege_list",
		ExtendedPrivilegeList: "extended_privilege_list",
		CreationTime:          "creation_time",
		LastModified:          "last_modified",
		Hidden:                "hidden",
	}
}

// Role - Retrieve roles and privileges for role-based access control.
type Role struct {
	// ID - Identifier for role.
	ID *string `json:"id,omitempty"`
	// Name - Name of role.
	Name *string `json:"name,omitempty"`
	// FullName - Full name of role.
	FullName *string `json:"full_name,omitempty"`
	// Description - Description of role.
	Description *string `json:"description,omitempty"`
	// PrivilegeList - List of privileges for this role.
	PrivilegeList []*NsPrivilege `json:"privilege_list,omitempty"`
	// ExtendedPrivilegeList - List of extended privileges for this role.
	ExtendedPrivilegeList []*NsExtendedPrivilege `json:"extended_privilege_list,omitempty"`
	// CreationTime - Time when this role was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this role was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// Hidden - Indicate whether the role is hidden.
	Hidden *bool `json:"hidden,omitempty"`
}

// RoleFieldHandles provides a string representation for each AccessControlRecord field.
type RoleFieldHandles struct {
	ID                    string
	Name                  string
	FullName              string
	Description           string
	PrivilegeList         string
	ExtendedPrivilegeList string
	CreationTime          string
	LastModified          string
	Hidden                string
}
