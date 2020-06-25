// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsSnapcollSummary - Select fields containing snapshot collection information.
// Export NsSnapcollSummaryFields for advance operations like search filter etc.
var NsSnapcollSummaryFields *NsSnapcollSummary

func init(){
	SnapcollIdfield:= "snapcoll_id"
	SnapcollNamefield:= "snapcoll_name"
		
	NsSnapcollSummaryFields= &NsSnapcollSummary{
		SnapcollId:            &SnapcollIdfield,
		SnapcollName:          &SnapcollNamefield,
	}
}

type NsSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
 	SnapcollId *string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
 	SnapcollName *string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
   	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
}
