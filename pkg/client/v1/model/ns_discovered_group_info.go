// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsDiscoveredGroupInfo - Discovered group details.
// Export NsDiscoveredGroupInfoFields for advance operations like search filter etc.
var NsDiscoveredGroupInfoFields *NsDiscoveredGroupInfo

func init(){
	IDfield:= "id"
	Namefield:= "name"
	VersionCurrentfield:= "version_current"
	ManagementIpfield:= "management_ip"
	DiscoveryIpfield:= "discovery_ip"
		
	NsDiscoveredGroupInfoFields= &NsDiscoveredGroupInfo{
	ID: &IDfield,
	Name: &Namefield,
	VersionCurrent: &VersionCurrentfield,
	ManagementIp: &ManagementIpfield,
	DiscoveryIp: &DiscoveryIpfield,
		
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
