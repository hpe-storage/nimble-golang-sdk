/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Subnet - Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets let you create logical addressing for selective routing.
// Export SubnetFields for advance operations like search filter etc.
var SubnetFields *Subnet

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Networkfield:= "network"
	Netmaskfield:= "netmask"
	DiscoveryIpfield:= "discovery_ip"
		
	SubnetFields= &Subnet{
		ID: &IDfield,
		Name: &Namefield,
		Network: &Networkfield,
		Netmask: &Netmaskfield,
		DiscoveryIp: &DiscoveryIpfield,
		
	}
}

type Subnet struct {
   
    // Identifier for the initiator group.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of subnet configuration.
    
 	Name *string `json:"name,omitempty"`
   
    // Subnet network address.
    
 	Network *string `json:"network,omitempty"`
   
    // Subnet netmask address.
    
 	Netmask *string `json:"netmask,omitempty"`
   
    // Subnet type. Options include 'mgmt', 'data', and 'mgmt,data'.
    
   	Type *NsSubnetType `json:"type,omitempty"`
   
    // Subnet type.
    
 	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
   
    // Subnet type.
    
 	AllowGroup *bool `json:"allow_group,omitempty"`
   
    // Subnet network address.
    
 	DiscoveryIp *string `json:"discovery_ip,omitempty"`
   
    // MTU for specified subnet. Valid MTU's are in the 512-16000 range.
    
   	Mtu *int64 `json:"mtu,omitempty"`
   
    // Specify Network Affinity Zone type for iSCSI enabled subnets. Valid types are Single, Bisect, and EvenOdd for iSCSI subnets.
    
   	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
   
    // VLAN ID for specified subnet. Valid ID's are in the 1-4094 range.
    
   	VlanID *int64 `json:"vlan_id,omitempty"`
   
    // Time when this subnet configuration was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this subnet configuration was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Failover setting of the subnet.
    
 	Failover *bool `json:"failover,omitempty"`
   
    // Failover for this subnet will be enabled again at the time specified by failover_enable_time.
    
   	FailoverEnableTime *int64 `json:"failover_enable_time,omitempty"`
}
