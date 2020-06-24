/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NetworkConfig - Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.
// Export NetworkConfigFields for advance operations like search filter etc.
var NetworkConfigFields *NetworkConfig

func init(){
	IDfield:= "id"
	MgmtIpfield:= "mgmt_ip"
	SecondaryMgmtIpfield:= "secondary_mgmt_ip"
	GroupLeaderArrayfield:= "group_leader_array"
		
	NetworkConfigFields= &NetworkConfig{
		ID: &IDfield,
		MgmtIp: &MgmtIpfield,
		SecondaryMgmtIp: &SecondaryMgmtIpfield,
		GroupLeaderArray: &GroupLeaderArrayfield,
		
	}
}

type NetworkConfig struct {
   
    // Identifier for network configuration.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the network configuration. Use the name 'draft' when creating a draft configuration.
    
   	Name *NsNetConfigName `json:"name,omitempty"`
   
    // Management IP address for the Group.
    
 	MgmtIp *string `json:"mgmt_ip,omitempty"`
   
    // Secondary management IP address for the Group.
    
 	SecondaryMgmtIp *string `json:"secondary_mgmt_ip,omitempty"`
   
    // Role of network configuration.
    
   	Role *NsNetConfigRole `json:"role,omitempty"`
   
    // Whether automatic connection method is enabled. Enabling this means means redirecting connections from the specified iSCSI discovery IP address to the best data IP address based on connection counts.
    
 	IscsiAutomaticConnectionMethod *bool `json:"iscsi_automatic_connection_method,omitempty"`
   
    // Whether rebalancing is enabled. Enabling this means rebalancing iSCSI connections by periodically breaking existing connections that are out-of-balance, allowing the host to reconnect to a more appropriate data IP address.
    
 	IscsiConnectionRebalancing *bool `json:"iscsi_connection_rebalancing,omitempty"`
   
    // List of static routes.
    
   	RouteList []*NsRoute `json:"route_list,omitempty"`
   
    // List of subnet configs.
    
   	SubnetList []*NsSubnet `json:"subnet_list,omitempty"`
   
    // List of array network configs.
    
   	ArrayList []*NsArrayNet `json:"array_list,omitempty"`
   
    // Name of the group leader array.
    
 	GroupLeaderArray *string `json:"group_leader_array,omitempty"`
   
    // Time when this net configuration was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this network configuration was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Start time of activity.
    
   	ActiveSince *int64 `json:"active_since,omitempty"`
   
    // Time of last activity.
    
   	LastActive *int64 `json:"last_active,omitempty"`
   
    // Indicates whether to ignore the validation.
    
   	IgnoreValIDationMask *int64 `json:"ignore_validation_mask,omitempty"`
}
