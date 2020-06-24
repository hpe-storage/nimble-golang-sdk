/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSnapSummary - Select fields containing snapshot information.
// Export NsSnapSummaryFields for advance operations like search filter etc.
var NsSnapSummaryFields *NsSnapSummary

func init(){
	SnapIDfield:= "snap_id"
	SnapNamefield:= "snap_name"
		
	NsSnapSummaryFields= &NsSnapSummary{
		SnapID: &SnapIDfield,
		SnapName: &SnapNamefield,
		
	}
}

type NsSnapSummary struct {
   
    // ID of snapshot.
    
 	SnapID *string `json:"snap_id,omitempty"`
   
    // Name of snapshot.
    
 	SnapName *string `json:"snap_name,omitempty"`
   
    // Creation time of snapshot.
    
   	SnapCreationTime *int64 `json:"snap_creation_time,omitempty"`
}
