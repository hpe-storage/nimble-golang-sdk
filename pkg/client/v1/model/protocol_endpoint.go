/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ProtocolEndpoint


// ProtocolEndpoint :
type ProtocolEndpoint struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // PoolName
   PoolName string `json:"pool_name,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
   // SerialNumber
   SerialNumber string `json:"serial_number,omitempty"`
   // TargetName
   TargetName string `json:"target_name,omitempty"`
   // GroupSpecificIDs
   GroupSpecificIDs bool `json:"group_specific_ids,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // NumConnections
   NumConnections float64 `json:"num_connections,omitempty"`
   // NumIscsiConnections
   NumIscsiConnections float64 `json:"num_iscsi_connections,omitempty"`
   // NumFcConnections
   NumFcConnections float64 `json:"num_fc_connections,omitempty"`
}
