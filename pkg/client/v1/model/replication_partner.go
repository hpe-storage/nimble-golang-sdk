/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// ReplicationPartner - Manage replication partner. Replication partners let one storage array talk to another for replication purposes. The two arrays must be able to communicate over a network, and use ports 4213 and 4214. Replication partners have the same name as the remote group. Replication partners can be reciprocal, upstream (the source of replicas), or downstream (the receiver of replicas) partners.
// Export ReplicationPartnerFields for advance operations like search filter etc.
var ReplicationPartnerFields *ReplicationPartner

func init(){
	IDfield:= "id"
	Namefield:= "name"
	FullNamefield:= "full_name"
	SearchNamefield:= "search_name"
	Descriptionfield:= "description"
	Aliasfield:= "alias"
	Secretfield:= "secret"
	Hostnamefield:= "hostname"
	ProxyHostnamefield:= "proxy_hostname"
	ProxyUserfield:= "proxy_user"
	ReplHostnamefield:= "repl_hostname"
	LastKeepaliveErrorfield:= "last_keepalive_error"
	LastSyncErrorfield:= "last_sync_error"
	ArraySerialfield:= "array_serial"
	PoolIDfield:= "pool_id"
	PoolNamefield:= "pool_name"
	FolderIDfield:= "folder_id"
	FolderNamefield:= "folder_name"
	SubnetLabelfield:= "subnet_label"
	SubnetNetworkfield:= "subnet_network"
	SubnetNetmaskfield:= "subnet_netmask"
		
	ReplicationPartnerFields= &ReplicationPartner{
		ID: &IDfield,
		Name: &Namefield,
		FullName: &FullNamefield,
		SearchName: &SearchNamefield,
		Description: &Descriptionfield,
		Alias: &Aliasfield,
		Secret: &Secretfield,
		Hostname: &Hostnamefield,
		ProxyHostname: &ProxyHostnamefield,
		ProxyUser: &ProxyUserfield,
		ReplHostname: &ReplHostnamefield,
		LastKeepaliveError: &LastKeepaliveErrorfield,
		LastSyncError: &LastSyncErrorfield,
		ArraySerial: &ArraySerialfield,
		PoolID: &PoolIDfield,
		PoolName: &PoolNamefield,
		FolderID: &FolderIDfield,
		FolderName: &FolderNamefield,
		SubnetLabel: &SubnetLabelfield,
		SubnetNetwork: &SubnetNetworkfield,
		SubnetNetmask: &SubnetNetmaskfield,
		
	}
}

