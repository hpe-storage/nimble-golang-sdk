/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsThrottle - A single throttle for the partner.
// Export NsThrottleFields for advance operations like search filter etc.
var NsThrottleFields *NsThrottle

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Descriptionfield:= "description"
	ThrPartnerIDfield:= "thr_partner_id"
	Daysfield:= "days"
		
	NsThrottleFields= &NsThrottle{
		ID: &IDfield,
		Name: &Namefield,
		Description: &Descriptionfield,
		ThrPartnerID: &ThrPartnerIDfield,
		Days: &Daysfield,
		
	}
}

type NsThrottle struct {
   
    // Id of the throttle.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the throttle.
    
 	Name *string `json:"name,omitempty"`
   
    // Description of the throttle.
    
 	Description *string `json:"description,omitempty"`
   
    // Creation time of the throttle.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Last modification time of the throttle.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // ID of the partner object.
    
 	ThrPartnerID *string `json:"thr_partner_id,omitempty"`
   
    // Mask for days that the throttle operates.
    
   	ThrDayMask *int64 `json:"thr_day_mask,omitempty"`
   
    // List of days that the throttle operates.
    
 	Days *string `json:"days,omitempty"`
   
    // Start time set for the throttle.
    
   	ThrAtTime *int64 `json:"thr_at_time,omitempty"`
   
    // End time set for the throttle.
    
   	ThrUntilTime *int64 `json:"thr_until_time,omitempty"`
   
    // Bandwidth set for the throttle in megabits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no limit. This atttibute is superseded by thr_bandwidth_limit_kbps.
    
   	ThrBandwIDth *int64 `json:"thr_bandwidth,omitempty"`
   
    // Bandwidth set for the throttle in kilobits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no limit. This atttibute is superseded by thr_bandwidth_limit_kbps.
    
   	ThrBandwIDthKbps *int64 `json:"thr_bandwidth_kbps,omitempty"`
   
    // Bandwidth set for the throttle in kilobits per second or -1 to indicate that there is no limit.
    
  	ThrBandwIDthLimitKbps  *int64 `json:"thr_bandwidth_limit_kbps,omitempty"`
}
