// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// UserGroup - Represents Active Directory groups configured to manage the system.
// Export UserGroupFields for advance operations like search filter etc.
var UserGroupFields *UserGroup

func init() {

	UserGroupFields = &UserGroup{
		ID:          "id",
		Name:        "name",
		Description: "description",
		RoleId:      "role_id",
		ExternalId:  "external_id",
		DomainId:    "domain_id",
		DomainName:  "domain_name",
	}
}

type UserGroup struct {
	// ID - Identifier for the user group.
	ID string `json:"id,omitempty"`
	// Name - Name of the user group.
	Name string `json:"name,omitempty"`
	// Description - Description of the user group.
	Description string `json:"description,omitempty"`
	// RoleId - Identifier for the user group's role.
	RoleId string `json:"role_id,omitempty"`
	// Role - Role of the user.
	Role *NsUserRoles `json:"role,omitempty"`
	// InactivityTimeout - The amount of time that the user session is inactive before timing out. A value of 0 indicates that the timeout is taken from the group setting.
	InactivityTimeout int64 `json:"inactivity_timeout,omitempty"`
	// CreationTime - Time when this user was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this user was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// Disabled - User is currently disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// ExternalId - External ID of the user group. In Active Directory, it is the group's SID (Security Identifier).
	ExternalId string `json:"external_id,omitempty"`
	// DomainId - Identifier of the domain this user group belongs to.
	DomainId string `json:"domain_id,omitempty"`
	// DomainName - Role of the user.
	DomainName string `json:"domain_name,omitempty"`
}

// sdk internal struct
type userGroup struct {
	ID                *string      `json:"id,omitempty"`
	Name              *string      `json:"name,omitempty"`
	Description       *string      `json:"description,omitempty"`
	RoleId            *string      `json:"role_id,omitempty"`
	Role              *NsUserRoles `json:"role,omitempty"`
	InactivityTimeout *int64       `json:"inactivity_timeout,omitempty"`
	CreationTime      *int64       `json:"creation_time,omitempty"`
	LastModified      *int64       `json:"last_modified,omitempty"`
	Disabled          *bool        `json:"disabled,omitempty"`
	ExternalId        *string      `json:"external_id,omitempty"`
	DomainId          *string      `json:"domain_id,omitempty"`
	DomainName        *string      `json:"domain_name,omitempty"`
}

// EncodeUserGroup - Transform UserGroup to userGroup type
func EncodeUserGroup(request interface{}) (*userGroup, error) {
	reqUserGroup := request.(*UserGroup)
	byte, err := json.Marshal(reqUserGroup)
	if err != nil {
		return nil, err
	}
	respUserGroupPtr := &userGroup{}
	err = json.Unmarshal(byte, respUserGroupPtr)
	return respUserGroupPtr, err
}

// DecodeUserGroup Transform userGroup to UserGroup type
func DecodeUserGroup(request interface{}) (*UserGroup, error) {
	reqUserGroup := request.(*userGroup)
	byte, err := json.Marshal(reqUserGroup)
	if err != nil {
		return nil, err
	}
	respUserGroup := &UserGroup{}
	err = json.Unmarshal(byte, respUserGroup)
	return respUserGroup, err
}
