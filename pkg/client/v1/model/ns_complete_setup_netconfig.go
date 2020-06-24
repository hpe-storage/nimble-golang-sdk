/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsCompleteSetupNetconfig - Information includes group, array and network configuration in complete setup.
// Export NsCompleteSetupNetconfigFields for advance operations like search filter etc.
var NsCompleteSetupNetconfigFields *NsCompleteSetupNetconfig

func init(){
	DiscoveryIpfield:= "discovery_ip"
	DiscoveryNetmaskfield:= "discovery_netmask"
	MgmtIpfield:= "mgmt_ip"
	SecondaryMgmtIpfield:= "secondary_mgmt_ip"
	MgmtNetmaskfield:= "mgmt_netmask"
		
	NsCompleteSetupNetconfigFields= &NsCompleteSetupNetconfig{
		DiscoveryIp: &DiscoveryIpfield,
		DiscoveryNetmask: &DiscoveryNetmaskfield,
		MgmtIp: &MgmtIpfield,
		SecondaryMgmtIp: &SecondaryMgmtIpfield,
		MgmtNetmask: &MgmtNetmaskfield,
		
	}
}

type NsCompleteSetupNetconfig struct {
   
    // Network configuration name, possible values: 'active', 'backup', 'draft'.
    
   	Name *NsNetConfigName `json:"name,omitempty"`
   
    // The discovery IP address of the group.
    
 	DiscoveryIp *string `json:"discovery_ip,omitempty"`
   
    // The netmask of the discovery IP address of the group.
    
 	DiscoveryNetmask *string `json:"discovery_netmask,omitempty"`
   
    // The management IP address of the group.
    
 	MgmtIp *string `json:"mgmt_ip,omitempty"`
   
    // The secondary management IP address of the group.
    
 	SecondaryMgmtIp *string `json:"secondary_mgmt_ip,omitempty"`
   
    // The netmask of the management IP address of the group.
    
 	MgmtNetmask *string `json:"mgmt_netmask,omitempty"`
   
    // Iscsi has atomatic connection method.
    
 	IscsiAutomaticConnectionMethod *bool `json:"iscsi_automatic_connection_method,omitempty"`
   
    // Iscsi connection rebalanced.
    
 	IscsiConnectionRebalancing *bool `json:"iscsi_connection_rebalancing,omitempty"`
   
    // Network route list.
    
   	RouteList []*NsRoute `json:"route_list,omitempty"`
   
    // Subnet list.
    
   	SubnetList []*NsSubnet `json:"subnet_list,omitempty"`
   
    // Array list.
    
   	ArrayList []*NsArrayNet `json:"array_list,omitempty"`
}
