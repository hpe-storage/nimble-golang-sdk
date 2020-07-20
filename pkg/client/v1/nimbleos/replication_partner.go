// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// ReplicationPartner - Manage replication partner. Replication partners let one storage array talk to another for replication purposes. The two arrays must be able to communicate over a network, and use ports 4213 and 4214. Replication partners have the same name as the remote group. Replication partners can be reciprocal, upstream (the source of replicas), or downstream (the receiver of replicas) partners.
// Export ReplicationPartnerFields for advance operations like search filter etc.
var ReplicationPartnerFields *ReplicationPartner

func init() {

	ReplicationPartnerFields = &ReplicationPartner{
		ID:                 "id",
		Name:               "name",
		FullName:           "full_name",
		SearchName:         "search_name",
		Description:        "description",
		Alias:              "alias",
		Secret:             "secret",
		Hostname:           "hostname",
		ProxyHostname:      "proxy_hostname",
		ProxyUser:          "proxy_user",
		ReplHostname:       "repl_hostname",
		LastKeepaliveError: "last_keepalive_error",
		LastSyncError:      "last_sync_error",
		ArraySerial:        "array_serial",
		PoolId:             "pool_id",
		PoolName:           "pool_name",
		FolderId:           "folder_id",
		FolderName:         "folder_name",
		SubnetLabel:        "subnet_label",
		SubnetNetwork:      "subnet_network",
		SubnetNetmask:      "subnet_netmask",
	}
}

