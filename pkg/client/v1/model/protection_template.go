/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ProtectionTemplate


// ProtectionTemplate :
type ProtectionTemplate struct {
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
   // AppServer
   AppServer string `json:"app_server,omitempty"`
   // AppClusterName
   AppClusterName string `json:"app_cluster_name,omitempty"`
   // AppServiceName
   AppServiceName string `json:"app_service_name,omitempty"`
   // VcenterHostname
   VcenterHostname string `json:"vcenter_hostname,omitempty"`
   // VcenterUsername
   VcenterUsername string `json:"vcenter_username,omitempty"`
   // VcenterPassword
   VcenterPassword string `json:"vcenter_password,omitempty"`
   // AgentHostname
   AgentHostname string `json:"agent_hostname,omitempty"`
   // AgentUsername
   AgentUsername string `json:"agent_username,omitempty"`
   // AgentPassword
   AgentPassword string `json:"agent_password,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
}
