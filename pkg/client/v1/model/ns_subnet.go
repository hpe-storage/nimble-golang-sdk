/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSubnet - A subnet configuration.
// Export NsSubnetFields for advance operations like search filter etc.
var NsSubnetFields *NsSubnet

func init(){
	Labelfield:= "label"
	Networkfield:= "network"
	Netmaskfield:= "netmask"
	DiscoveryIpfield:= "discovery_ip"
		
	NsSubnetFields= &NsSubnet{
		Label: &Labelfield,
		Network: &Networkfield,
		Netmask: &Netmaskfield,
		DiscoveryIp: &DiscoveryIpfield,
		
	}
}

type NsSubnet struct {
   
    // Subnet label.
    
 	Label *string `json:"label,omitempty"`
   
    // Network IP address.
    
 	Network *string `json:"network,omitempty"`
   
    // Subnet netmask address.
    
 	Netmask *string `json:"netmask,omitempty"`
   
    // Netzone type.
    
   	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
   
    // Subnet type.
    
   	Type *NsSubnetType `json:"type,omitempty"`
   
    // Allow iSCSI.
    
 	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
   
    // Allow group.
    
 	AllowGroup *bool `json:"allow_group,omitempty"`
   
    // Discovery IP address.
    
 	DiscoveryIp *string `json:"discovery_ip,omitempty"`
   
    // MTU for specified subnet.
    
   	Mtu *int64 `json:"mtu,omitempty"`
   
    // VLAN ID for specified subnet.
    
   	VlanID *int64 `json:"vlan_id,omitempty"`
   
    // Failover setting of the subnet.
    
 	Failover *bool `json:"failover,omitempty"`
   
    // Failover for this subnet will be enabled again at the time specified by failover_enable_time.
    
   	FailoverEnableTime *int64 `json:"failover_enable_time,omitempty"`
}
