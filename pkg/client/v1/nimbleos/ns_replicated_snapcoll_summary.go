// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsReplicatedSnapcollSummary - Select fields containing snapshot collection information for replicated snapshot collections.
// Export NsReplicatedSnapcollSummaryFields for advance operations like search filter etc.
var NsReplicatedSnapcollSummaryFields *NsReplicatedSnapcollSummary

func init() {

	NsReplicatedSnapcollSummaryFields = &NsReplicatedSnapcollSummary{
		SnapcollId:            "snapcoll_id",
		SnapcollName:          "snapcoll_name",
		DownstreamPartnerName: "downstream_partner_name",
	}
}

type NsReplicatedSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
	SnapcollId string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
	SnapcollName string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
	SnapcollCreationTime int64 `json:"snapcoll_creation_time,omitempty"`
	// DownstreamPartnerName - Name of partner the snapshot collection is replicated to.
	DownstreamPartnerName string `json:"downstream_partner_name,omitempty"`
	// VolIds - Id of volumes that have snapshots which are part of the snapcoll.
	VolIds []*string `json:"vol_ids,omitempty"`
}

// sdk internal struct
type nsReplicatedSnapcollSummary struct {
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

// EncodeNsReplicatedSnapcollSummary - Transform NsReplicatedSnapcollSummary to nsReplicatedSnapcollSummary type
func EncodeNsReplicatedSnapcollSummary(request interface{}) (*nsReplicatedSnapcollSummary, error) {
	reqNsReplicatedSnapcollSummary := request.(*NsReplicatedSnapcollSummary)
	byte, err := json.Marshal(reqNsReplicatedSnapcollSummary)
	objPtr := &nsReplicatedSnapcollSummary{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsReplicatedSnapcollSummary Transform nsReplicatedSnapcollSummary to NsReplicatedSnapcollSummary type
func DecodeNsReplicatedSnapcollSummary(request interface{}) (*NsReplicatedSnapcollSummary, error) {
	reqNsReplicatedSnapcollSummary := request.(*nsReplicatedSnapcollSummary)
	byte, err := json.Marshal(reqNsReplicatedSnapcollSummary)
	obj := &NsReplicatedSnapcollSummary{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
