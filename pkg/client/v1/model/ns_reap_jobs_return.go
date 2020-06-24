/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsReapJobsReturn - Response from reaping jobs.
// Export NsReapJobsReturnFields for advance operations like search filter etc.
var NsReapJobsReturnFields *NsReapJobsReturn

func init(){
		
	NsReapJobsReturnFields= &NsReapJobsReturn{
		
	}
}

type NsReapJobsReturn struct {
   
    // Number of jobs reaped.
    
   	Reaped *int64 `json:"reaped,omitempty"`
   
    // Number of jobs remaining.
    
   	Remaining *int64 `json:"remaining,omitempty"`
}
