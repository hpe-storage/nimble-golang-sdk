// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// Role - Retrieve roles and privileges for role-based access control.
// Export RoleFields for advance operations like search filter etc.
var RoleFields *Role

func init() {

	RoleFields = &Role{
		ID:          "id",
		Name:        "name",
		FullName:    "full_name",
		Description: "description",
	}
}

type Role struct {
	// ID - Identifier for role.
	ID string `json:"id,omitempty"`
	// Name - Name of role.
	Name string `json:"name,omitempty"`
	// FullName - Full name of role.
	FullName string `json:"full_name,omitempty"`
	// Description - Description of role.
	Description string `json:"description,omitempty"`
	// PrivilegeList - List of privileges for this role.
	PrivilegeList []*NsPrivilege `json:"privilege_list,omitempty"`
	// ExtendedPrivilegeList - List of extended privileges for this role.
	ExtendedPrivilegeList []*NsExtendedPrivilege `json:"extended_privilege_list,omitempty"`
	// CreationTime - Time when this role was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this role was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// Hidden - Indicate whether the role is hidden.
	Hidden *bool `json:"hidden,omitempty"`
}

// sdk internal struct
type role struct {
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

// EncodeRole - Transform Role to role type
func EncodeRole(request interface{}) (*role, error) {
	reqRole := request.(*Role)
	byte, err := json.Marshal(reqRole)
	objPtr := &role{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeRole Transform role to Role type
func DecodeRole(request interface{}) (*Role, error) {
	reqRole := request.(*role)
	byte, err := json.Marshal(reqRole)
	obj := &Role{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
