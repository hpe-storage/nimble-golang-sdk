// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsReplicatedSnapcollSummaryFields provides field names to use in filter parameters, for example.
var NsReplicatedSnapcollSummaryFields *NsReplicatedSnapcollSummaryFieldHandles

func init() {
	NsReplicatedSnapcollSummaryFields = &NsReplicatedSnapcollSummaryFieldHandles{
		SnapcollId:            "snapcoll_id",
		SnapcollName:          "snapcoll_name",
		SnapcollCreationTime:  "snapcoll_creation_time",
		DownstreamPartnerName: "downstream_partner_name",
		VolIds:                "vol_ids",
	}
}

// NsReplicatedSnapcollSummary - Select fields containing snapshot collection information for replicated snapshot collections.
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

// NsReplicatedSnapcollSummaryFieldHandles provides a string representation for each AccessControlRecord field.
type NsReplicatedSnapcollSummaryFieldHandles struct {
	SnapcollId            string
	SnapcollName          string
	SnapcollCreationTime  string
	DownstreamPartnerName string
	VolIds                string
}
