/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSnapcollSummary - Select fields containing snapshot collection information.
// Export NsSnapcollSummaryFields for advance operations like search filter etc.
var NsSnapcollSummaryFields *NsSnapcollSummary

func init(){
	SnapcollIDfield:= "snapcoll_id"
	SnapcollNamefield:= "snapcoll_name"
		
	NsSnapcollSummaryFields= &NsSnapcollSummary{
		SnapcollID: &SnapcollIDfield,
		SnapcollName: &SnapcollNamefield,
		
	}
}

type NsSnapcollSummary struct {
   
    // ID of snapshot collection.
    
 	SnapcollID *string `json:"snapcoll_id,omitempty"`
   
    // Name of snapshot collection.
    
 	SnapcollName *string `json:"snapcoll_name,omitempty"`
   
    // Creation time of snapshot collection.
    
   	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
}
