/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/AuditLog


// AuditLog :
type AuditLog struct {
   // ID
   ID string `json:"id,omitempty"`
   // Type
   Type float64 `json:"type,omitempty"`
   // ObjectID
   ObjectID string `json:"object_id,omitempty"`
   // ObjectName
   ObjectName string `json:"object_name,omitempty"`
   // Scope
   Scope string `json:"scope,omitempty"`
   // Time
   Time float64 `json:"time,omitempty"`
   // ErrorCode
   ErrorCode string `json:"error_code,omitempty"`
   // UserID
   UserID string `json:"user_id,omitempty"`
   // UserName
   UserName string `json:"user_name,omitempty"`
   // UserFullName
   UserFullName string `json:"user_full_name,omitempty"`
   // SourceIp
   SourceIp string `json:"source_ip,omitempty"`
   // ExtUserID
   ExtUserID string `json:"ext_user_id,omitempty"`
   // ExtUserGroupID
   ExtUserGroupID string `json:"ext_user_group_id,omitempty"`
   // ExtUserGroupName
   ExtUserGroupName string `json:"ext_user_group_name,omitempty"`
   // AppName
   AppName string `json:"app_name,omitempty"`
   // AccessType
   AccessType string `json:"access_type,omitempty"`
   // Activity
   Activity string `json:"activity,omitempty"`
}
