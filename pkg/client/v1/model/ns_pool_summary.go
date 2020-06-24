// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsPoolSummary - Object containing pool ID and name.
// Export NsPoolSummaryFields for advance operations like search filter etc.
var NsPoolSummaryFields *NsPoolSummary

func init(){
	IDfield:= "id"
	Namefield:= "name"
		
	NsPoolSummaryFields= &NsPoolSummary{
	ID: &IDfield,
	Name: &Namefield,
		
	}
}

type NsPoolSummary struct {
	// ID - ID of pool.
 	ID *string `json:"id,omitempty"`
	// Name - Name of pool.
 	Name *string `json:"name,omitempty"`
}
