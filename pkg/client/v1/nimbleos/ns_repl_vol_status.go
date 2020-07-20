// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsReplVolStatus - The replication status of a volume undergoing replication.
// Export NsReplVolStatusFields for advance operations like search filter etc.
var NsReplVolStatusFields *NsReplVolStatus

func init() {

	NsReplVolStatusFields = &NsReplVolStatus{
		Name:     "name",
		SnapName: "snap_name",
	}
}

type NsReplVolStatus struct {
	// Name - Name of the volume being replicated.
	Name string `json:"name,omitempty"`
	// SnapName - Name of the snapshot being replicated.
	SnapName string `json:"snap_name,omitempty"`
	// Status - Replication status of the volume.
	Status *NsReplVolPartnerStatus `json:"status,omitempty"`
	// InternalStatus - Internal replication status of the volume.
	InternalStatus *NsReplVolPartnerStatus `json:"internal_status,omitempty"`
	// ReplBytesDone - Transferred bytes.
	ReplBytesDone int64 `json:"repl_bytes_done,omitempty"`
	// ReplBytesTotal - Total number of bytes to be transferred.
	ReplBytesTotal int64 `json:"repl_bytes_total,omitempty"`
}

// sdk internal struct
type nsReplVolStatus struct {
	Name           *string                 `json:"name,omitempty"`
	SnapName       *string                 `json:"snap_name,omitempty"`
	Status         *NsReplVolPartnerStatus `json:"status,omitempty"`
	InternalStatus *NsReplVolPartnerStatus `json:"internal_status,omitempty"`
	ReplBytesDone  *int64                  `json:"repl_bytes_done,omitempty"`
	ReplBytesTotal *int64                  `json:"repl_bytes_total,omitempty"`
}

// EncodeNsReplVolStatus - Transform NsReplVolStatus to nsReplVolStatus type
func EncodeNsReplVolStatus(request interface{}) (*nsReplVolStatus, error) {
	reqNsReplVolStatus := request.(*NsReplVolStatus)
	byte, err := json.Marshal(reqNsReplVolStatus)
	if err != nil {
		return nil, err
	}
	respNsReplVolStatusPtr := &nsReplVolStatus{}
	err = json.Unmarshal(byte, respNsReplVolStatusPtr)
	return respNsReplVolStatusPtr, err
}

// DecodeNsReplVolStatus Transform nsReplVolStatus to NsReplVolStatus type
func DecodeNsReplVolStatus(request interface{}) (*NsReplVolStatus, error) {
	reqNsReplVolStatus := request.(*nsReplVolStatus)
	byte, err := json.Marshal(reqNsReplVolStatus)
	if err != nil {
		return nil, err
	}
	respNsReplVolStatus := &NsReplVolStatus{}
	err = json.Unmarshal(byte, respNsReplVolStatus)
	return respNsReplVolStatus, err
}
