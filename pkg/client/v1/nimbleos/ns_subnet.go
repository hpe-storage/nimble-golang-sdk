// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSubnet - A subnet configuration.
// Export NsSubnetFields for advance operations like search filter etc.
var NsSubnetFields *NsSubnet

func init() {

	NsSubnetFields = &NsSubnet{
		Label:       "label",
		Network:     "network",
		Netmask:     "netmask",
		DiscoveryIp: "discovery_ip",
	}
}

type NsSubnet struct {
	// Label - Subnet label.
	Label string `json:"label,omitempty"`
	// Network - Network IP address.
	Network string `json:"network,omitempty"`
	// Netmask - Subnet netmask address.
	Netmask string `json:"netmask,omitempty"`
	// NetzoneType - Netzone type.
	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
	// Type - Subnet type.
	Type *NsSubnetType `json:"type,omitempty"`
	// AllowIscsi - Allow iSCSI.
	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
	// AllowGroup - Allow group.
	AllowGroup *bool `json:"allow_group,omitempty"`
	// DiscoveryIp - Discovery IP address.
	DiscoveryIp string `json:"discovery_ip,omitempty"`
	// Mtu - MTU for specified subnet.
	Mtu int64 `json:"mtu,omitempty"`
	// VlanId - VLAN ID for specified subnet.
	VlanId int64 `json:"vlan_id,omitempty"`
	// Failover - Failover setting of the subnet.
	Failover *bool `json:"failover,omitempty"`
	// FailoverEnableTime - Failover for this subnet will be enabled again at the time specified by failover_enable_time.
	FailoverEnableTime int64 `json:"failover_enable_time,omitempty"`
}

// sdk internal struct
type nsSubnet struct {
	// Label - Subnet label.
	Label *string `json:"label,omitempty"`
	// Network - Network IP address.
	Network *string `json:"network,omitempty"`
	// Netmask - Subnet netmask address.
	Netmask *string `json:"netmask,omitempty"`
	// NetzoneType - Netzone type.
	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
	// Type - Subnet type.
	Type *NsSubnetType `json:"type,omitempty"`
	// AllowIscsi - Allow iSCSI.
	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
	// AllowGroup - Allow group.
	AllowGroup *bool `json:"allow_group,omitempty"`
	// DiscoveryIp - Discovery IP address.
	DiscoveryIp *string `json:"discovery_ip,omitempty"`
	// Mtu - MTU for specified subnet.
	Mtu *int64 `json:"mtu,omitempty"`
	// VlanId - VLAN ID for specified subnet.
	VlanId *int64 `json:"vlan_id,omitempty"`
	// Failover - Failover setting of the subnet.
	Failover *bool `json:"failover,omitempty"`
	// FailoverEnableTime - Failover for this subnet will be enabled again at the time specified by failover_enable_time.
	FailoverEnableTime *int64 `json:"failover_enable_time,omitempty"`
}

// EncodeNsSubnet - Transform NsSubnet to nsSubnet type
func EncodeNsSubnet(request interface{}) (*nsSubnet, error) {
	reqNsSubnet := request.(*NsSubnet)
	byte, err := json.Marshal(reqNsSubnet)
	objPtr := &nsSubnet{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSubnet Transform nsSubnet to NsSubnet type
func DecodeNsSubnet(request interface{}) (*NsSubnet, error) {
	reqNsSubnet := request.(*nsSubnet)
	byte, err := json.Marshal(reqNsSubnet)
	obj := &NsSubnet{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
