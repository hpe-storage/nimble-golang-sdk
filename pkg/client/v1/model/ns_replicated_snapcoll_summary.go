// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsReplicatedSnapcollSummary - Select fields containing snapshot collection information for replicated snapshot collections.
// Export NsReplicatedSnapcollSummaryFields for advance operations like search filter etc.
var NsReplicatedSnapcollSummaryFields *NsReplicatedSnapcollSummary

func init(){
	SnapcollIdfield:= "snapcoll_id"
	SnapcollNamefield:= "snapcoll_name"
	DownstreamPartnerNamefield:= "downstream_partner_name"
		
	NsReplicatedSnapcollSummaryFields= &NsReplicatedSnapcollSummary{
	SnapcollId: &SnapcollIdfield,
	SnapcollName: &SnapcollNamefield,
	DownstreamPartnerName: &DownstreamPartnerNamefield,
		
	}
}

type NsReplicatedSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
 	SnapcollId *string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
 	SnapcollName *string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
   	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
	// DownstreamPartnerName - Name of partner the snapshot collection is replicated to.
 	DownstreamPartnerName *string `json:"downstream_partner_name,omitempty"`
	// VolIds - Id of volumes that have snapshots which are part of the snapcoll.
	VolIds []*string `json:"vol_ids,omitempty"`
}
