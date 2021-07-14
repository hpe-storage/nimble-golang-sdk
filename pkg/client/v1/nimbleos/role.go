// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// Role - Retrieve roles and privileges for role-based access control.
// Export RoleFields for advance operations like search filter etc.
var RoleFields *Role

func init(){
 IDfield:= "id"
 Namefield:= "name"
 FullNamefield:= "full_name"
 Descriptionfield:= "description"

 RoleFields= &Role{
  ID:                     &IDfield,
  Name:                   &Namefield,
  FullName:               &FullNamefield,
  Description:            &Descriptionfield,
 }
}


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
