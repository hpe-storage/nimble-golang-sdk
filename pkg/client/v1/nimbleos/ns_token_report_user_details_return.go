// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTokenReportUserDetailsReturn - Return values of token reporting user details.

// Export NsTokenReportUserDetailsReturnFields provides field names to use in filter parameters, for example.
var NsTokenReportUserDetailsReturnFields *NsTokenReportUserDetailsReturnStringFields

func init() {
	fieldUserName := "user_name"
	fieldPrimaryGroupId := "primary_group_id"
	fieldPrimaryGroupName := "primary_group_name"
	fieldGroupCount := "group_count"
	fieldRole := "role"
	fieldInactivityTimeout := "inactivity_timeout"
	fieldUserId := "user_id"
	fieldDomainType := "domain_type"
	fieldGroups := "groups"

	NsTokenReportUserDetailsReturnFields = &NsTokenReportUserDetailsReturnStringFields{
		UserName:          &fieldUserName,
		PrimaryGroupId:    &fieldPrimaryGroupId,
		PrimaryGroupName:  &fieldPrimaryGroupName,
		GroupCount:        &fieldGroupCount,
		Role:              &fieldRole,
		InactivityTimeout: &fieldInactivityTimeout,
		UserId:            &fieldUserId,
		DomainType:        &fieldDomainType,
		Groups:            &fieldGroups,
	}
}

type NsTokenReportUserDetailsReturn struct {
	// UserName - User name for the session.
	UserName *string `json:"user_name,omitempty"`
	// PrimaryGroupId - The ID of the primary directory service user group the user belongs to. RBAC is granted based on this group.
	PrimaryGroupId *string `json:"primary_group_id,omitempty"`
	// PrimaryGroupName - The primary directory user group the user belongs to. RBAC is granted based on this group.
	PrimaryGroupName *string `json:"primary_group_name,omitempty"`
	// GroupCount - The number of directory service user groups the user belongs to.
	GroupCount *int64 `json:"group_count,omitempty"`
	// Role - Role of the user.
	Role *NsUserRoles `json:"role,omitempty"`
	// InactivityTimeout - The amount of time that the user session is inactive before timing out.
	InactivityTimeout *int64 `json:"inactivity_timeout,omitempty"`
	// UserId - Global ID of the user.
	UserId *string `json:"user_id,omitempty"`
	// DomainType - Connected directory service type. Either 'ad, 'ldap' or 'local'.
	DomainType *NsDomainType `json:"domain_type,omitempty"`
	// Groups - The list of directory service groups the user belongs to.
	Groups []*string `json:"groups,omitempty"`
}

// Struct for NsTokenReportUserDetailsReturnFields
type NsTokenReportUserDetailsReturnStringFields struct {
	UserName          *string
	PrimaryGroupId    *string
	PrimaryGroupName  *string
	GroupCount        *string
	Role              *string
	InactivityTimeout *string
	UserId            *string
	DomainType        *string
	Groups            *string
}
