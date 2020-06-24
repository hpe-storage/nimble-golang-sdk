// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsSensorCumulativeData - Stat sensor cumulative data.
// Export NsSensorCumulativeDataFields for advance operations like search filter etc.
var NsSensorCumulativeDataFields *NsSensorCumulativeData

func init(){
	Namefield:= "name"
		
	NsSensorCumulativeDataFields= &NsSensorCumulativeData{
	Name: &Namefield,
		
	}
}

type NsSensorCumulativeData struct {
	// Name - Name of sensor.
 	Name *string `json:"name,omitempty"`
	// Index - Index of sensor in stat for versioning.
   	Index *int64 `json:"index,omitempty"`
	// Msec - Creation time of stat sensor in millisecs.
   	Msec *int64 `json:"msec,omitempty"`
	// PrevUnavail - Previous stat value is not valid.
 	PrevUnavail *bool `json:"prev_unavail,omitempty"`
	// CurrUnavail - Current stat value is not valid.
 	CurrUnavail *bool `json:"curr_unavail,omitempty"`
	// Curr - Current stat value.
   	Curr *int64 `json:"curr,omitempty"`
	// Prev - Previous stat value from last sample period.
   	Prev *int64 `json:"prev,omitempty"`
}
