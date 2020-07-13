// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSnapSummary - Select fields containing snapshot information.
// Export NsSnapSummaryFields for advance operations like search filter etc.
var NsSnapSummaryFields *NsSnapSummary

func init() {

	NsSnapSummaryFields = &NsSnapSummary{
		SnapId:   "snap_id",
		SnapName: "snap_name",
	}
}

type NsSnapSummary struct {
	// SnapId - ID of snapshot.
	SnapId string `json:"snap_id,omitempty"`
	// SnapName - Name of snapshot.
	SnapName string `json:"snap_name,omitempty"`
	// SnapCreationTime - Creation time of snapshot.
	SnapCreationTime int64 `json:"snap_creation_time,omitempty"`
}

// sdk internal struct
type nsSnapSummary struct {
	// SnapId - ID of snapshot.
	SnapId *string `json:"snap_id,omitempty"`
	// SnapName - Name of snapshot.
	SnapName *string `json:"snap_name,omitempty"`
	// SnapCreationTime - Creation time of snapshot.
	SnapCreationTime *int64 `json:"snap_creation_time,omitempty"`
}

// EncodeNsSnapSummary - Transform NsSnapSummary to nsSnapSummary type
func EncodeNsSnapSummary(request interface{}) (*nsSnapSummary, error) {
	reqNsSnapSummary := request.(*NsSnapSummary)
	byte, err := json.Marshal(reqNsSnapSummary)
	objPtr := &nsSnapSummary{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSnapSummary Transform nsSnapSummary to NsSnapSummary type
func DecodeNsSnapSummary(request interface{}) (*NsSnapSummary, error) {
	reqNsSnapSummary := request.(*nsSnapSummary)
	byte, err := json.Marshal(reqNsSnapSummary)
	obj := &NsSnapSummary{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
