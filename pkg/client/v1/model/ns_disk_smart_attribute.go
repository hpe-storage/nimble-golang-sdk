// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsDiskSmartAttribute - One Smart attribute of a disk.
// Export NsDiskSmartAttributeFields for advance operations like search filter etc.
var NsDiskSmartAttributeFields *NsDiskSmartAttribute

func init(){
	Namefield:= "name"
		
	NsDiskSmartAttributeFields= &NsDiskSmartAttribute{
		Name:                   &Namefield,
	}
}

type NsDiskSmartAttribute struct {
	// Name - Name of the Smart attribute.
 	Name *string `json:"name,omitempty"`
	// SmartId - Smart attribute ID.
   	SmartId *int64 `json:"smart_id,omitempty"`
	// CurValue - Current value.
   	CurValue *int64 `json:"cur_value,omitempty"`
	// WorstValue - Worst value.
   	WorstValue *int64 `json:"worst_value,omitempty"`
	// Threshold - Smart threshold.
   	Threshold *int64 `json:"threshold,omitempty"`
	// Flags - Smart flags.
   	Flags *int64 `json:"flags,omitempty"`
	// RawValue - Raw value.
   	RawValue *int64 `json:"raw_value,omitempty"`
	// LastUpdatedEpochSecs - Last update time in epoch seconds.
   	LastUpdatedEpochSecs *int64 `json:"last_updated_epoch_secs,omitempty"`
}
