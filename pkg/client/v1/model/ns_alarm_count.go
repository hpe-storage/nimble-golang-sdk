/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAlarmCount - List of alarm count for each category.
// Export NsAlarmCountFields for advance operations like search filter etc.
var NsAlarmCountFields *NsAlarmCount

func init(){
		
	NsAlarmCountFields= &NsAlarmCount{
		
	}
}

type NsAlarmCount struct {
   
    // Alert category.
    
   	Category *NsEventCategory `json:"category,omitempty"`
   
    // Critical alarm count of a particular category.
    
   	Critical *int64 `json:"critical,omitempty"`
   
    // Warning alarm count of a particular category.
    
   	Warning *int64 `json:"warning,omitempty"`
}