type ReplicationPartner struct {
	// ID - Identifier for a replication partner.
	ID string `json:"id,omitempty"`
	// Name - Name of replication partner.
	Name string `json:"name,omitempty"`
	// FullName - Fully qualified name of replication partner.
	FullName string `json:"full_name,omitempty"`
	// SearchName - Name of replication partner used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description - Description of replication partner.
	Description string `json:"description,omitempty"`
	// PartnerType - Replication partner type. Possible values are group or pool.
	PartnerType *NsPartnerType `json:"partner_type,omitempty"`
	// Alias - Name this group uses to identify itself to this partner.
	Alias string `json:"alias,omitempty"`
	// Secret - Replication partner shared secret, used for mutual authentication of the partners.
	Secret string `json:"secret,omitempty"`
	// CreationTime - Time when this replication partner was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this replication partner was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// ControlPort - Port number of partner control interface.
	ControlPort int64 `json:"control_port,omitempty"`
	// Hostname - IP address or hostname of partner interface.  This must be the partner's Group Management IP address.
	Hostname string `json:"hostname,omitempty"`
	// PortRangeStart - For tunnel_endpoint partner types, first port available on the ssh proxy available for reverse forwarding. It must be guaranteed that the proxy has the next N ports reserved for this partner, where N is the count of DSDs in this group. This attribute is only valid for tunnel_endpoint partner type.
	PortRangeStart int64 `json:"port_range_start,omitempty"`
	// ProxyHostname - IP address of tunnel endpoint. Only valid for tunnel_endpoint partner types.
	ProxyHostname string `json:"proxy_hostname,omitempty"`
	// ProxyUser - User to authenticate with tunnel endpoint. Only valid for tunnel_endpoint partner types.
	ProxyUser string `json:"proxy_user,omitempty"`
	// ReplHostname - IP address or hostname of partner data interface.
	ReplHostname string `json:"repl_hostname,omitempty"`
	// DataPort - Port number of partner data interface.
	DataPort int64 `json:"data_port,omitempty"`
	// IsAlive - Whether the partner is available, and responding to pings.
	IsAlive *bool `json:"is_alive,omitempty"`
	// PartnerGroupUid - Replication partner group uid.
	PartnerGroupUid int64 `json:"partner_group_uid,omitempty"`
	// LastKeepaliveError - Most recent error while attempting to ping the partner.
	LastKeepaliveError string `json:"last_keepalive_error,omitempty"`
	// CfgSyncStatus - Indicates whether all volumes and volume collections have been synced to the partner.
	CfgSyncStatus *NsPartnerCfgSyncStatus `json:"cfg_sync_status,omitempty"`
	// LastSyncError - Most recent error seen while attempting to sync objects to the partner.
	LastSyncError string `json:"last_sync_error,omitempty"`
	// ArraySerial - Serial number of group leader array of the partner.
	ArraySerial string `json:"array_serial,omitempty"`
	// Version - Replication version of the partner.
	Version int64 `json:"version,omitempty"`
	// PoolId - The pool ID where volumes replicated from this partner will be created. Replica volumes created as clones ignore this parameter and are always created in the same pool as their parent volume.
	PoolId string `json:"pool_id,omitempty"`
	// PoolName - The pool name where volumes replicated from this partner will be created.
	PoolName string `json:"pool_name,omitempty"`
	// FolderId - The Folder ID within the pool where volumes replicated from this partner will be created. This is not supported for pool partners.
	FolderId string `json:"folder_id,omitempty"`
	// FolderName - The Folder name within the pool where volumes replicated from this partner will be created.
	FolderName string `json:"folder_name,omitempty"`
	// MatchFolder - Indicates whether to match the upstream volume's folder on the downstream.
	MatchFolder *bool `json:"match_folder,omitempty"`
	// Paused - Indicates whether replication traffic from/to this partner has been halted.
	Paused *bool `json:"paused,omitempty"`
	// UniqueName - Indicates whether this partner actively mangles object names to avoid name conflicts during replication.
	UniqueName *bool `json:"unique_name,omitempty"`
	// SubnetLabel - Label of the subnet used to replicate to this partner.
	SubnetLabel string `json:"subnet_label,omitempty"`
	// SubnetType - Type of the subnet used to replicate to this partner.
	SubnetType *NsSubnetType `json:"subnet_type,omitempty"`
	// Throttles - Throttles used while replicating from/to this partner.
	Throttles []*NsThrottle `json:"throttles,omitempty"`
	// ThrottledBandwidth - Current bandwidth throttle for this partner, expressed either as megabits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no throttle. This attribute is superseded by throttled_bandwidth_current.
	ThrottledBandwidth int64 `json:"throttled_bandwidth,omitempty"`
	// ThrottledBandwidthCurrent - Current bandwidth throttle for this partner, expressed either as megabits per second or as -1 to indicate that there is no throttle.
	ThrottledBandwidthCurrent int64 `json:"throttled_bandwidth_current,omitempty"`
	// ThrottledBandwidthKbps - Current bandwidth throttle for this partner, expressed either as kilobits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no throttle. This attribute is superseded by throttled_bandwidth_current_kbps.
	ThrottledBandwidthKbps int64 `json:"throttled_bandwidth_kbps,omitempty"`
	// ThrottledBandwidthCurrentKbps - Current bandwidth throttle for this partner, expressed either as kilobits per second or as -1 to indicate that there is no throttle.
	ThrottledBandwidthCurrentKbps int64 `json:"throttled_bandwidth_current_kbps,omitempty"`
	// SubnetNetwork - Subnet used to replicate to this partner.
	SubnetNetwork string `json:"subnet_network,omitempty"`
	// SubnetNetmask - Subnet mask used to replicate to this partner.
	SubnetNetmask string `json:"subnet_netmask,omitempty"`
	// VolumeCollectionList - List of volume collections that are replicating from/to this partner.
	VolumeCollectionList []*NsVolumeCollectionSummary `json:"volume_collection_list,omitempty"`
	// VolumeCollectionListCount - Count of volume collections that are replicating from/to this partner.
	VolumeCollectionListCount int64 `json:"volume_collection_list_count,omitempty"`
	// VolumeList - List of volumes that are replicating from/to this partner.
	VolumeList []*NsVolumeSummary `json:"volume_list,omitempty"`
	// VolumeListCount - Count of volumes that are replicating from/to this partner.
	VolumeListCount int64 `json:"volume_list_count,omitempty"`
	// ReplicationDirection - Direction of replication configured with this partner.
	ReplicationDirection *NsReplDirection `json:"replication_direction,omitempty"`
}

