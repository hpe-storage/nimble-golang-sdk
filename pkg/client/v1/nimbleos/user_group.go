// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// UserGroup - Represents Active Directory groups configured to manage the system.
// Export UserGroupFields for advance operations like search filter etc.
var UserGroupFields *UserGroupStringFields

func init() {
	IDfield := "id"
	Namefield := "name"
	Descriptionfield := "description"
	RoleIdfield := "role_id"
	Rolefield := "role"
	InactivityTimeoutfield := "inactivity_timeout"
	CreationTimefield := "creation_time"
	LastModifiedfield := "last_modified"
	Disabledfield := "disabled"
	ExternalIdfield := "external_id"
	DomainIdfield := "domain_id"
	DomainNamefield := "domain_name"

	UserGroupFields = &UserGroupStringFields{
		ID:                &IDfield,
		Name:              &Namefield,
		Description:       &Descriptionfield,
		RoleId:            &RoleIdfield,
		Role:              &Rolefield,
		InactivityTimeout: &InactivityTimeoutfield,
		CreationTime:      &CreationTimefield,
		LastModified:      &LastModifiedfield,
		Disabled:          &Disabledfield,
		ExternalId:        &ExternalIdfield,
		DomainId:          &DomainIdfield,
		DomainName:        &DomainNamefield,
	}
}

type UserGroup struct {
	// ID - Identifier for the user group.
	ID *string `json:"id,omitempty"`
	// Name - Name of the user group.
	Name *string `json:"name,omitempty"`
	// Description - Description of the user group.
	Description *string `json:"description,omitempty"`
	// RoleId - Identifier for the user group's role.
	RoleId *string `json:"role_id,omitempty"`
	// Role - Role of the user.
	Role *NsUserRoles `json:"role,omitempty"`
	// InactivityTimeout - The amount of time that the user session is inactive before timing out. A value of 0 indicates that the timeout is taken from the group setting.
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`
	// CreationTime - Time when this user was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this user was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// Disabled - User is currently disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// ExternalId - External ID of the user group. In Active Directory, it is the group's SID (Security Identifier).
	ExternalId *string `json:"external_id,omitempty"`
	// DomainId - Identifier of the domain this user group belongs to.
	DomainId *string `json:"domain_id,omitempty"`
	// DomainName - Role of the user.
	DomainName *string `json:"domain_name,omitempty"`
}

// Struct for UserGroupFields
type UserGroupStringFields struct {
	ID                *string
	Name              *string
	Description       *string
	RoleId            *string
	Role              *string
	InactivityTimeout *string
	CreationTime      *string
	LastModified      *string
	Disabled          *string
	ExternalId        *string
	DomainId          *string
	DomainName        *string
}
