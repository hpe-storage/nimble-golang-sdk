// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsDiscoveredGroupInfo - Discovered group details.

// Export NsDiscoveredGroupInfoFields provides field names to use in filter parameters, for example.
var NsDiscoveredGroupInfoFields *NsDiscoveredGroupInfoStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldVersionCurrent := "version_current"
	fieldCountOfMembers := "count_of_members"
	fieldManagementIp := "management_ip"
	fieldDiscoveryIp := "discovery_ip"
	fieldIscsiEnabled := "iscsi_enabled"
	fieldFcEnabled := "fc_enabled"
	fieldArrayMemberList := "array_member_list"

	NsDiscoveredGroupInfoFields = &NsDiscoveredGroupInfoStringFields{
		ID:              &fieldID,
		Name:            &fieldName,
		VersionCurrent:  &fieldVersionCurrent,
		CountOfMembers:  &fieldCountOfMembers,
		ManagementIp:    &fieldManagementIp,
		DiscoveryIp:     &fieldDiscoveryIp,
		IscsiEnabled:    &fieldIscsiEnabled,
		FcEnabled:       &fieldFcEnabled,
		ArrayMemberList: &fieldArrayMemberList,
	}
}

type NsDiscoveredGroupInfo struct {
	// ID - Identifier of the group.
	ID *string `json:"id,omitempty"`
	// Name - Name of the group.
	Name *string `json:"name,omitempty"`
	// VersionCurrent - Version of software running on the group.
	VersionCurrent *string `json:"version_current,omitempty"`
	// CountOfMembers - Number of member arrays in the group.
	CountOfMembers *int64 `json:"count_of_members,omitempty"`
	// ManagementIp - The management IP address of the group.
	ManagementIp *string `json:"management_ip,omitempty"`
	// DiscoveryIp - The discovery IP address of the group.
	DiscoveryIp *string `json:"discovery_ip,omitempty"`
	// IscsiEnabled - Whether iscsi is enabled on the group.
	IscsiEnabled *bool `json:"iscsi_enabled,omitempty"`
	// FcEnabled - Whether fibre channel is enabled on the group.
	FcEnabled *bool `json:"fc_enabled,omitempty"`
	// ArrayMemberList - Members of this group.
	ArrayMemberList []*NsArraySummaryInfo `json:"array_member_list,omitempty"`
}

// Struct for NsDiscoveredGroupInfoFields
type NsDiscoveredGroupInfoStringFields struct {
	ID              *string
	Name            *string
	VersionCurrent  *string
	CountOfMembers  *string
	ManagementIp    *string
	DiscoveryIp     *string
	IscsiEnabled    *string
	FcEnabled       *string
	ArrayMemberList *string
}