// sdk internal struct
type replicationPartner struct {
	ID                            *string                      `json:"id,omitempty"`
	Name                          *string                      `json:"name,omitempty"`
	FullName                      *string                      `json:"full_name,omitempty"`
	SearchName                    *string                      `json:"search_name,omitempty"`
	Description                   *string                      `json:"description,omitempty"`
	PartnerType                   *NsPartnerType               `json:"partner_type,omitempty"`
	Alias                         *string                      `json:"alias,omitempty"`
	Secret                        *string                      `json:"secret,omitempty"`
	CreationTime                  *int64                       `json:"creation_time,omitempty"`
	LastModified                  *int64                       `json:"last_modified,omitempty"`
	ControlPort                   *int64                       `json:"control_port,omitempty"`
	Hostname                      *string                      `json:"hostname,omitempty"`
	PortRangeStart                *int64                       `json:"port_range_start,omitempty"`
	ProxyHostname                 *string                      `json:"proxy_hostname,omitempty"`
	ProxyUser                     *string                      `json:"proxy_user,omitempty"`
	ReplHostname                  *string                      `json:"repl_hostname,omitempty"`
	DataPort                      *int64                       `json:"data_port,omitempty"`
	IsAlive                       *bool                        `json:"is_alive,omitempty"`
	PartnerGroupUid               *int64                       `json:"partner_group_uid,omitempty"`
	LastKeepaliveError            *string                      `json:"last_keepalive_error,omitempty"`
	CfgSyncStatus                 *NsPartnerCfgSyncStatus      `json:"cfg_sync_status,omitempty"`
	LastSyncError                 *string                      `json:"last_sync_error,omitempty"`
	ArraySerial                   *string                      `json:"array_serial,omitempty"`
	Version                       *int64                       `json:"version,omitempty"`
	PoolId                        *string                      `json:"pool_id,omitempty"`
	PoolName                      *string                      `json:"pool_name,omitempty"`
	FolderId                      *string                      `json:"folder_id,omitempty"`
	FolderName                    *string                      `json:"folder_name,omitempty"`
	MatchFolder                   *bool                        `json:"match_folder,omitempty"`
	Paused                        *bool                        `json:"paused,omitempty"`
	UniqueName                    *bool                        `json:"unique_name,omitempty"`
	SubnetLabel                   *string                      `json:"subnet_label,omitempty"`
	SubnetType                    *NsSubnetType                `json:"subnet_type,omitempty"`
	Throttles                     []*NsThrottle                `json:"throttles,omitempty"`
	ThrottledBandwidth            *int64                       `json:"throttled_bandwidth,omitempty"`
	ThrottledBandwidthCurrent     *int64                       `json:"throttled_bandwidth_current,omitempty"`
	ThrottledBandwidthKbps        *int64                       `json:"throttled_bandwidth_kbps,omitempty"`
	ThrottledBandwidthCurrentKbps *int64                       `json:"throttled_bandwidth_current_kbps,omitempty"`
	SubnetNetwork                 *string                      `json:"subnet_network,omitempty"`
	SubnetNetmask                 *string                      `json:"subnet_netmask,omitempty"`
	VolumeCollectionList          []*NsVolumeCollectionSummary `json:"volume_collection_list,omitempty"`
	VolumeCollectionListCount     *int64                       `json:"volume_collection_list_count,omitempty"`
	VolumeList                    []*NsVolumeSummary           `json:"volume_list,omitempty"`
	VolumeListCount               *int64                       `json:"volume_list_count,omitempty"`
	ReplicationDirection          *NsReplDirection             `json:"replication_direction,omitempty"`
}

// EncodeReplicationPartner - Transform ReplicationPartner to replicationPartner type
func EncodeReplicationPartner(request interface{}) (*replicationPartner, error) {
	reqReplicationPartner := request.(*ReplicationPartner)
	byte, err := json.Marshal(reqReplicationPartner)
	if err != nil {
		return nil, err
	}
	respReplicationPartnerPtr := &replicationPartner{}
	err = json.Unmarshal(byte, respReplicationPartnerPtr)
	return respReplicationPartnerPtr, err
}

// DecodeReplicationPartner Transform replicationPartner to ReplicationPartner type
func DecodeReplicationPartner(request interface{}) (*ReplicationPartner, error) {
	reqReplicationPartner := request.(*replicationPartner)
	byte, err := json.Marshal(reqReplicationPartner)
	if err != nil {
		return nil, err
	}
	respReplicationPartner := &ReplicationPartner{}
	err = json.Unmarshal(byte, respReplicationPartner)
	return respReplicationPartner, err
}
