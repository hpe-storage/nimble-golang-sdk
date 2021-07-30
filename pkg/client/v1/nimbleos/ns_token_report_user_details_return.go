// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTokenReportUserDetailsReturn - Return values of token reporting user details.
// Export NsTokenReportUserDetailsReturnFields for advance operations like search filter etc.
var NsTokenReportUserDetailsReturnFields *NsTokenReportUserDetailsReturnStringFields

func init() {
	UserNamefield := "user_name"
	PrimaryGroupIdfield := "primary_group_id"
	PrimaryGroupNamefield := "primary_group_name"
	GroupCountfield := "group_count"
	Rolefield := "role"
	InactivityTimeoutfield := "inactivity_timeout"
	UserIdfield := "user_id"
	DomainTypefield := "domain_type"
	Groupsfield := "groups"

	NsTokenReportUserDetailsReturnFields = &NsTokenReportUserDetailsReturnStringFields{
		UserName:          &UserNamefield,
		PrimaryGroupId:    &PrimaryGroupIdfield,
		PrimaryGroupName:  &PrimaryGroupNamefield,
		GroupCount:        &GroupCountfield,
		Role:              &Rolefield,
		InactivityTimeout: &InactivityTimeoutfield,
		UserId:            &UserIdfield,
		DomainType:        &DomainTypefield,
		Groups:            &Groupsfield,
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
