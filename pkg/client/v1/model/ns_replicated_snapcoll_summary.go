/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsReplicatedSnapcollSummary - Select fields containing snapshot collection information for replicated snapshot collections.
// Export NsReplicatedSnapcollSummaryFields for advance operations like search filter etc.
var NsReplicatedSnapcollSummaryFields *NsReplicatedSnapcollSummary

func init(){
	SnapcollIDfield:= "snapcoll_id"
	SnapcollNamefield:= "snapcoll_name"
	DownstreamPartnerNamefield:= "downstream_partner_name"
		
	NsReplicatedSnapcollSummaryFields= &NsReplicatedSnapcollSummary{
		SnapcollID: &SnapcollIDfield,
		SnapcollName: &SnapcollNamefield,
		DownstreamPartnerName: &DownstreamPartnerNamefield,
		
	}
}

type NsReplicatedSnapcollSummary struct {
   
    // ID of snapshot collection.
    
 	SnapcollID *string `json:"snapcoll_id,omitempty"`
   
    // Name of snapshot collection.
    
 	SnapcollName *string `json:"snapcoll_name,omitempty"`
   
    // Creation time of snapshot collection.
    
   	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
   
    // Name of partner the snapshot collection is replicated to.
    
 	DownstreamPartnerName *string `json:"downstream_partner_name,omitempty"`
   
    // Id of volumes that have snapshots which are part of the snapcoll.
    
	VolIDs []*string `json:"vol_ids,omitempty"`
}
