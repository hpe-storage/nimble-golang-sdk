// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// UserGroupFields provides field names to use in filter parameters, for example.
var UserGroupFields *UserGroupFieldHandles

func init() {
	UserGroupFields = &UserGroupFieldHandles{
		ID:                "id",
		Name:              "name",
		Description:       "description",
		RoleId:            "role_id",
		Role:              "role",
		InactivityTimeout: "inactivity_timeout",
		CreationTime:      "creation_time",
		LastModified:      "last_modified",
		Disabled:          "disabled",
		ExternalId:        "external_id",
		DomainId:          "domain_id",
		DomainName:        "domain_name",
	}
}

// UserGroup - Represents Active Directory groups configured to manage the system.
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

// UserGroupFieldHandles provides a string representation for each AccessControlRecord field.
type UserGroupFieldHandles struct {
	ID                string
	Name              string
	Description       string
	RoleId            string
	Role              string
	InactivityTimeout string
	CreationTime      string
	LastModified      string
	Disabled          string
	ExternalId        string
	DomainId          string
	DomainName        string
}
