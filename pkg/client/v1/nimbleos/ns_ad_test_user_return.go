// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsADTestUserReturn - Active Directory user details.
// Export NsADTestUserReturnFields for advance operations like search filter etc.
var NsADTestUserReturnFields *NsADTestUserReturn

func init() {
	Usernamefield := "username"
	PrimaryGroupNamefield := "primary_group_name"
	PrimaryGroupIdfield := "primary_group_id"

	NsADTestUserReturnFields = &NsADTestUserReturn{
		Username:         &Usernamefield,
		PrimaryGroupName: &PrimaryGroupNamefield,
		PrimaryGroupId:   &PrimaryGroupIdfield,
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
