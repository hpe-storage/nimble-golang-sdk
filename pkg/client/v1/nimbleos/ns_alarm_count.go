// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsAlarmCount - List of alarm count for each category.
// Export NsAlarmCountFields for advance operations like search filter etc.
var NsAlarmCountFields *NsAlarmCount

func init(){

 NsAlarmCountFields= &NsAlarmCount{
 }
}


type NsAlarmCount struct {
 // Category - Alert category.
    Category *NsEventCategory `json:"category,omitempty"`
 // Critical - Critical alarm count of a particular category.
    Critical *int64 `json:"critical,omitempty"`
 // Warning - Warning alarm count of a particular category.
    Warning *int64 `json:"warning,omitempty"`
}
