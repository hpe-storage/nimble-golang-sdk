/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Token


// Token :
type Token struct {
   // ID
   ID string `json:"id,omitempty"`
   // SessionToken
   SessionToken string `json:"session_token,omitempty"`
   // Username
   Username string `json:"username,omitempty"`
   // Password
   Password string `json:"password,omitempty"`
   // AppName
   AppName string `json:"app_name,omitempty"`
   // SdkName
   SdkName string `json:"sdk_name,omitempty"`
   // SourceIp
   SourceIp string `json:"source_ip,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // ExpiryTime
   ExpiryTime float64 `json:"expiry_time,omitempty"`
}
