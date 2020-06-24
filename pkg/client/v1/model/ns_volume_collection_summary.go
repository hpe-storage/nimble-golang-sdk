/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumeCollectionSummary - Select fields of volume collection info.
// Export NsVolumeCollectionSummaryFields for advance operations like search filter etc.
var NsVolumeCollectionSummaryFields *NsVolumeCollectionSummary

func init(){
	IDfield:= "id"
	Namefield:= "name"
		
	NsVolumeCollectionSummaryFields= &NsVolumeCollectionSummary{
		ID: &IDfield,
		Name: &Namefield,
		
	}
}

type NsVolumeCollectionSummary struct {
   
    // Identifier of volume collection.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of volume collection.
    
 	Name *string `json:"name,omitempty"`
}
