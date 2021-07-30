// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Subnet - Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets let you create logical addressing for selective routing.
// Export SubnetFields for advance operations like search filter etc.
var SubnetFields *SubnetStringFields

func init() {
	IDfield := "id"
	Namefield := "name"
	Networkfield := "network"
	Netmaskfield := "netmask"
	Typefield := "type"
	AllowIscsifield := "allow_iscsi"
	AllowGroupfield := "allow_group"
	DiscoveryIpfield := "discovery_ip"
	Mtufield := "mtu"
	NetzoneTypefield := "netzone_type"
	VlanIdfield := "vlan_id"
	CreationTimefield := "creation_time"
	LastModifiedfield := "last_modified"
	Failoverfield := "failover"
	FailoverEnableTimefield := "failover_enable_time"

	SubnetFields = &SubnetStringFields{
		ID:                 &IDfield,
		Name:               &Namefield,
		Network:            &Networkfield,
		Netmask:            &Netmaskfield,
		Type:               &Typefield,
		AllowIscsi:         &AllowIscsifield,
		AllowGroup:         &AllowGroupfield,
		DiscoveryIp:        &DiscoveryIpfield,
		Mtu:                &Mtufield,
		NetzoneType:        &NetzoneTypefield,
		VlanId:             &VlanIdfield,
		CreationTime:       &CreationTimefield,
		LastModified:       &LastModifiedfield,
		Failover:           &Failoverfield,
		FailoverEnableTime: &FailoverEnableTimefield,
	}
}

type Subnet struct {
	// ID - Identifier for the initiator group.
	ID *string `json:"id,omitempty"`
	// Name - Name of subnet configuration.
	Name *string `json:"name,omitempty"`
	// Network - Subnet network address.
	Network *string `json:"network,omitempty"`
	// Netmask - Subnet netmask address.
	Netmask *string `json:"netmask,omitempty"`
	// Type - Subnet type. Options include 'mgmt', 'data', and 'mgmt,data'.
	Type *NsSubnetType `json:"type,omitempty"`
	// AllowIscsi - Subnet type.
	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
	// AllowGroup - Subnet type.
	AllowGroup *bool `json:"allow_group,omitempty"`
	// DiscoveryIp - Subnet network address.
	DiscoveryIp *string `json:"discovery_ip,omitempty"`
	// Mtu - MTU for specified subnet. Valid MTU's are in the 512-16000 range.
	Mtu *int64 `json:"mtu,omitempty"`
	// NetzoneType - Specify Network Affinity Zone type for iSCSI enabled subnets. Valid types are Single, Bisect, and EvenOdd for iSCSI subnets.
	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
	// VlanId - VLAN ID for specified subnet. Valid ID's are in the 1-4094 range.
	VlanId *int64 `json:"vlan_id,omitempty"`
	// CreationTime - Time when this subnet configuration was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this subnet configuration was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// Failover - Failover setting of the subnet.
	Failover *bool `json:"failover,omitempty"`
	// FailoverEnableTime - Failover for this subnet will be enabled again at the time specified by failover_enable_time.
	FailoverEnableTime *int64 `json:"failover_enable_time,omitempty"`
}

// Struct for SubnetFields
type SubnetStringFields struct {
	ID                 *string
	Name               *string
	Network            *string
	Netmask            *string
	Type               *string
	AllowIscsi         *string
	AllowGroup         *string
	DiscoveryIp        *string
	Mtu                *string
	NetzoneType        *string
	VlanId             *string
	CreationTime       *string
	LastModified       *string
	Failover           *string
	FailoverEnableTime *string
}
