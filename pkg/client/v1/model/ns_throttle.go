/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsThrottle


// NsThrottle :
type NsThrottle struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // ThrPartnerID
   ThrPartnerID string `json:"thr_partner_id,omitempty"`
   // ThrDayMask
   ThrDayMask float64 `json:"thr_day_mask,omitempty"`
   // Days
   Days string `json:"days,omitempty"`
   // ThrAtTime
   ThrAtTime float64 `json:"thr_at_time,omitempty"`
   // ThrUntilTime
   ThrUntilTime float64 `json:"thr_until_time,omitempty"`
   // ThrBandwIDth
   ThrBandwIDth float64 `json:"thr_bandwidth,omitempty"`
   // ThrBandwIDthKbps
   ThrBandwIDthKbps float64 `json:"thr_bandwidth_kbps,omitempty"`
   // ThrBandwIDthLimitKbps
   ThrBandwIDthLimitKbps  int32 `json:"thr_bandwidth_limit_kbps,omitempty"`
}
