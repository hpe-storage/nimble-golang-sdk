// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSnapcollSummary - Select fields containing snapshot collection information.
// Export NsSnapcollSummaryFields for advance operations like search filter etc.
var NsSnapcollSummaryFields *NsSnapcollSummary

func init() {

	NsSnapcollSummaryFields = &NsSnapcollSummary{
		SnapcollId:   "snapcoll_id",
		SnapcollName: "snapcoll_name",
	}
}

type NsSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
	SnapcollId string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
	SnapcollName string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
	SnapcollCreationTime int64 `json:"snapcoll_creation_time,omitempty"`
}

// sdk internal struct
type nsSnapcollSummary struct {
	// SnapcollId - ID of snapshot collection.
	SnapcollId *string `json:"snapcoll_id,omitempty"`
	// SnapcollName - Name of snapshot collection.
	SnapcollName *string `json:"snapcoll_name,omitempty"`
	// SnapcollCreationTime - Creation time of snapshot collection.
	SnapcollCreationTime *int64 `json:"snapcoll_creation_time,omitempty"`
}

// EncodeNsSnapcollSummary - Transform NsSnapcollSummary to nsSnapcollSummary type
func EncodeNsSnapcollSummary(request interface{}) (*nsSnapcollSummary, error) {
	reqNsSnapcollSummary := request.(*NsSnapcollSummary)
	byte, err := json.Marshal(reqNsSnapcollSummary)
	objPtr := &nsSnapcollSummary{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSnapcollSummary Transform nsSnapcollSummary to NsSnapcollSummary type
func DecodeNsSnapcollSummary(request interface{}) (*NsSnapcollSummary, error) {
	reqNsSnapcollSummary := request.(*nsSnapcollSummary)
	byte, err := json.Marshal(reqNsSnapcollSummary)
	obj := &NsSnapcollSummary{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
