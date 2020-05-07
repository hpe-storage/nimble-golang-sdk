/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ReplicationPartner


// ReplicationPartner :
type ReplicationPartner struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // Alias
   Alias string `json:"alias,omitempty"`
   // Secret
   Secret string `json:"secret,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // ControlPort
   ControlPort  int32 `json:"control_port,omitempty"`
   // Hostname
   Hostname string `json:"hostname,omitempty"`
   // PortRangeStart
   PortRangeStart float64 `json:"port_range_start,omitempty"`
   // ProxyHostname
   ProxyHostname string `json:"proxy_hostname,omitempty"`
   // ProxyUser
   ProxyUser string `json:"proxy_user,omitempty"`
   // ReplHostname
   ReplHostname string `json:"repl_hostname,omitempty"`
   // DataPort
   DataPort  int32 `json:"data_port,omitempty"`
   // IsAlive
   IsAlive bool `json:"is_alive,omitempty"`
   // PartnerGroupUID
   PartnerGroupUID float64 `json:"partner_group_uid,omitempty"`
   // LastKeepaliveError
   LastKeepaliveError string `json:"last_keepalive_error,omitempty"`
   // LastSyncError
   LastSyncError string `json:"last_sync_error,omitempty"`
   // ArraySerial
   ArraySerial string `json:"array_serial,omitempty"`
   // Version
   Version  int32 `json:"version,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
   // FolderID
   FolderID string `json:"folder_id,omitempty"`
   // FolderName
   FolderName string `json:"folder_name,omitempty"`
   // MatchFolder
   MatchFolder bool `json:"match_folder,omitempty"`
   // Paused
   Paused bool `json:"paused,omitempty"`
   // UniqueName
   UniqueName bool `json:"unique_name,omitempty"`
   // SubnetLabel
   SubnetLabel string `json:"subnet_label,omitempty"`
   // ThrottledBandwIDth
   ThrottledBandwIDth  int32 `json:"throttled_bandwidth,omitempty"`
   // ThrottledBandwIDthCurrent
   ThrottledBandwIDthCurrent  int32 `json:"throttled_bandwidth_current,omitempty"`
   // ThrottledBandwIDthKbps
   ThrottledBandwIDthKbps  int32 `json:"throttled_bandwidth_kbps,omitempty"`
   // ThrottledBandwIDthCurrentKbps
   ThrottledBandwIDthCurrentKbps  int32 `json:"throttled_bandwidth_current_kbps,omitempty"`
   // SubnetNetwork
   SubnetNetwork string `json:"subnet_network,omitempty"`
   // SubnetNetmask
   SubnetNetmask string `json:"subnet_netmask,omitempty"`
   // VolumeCollectionListCount
   VolumeCollectionListCount float64 `json:"volume_collection_list_count,omitempty"`
   // VolumeListCount
   VolumeListCount float64 `json:"volume_list_count,omitempty"`
}
