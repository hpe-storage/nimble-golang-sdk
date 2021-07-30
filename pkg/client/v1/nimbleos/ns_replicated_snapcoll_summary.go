// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReplicatedSnapcollSummary - Select fields containing snapshot collection information for replicated snapshot collections.
// Export NsReplicatedSnapcollSummaryFields for advance operations like search filter etc.
var NsReplicatedSnapcollSummaryFields *NsReplicatedSnapcollSummaryStringFields

func init() {
	SnapcollIdfield := "snapcoll_id"
	SnapcollNamefield := "snapcoll_name"
	SnapcollCreationTimefield := "snapcoll_creation_time"
	DownstreamPartnerNamefield := "downstream_partner_name"
	VolIdsfield := "vol_ids"

	NsReplicatedSnapcollSummaryFields = &NsReplicatedSnapcollSummaryStringFields{
		SnapcollId:            &SnapcollIdfield,
		SnapcollName:          &SnapcollNamefield,
		SnapcollCreationTime:  &SnapcollCreationTimefield,
		DownstreamPartnerName: &DownstreamPartnerNamefield,
		VolIds:                &VolIdsfield,
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

// Struct for NsReplicatedSnapcollSummaryFields
type NsReplicatedSnapcollSummaryStringFields struct {
	SnapcollId            *string
	SnapcollName          *string
	SnapcollCreationTime  *string
	DownstreamPartnerName *string
	VolIds                *string
}
