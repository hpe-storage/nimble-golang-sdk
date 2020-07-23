// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsThrottle - A single throttle for the partner.
// Export NsThrottleFields for advance operations like search filter etc.
var NsThrottleFields *NsThrottle

func init() {
	IDfield := "id"
	Namefield := "name"
	Descriptionfield := "description"
	ThrPartnerIdfield := "thr_partner_id"
	Daysfield := "days"

	NsThrottleFields = &NsThrottle{
		ID:           &IDfield,
		Name:         &Namefield,
		Description:  &Descriptionfield,
		ThrPartnerId: &ThrPartnerIdfield,
		Days:         &Daysfield,
	}
}

type NsThrottle struct {
	// ID - Id of the throttle.
	ID *string `json:"id,omitempty"`
	// Name - Name of the throttle.
	Name *string `json:"name,omitempty"`
	// Description - Description of the throttle.
	Description *string `json:"description,omitempty"`
	// CreationTime - Creation time of the throttle.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Last modification time of the throttle.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ThrPartnerId - ID of the partner object.
	ThrPartnerId *string `json:"thr_partner_id,omitempty"`
	// ThrDayMask - Mask for days that the throttle operates.
	ThrDayMask *int64 `json:"thr_day_mask,omitempty"`
	// Days - List of days that the throttle operates.
	Days *string `json:"days,omitempty"`
	// ThrAtTime - Start time set for the throttle.
	ThrAtTime *int64 `json:"thr_at_time,omitempty"`
	// ThrUntilTime - End time set for the throttle.
	ThrUntilTime *int64 `json:"thr_until_time,omitempty"`
	// ThrBandwidth - Bandwidth set for the throttle in megabits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no limit. This atttibute is superseded by thr_bandwidth_limit_kbps.
	ThrBandwidth *int64 `json:"thr_bandwidth,omitempty"`
	// ThrBandwidthKbps - Bandwidth set for the throttle in kilobits per second or as the largest possible 64-bit signed integer (9223372036854775807) to indicate that there is no limit. This atttibute is superseded by thr_bandwidth_limit_kbps.
	ThrBandwidthKbps *int64 `json:"thr_bandwidth_kbps,omitempty"`
	// ThrBandwidthLimitKbps - Bandwidth set for the throttle in kilobits per second or -1 to indicate that there is no limit.
	ThrBandwidthLimitKbps *int64 `json:"thr_bandwidth_limit_kbps,omitempty"`
}
