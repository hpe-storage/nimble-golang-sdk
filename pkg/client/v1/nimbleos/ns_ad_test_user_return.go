// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsADTestUserReturn - Active Directory user details.

// Export NsADTestUserReturnFields provides field names to use in filter parameters, for example.
var NsADTestUserReturnFields *NsADTestUserReturnStringFields

func init() {
	fieldUsername := "username"
	fieldPrimaryGroupName := "primary_group_name"
	fieldPrimaryGroupId := "primary_group_id"
	fieldGroupCount := "group_count"
	fieldGroups := "groups"
	fieldRole := "role"

	NsADTestUserReturnFields = &NsADTestUserReturnStringFields{
		Username:         &fieldUsername,
		PrimaryGroupName: &fieldPrimaryGroupName,
		PrimaryGroupId:   &fieldPrimaryGroupId,
		GroupCount:       &fieldGroupCount,
		Groups:           &fieldGroups,
		Role:             &fieldRole,
	}
}

type NsADTestUserReturn struct {
	// Username - Name of the Active Directory user.
	Username *string `json:"username,omitempty"`
	// PrimaryGroupName - Name of the Active Directory group the user belongs to. RBAC is based on this primary group.
	PrimaryGroupName *string `json:"primary_group_name,omitempty"`
	// PrimaryGroupId - ID of the Active Directory group the user belongs to. RBAC is based on this primary group.
	PrimaryGroupId *string `json:"primary_group_id,omitempty"`
	// GroupCount - Number of Active Directory groups the user belongs to.
	GroupCount *int64 `json:"group_count,omitempty"`
	// Groups - List of Active Directory groups the user belongs to.
	Groups []*string `json:"groups,omitempty"`
	// Role - The role the user belongs to.
	Role *NsUserRoles `json:"role,omitempty"`
}

// Struct for NsADTestUserReturnFields
type NsADTestUserReturnStringFields struct {
	Username         *string
	PrimaryGroupName *string
	PrimaryGroupId   *string
	GroupCount       *string
	Groups           *string
	Role             *string
}