type ReplicationPartner struct {
   
    // Identifier for a replication partner.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of replication partner.
    
 	Name *string `json:"name,omitempty"`
   
    // Fully qualified name of replication partner.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Name of replication partner used for object search.
    
 	SearchName *string `json:"search_name,omitempty"`
   
    // Description of replication partner.
    
 	Description *string `json:"description,omitempty"`
   
    // Replication partner type. Possible values are group or pool.
    
   	PartnerType *NsPartnerType `json:"partner_type,omitempty"`
   
    // Name this group uses to identify itself to this partner.
    
 	Alias *string `json:"alias,omitempty"`
   
    // Replication partner shared secret, used for mutual authentication of the partners.
    
 	Secret *string `json:"secret,omitempty"`
   
    // Time when this replication partner was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this replication partner was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Port number of partner control interface.
    
  	ControlPort  *int64 `json:"control_port,omitempty"`
   
    // IP address or hostname of partner interface.  This must be the partner's Group Management IP address.
    
 	Hostname *string `json:"hostname,omitempty"`
   
    // For tunnel_endpoint partner types, first port available on the ssh proxy available for reverse forwarding. It must be guaranteed that the proxy has the next N ports reserved for this partner, where N is the count of DSDs in this group. This attribute is only valid for tunnel_endpoint partner type.
    
   	PortRangeStart *int64 `json:"port_range_start,omitempty"`
   
    // IP address of tunnel endpoint. Only valid for tunnel_endpoint partner types.
    
 	ProxyHostname *string `json:"proxy_hostname,omitempty"`
   
    // User to authenticate with tunnel endpoint. Only valid for tunnel_endpoint partner types.
    
 	ProxyUser *string `json:"proxy_user,omitempty"`
   
    // IP address or hostname of partner data interface.
    
 	ReplHostname *string `json:"repl_hostname,omitempty"`
   
    // Port number of partner data interface.
    
  	DataPort  *int64 `json:"data_port,omitempty"`
   
    // Whether the partner is available, and responding to pings.
    
 	IsAlive *bool `json:"is_alive,omitempty"`
   
    // Replication partner group uid.
    
   	PartnerGroupUID *int64 `json:"partner_group_uid,omitempty"`
   
    // Most recent error while attempting to ping the partner.
    
 	LastKeepaliveError *string `json:"last_keepalive_error,omitempty"`
   
    // Indicates whether all volumes and volume collections have been synced to the partner.
    
   	CfgSyncStatus *NsPartnerCfgSyncStatus `json:"cfg_sync_status,omitempty"`
   
    // Most recent error seen while attempting to sync objects to the partner.
    
 	LastSyncError *string `json:"last_sync_error,omitempty"`
   
    // Serial number of group leader array of the partner.
    
 	ArraySerial *string `json:"array_serial,omitempty"`
   
    // Replication version of the partner.
    
  	Version  *int64 `json:"version,omitempty"`
   
    // The pool ID where volumes replicated from this partner will be created. Replica volumes created as clones ignore this parameter and are always created in the same pool as their parent volume.
    
 	PoolID *string `json:"pool_id,omitempty"`
   
    // The pool name where volumes replicated from this partner will be created.
    
 	PoolName *string `json:"pool_name,omitempty"`
   
    // The Folder ID within the pool where volumes replicated from this partner will be created. This is not supported for pool partners.
    
 	FolderID *string `json:"folder_id,omitempty"`
   
    // The Folder name within the pool where volumes replicated from this partner will be created.
    
 	FolderName *string `json:"folder_name,omitempty"`
   
    // Indicates whether to match the upstream volume's folder on the downstream.
    
 	MatchFolder *bool `json:"match_folder,omitempty"`
   
    // Indicates whether replication traffic from/to this partner has been halted.
    
 	Paused *bool `json:"paused,omitempty"`
   
    // Indicates whether this partner actively mangles object names to avoid name conflicts during replication.
    
 	UniqueName *bool `json:"unique_name,omitempty"`
   
    // Label of the subnet used to replicate to this partner.
    
 	SubnetLabel *string `json:"subnet_label,omitempty"`
   
    // Type of the subnet used to replicate to this partner.
    
   	SubnetType *NsSubnetType `json:"subnet_type,omitempty"`
   
    // Throttles used while replicating from/to this partner.
    
   	Throttles []*NsThrottle `json:"throttles,omitempty"`
   
    // Current bandwidth throttle for this partner, expressed either as megabits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no throttle. This attribute is superseded by throttled_bandwidth_current.
    
  	ThrottledBandwIDth  *int64 `json:"throttled_bandwidth,omitempty"`
   
    // Current bandwidth throttle for this partner, expressed either as megabits per second or as -1 to indicate that there is no throttle.
    
  	ThrottledBandwIDthCurrent  *int64 `json:"throttled_bandwidth_current,omitempty"`
   
    // Current bandwidth throttle for this partner, expressed either as kilobits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no throttle. This attribute is superseded by throttled_bandwidth_current_kbps.
    
  	ThrottledBandwIDthKbps  *int64 `json:"throttled_bandwidth_kbps,omitempty"`
   
    // Current bandwidth throttle for this partner, expressed either as kilobits per second or as -1 to indicate that there is no throttle.
    
  	ThrottledBandwIDthCurrentKbps  *int64 `json:"throttled_bandwidth_current_kbps,omitempty"`
   
    // Subnet used to replicate to this partner.
    
 	SubnetNetwork *string `json:"subnet_network,omitempty"`
   
    // Subnet mask used to replicate to this partner.
    
 	SubnetNetmask *string `json:"subnet_netmask,omitempty"`
   
    // List of volume collections that are replicating from/to this partner.
    
   	VolumeCollectionList []*NsVolumeCollectionSummary `json:"volume_collection_list,omitempty"`
   
    // Count of volume collections that are replicating from/to this partner.
    
   	VolumeCollectionListCount *int64 `json:"volume_collection_list_count,omitempty"`
   
    // List of volumes that are replicating from/to this partner.
    
   	VolumeList []*NsVolumeSummary `json:"volume_list,omitempty"`
   
    // Count of volumes that are replicating from/to this partner.
    
   	VolumeListCount *int64 `json:"volume_list_count,omitempty"`
   
    // Direction of replication configured with this partner.
    
   	ReplicationDirection *NsReplDirection `json:"replication_direction,omitempty"`
}
