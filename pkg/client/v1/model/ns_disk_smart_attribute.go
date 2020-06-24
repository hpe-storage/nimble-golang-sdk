/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsDiskSmartAttribute - One Smart attribute of a disk.
// Export NsDiskSmartAttributeFields for advance operations like search filter etc.
var NsDiskSmartAttributeFields *NsDiskSmartAttribute

func init(){
	Namefield:= "name"
		
	NsDiskSmartAttributeFields= &NsDiskSmartAttribute{
		Name: &Namefield,
		
	}
}

type NsDiskSmartAttribute struct {
   
    // Name of the Smart attribute.
    
 	Name *string `json:"name,omitempty"`
   
    // Smart attribute ID.
    
   	SmartID *int64 `json:"smart_id,omitempty"`
   
    // Current value.
    
   	CurValue *int64 `json:"cur_value,omitempty"`
   
    // Worst value.
    
   	WorstValue *int64 `json:"worst_value,omitempty"`
   
    // Smart threshold.
    
   	Threshold *int64 `json:"threshold,omitempty"`
   
    // Smart flags.
    
   	Flags *int64 `json:"flags,omitempty"`
   
    // Raw value.
    
   	RawValue *int64 `json:"raw_value,omitempty"`
   
    // Last update time in epoch seconds.
    
   	LastUpdatedEpochSecs *int64 `json:"last_updated_epoch_secs,omitempty"`
}
