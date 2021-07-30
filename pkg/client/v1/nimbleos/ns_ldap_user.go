// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsLdapUser - Information about the current status of an LDAP user.

// Export NsLdapUserFields provides field names to use in filter parameters, for example.
var NsLdapUserFields *NsLdapUserStringFields

func init() {
	fieldUser := "user"
	fieldPrimaryGroupName := "primary_group_name"
	fieldPrimaryGroupId := "primary_group_id"
	fieldGroupCount := "group_count"
	fieldGroups := "groups"
	fieldRole := "role"

	NsLdapUserFields = &NsLdapUserStringFields{
		User:             &fieldUser,
		PrimaryGroupName: &fieldPrimaryGroupName,
		PrimaryGroupId:   &fieldPrimaryGroupId,
		GroupCount:       &fieldGroupCount,
		Groups:           &fieldGroups,
		Role:             &fieldRole,
	}
}

type NsLdapUser struct {
	// User - The user's username.
	User *string `json:"user,omitempty"`
	// PrimaryGroupName - The name of the configured LDAP group used to determine the user's role.
	PrimaryGroupName *string `json:"primary_group_name,omitempty"`
	// PrimaryGroupId - The external identifier for the configured LDAP group uset to determine the user's role.
	PrimaryGroupId *string `json:"primary_group_id,omitempty"`
	// GroupCount - How many configured LDAP groups is the user is currently reported as having membership in.
	GroupCount *int64 `json:"group_count,omitempty"`
	// Groups - A list of all the configured LDAP groups the user is currently reported to have membership in.
	Groups []*string `json:"groups,omitempty"`
	// Role - The role currently available to the user.
	Role *NsUserRoles `json:"role,omitempty"`
}

// Struct for NsLdapUserFields
type NsLdapUserStringFields struct {
	User             *string
	PrimaryGroupName *string
	PrimaryGroupId   *string
	GroupCount       *string
	Groups           *string
	Role             *string
}
