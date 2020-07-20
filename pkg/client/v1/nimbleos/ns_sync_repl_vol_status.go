// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSyncReplVolStatus - The sync replication status of a volume in volume collection.
// Export NsSyncReplVolStatusFields for advance operations like search filter etc.
var NsSyncReplVolStatusFields *NsSyncReplVolStatus

func init() {

	NsSyncReplVolStatusFields = &NsSyncReplVolStatus{}
}

type NsSyncReplVolStatus struct {
	// ResyncActive - Sync replication active status.
	ResyncActive *bool `json:"resync_active,omitempty"`
	// ResyncBytesDone - Transferred bytes.
	ResyncBytesDone int64 `json:"resync_bytes_done,omitempty"`
	// ResyncBytesTotal - Total number of bytes to be transferred.
	ResyncBytesTotal int64 `json:"resync_bytes_total,omitempty"`
}

// sdk internal struct
type nsSyncReplVolStatus struct {
	ResyncActive     *bool  `json:"resync_active,omitempty"`
	ResyncBytesDone  *int64 `json:"resync_bytes_done,omitempty"`
	ResyncBytesTotal *int64 `json:"resync_bytes_total,omitempty"`
}

// EncodeNsSyncReplVolStatus - Transform NsSyncReplVolStatus to nsSyncReplVolStatus type
func EncodeNsSyncReplVolStatus(request interface{}) (*nsSyncReplVolStatus, error) {
	reqNsSyncReplVolStatus := request.(*NsSyncReplVolStatus)
	byte, err := json.Marshal(reqNsSyncReplVolStatus)
	if err != nil {
		return nil, err
	}
	respNsSyncReplVolStatusPtr := &nsSyncReplVolStatus{}
	err = json.Unmarshal(byte, respNsSyncReplVolStatusPtr)
	return respNsSyncReplVolStatusPtr, err
}

// DecodeNsSyncReplVolStatus Transform nsSyncReplVolStatus to NsSyncReplVolStatus type
func DecodeNsSyncReplVolStatus(request interface{}) (*NsSyncReplVolStatus, error) {
	reqNsSyncReplVolStatus := request.(*nsSyncReplVolStatus)
	byte, err := json.Marshal(reqNsSyncReplVolStatus)
	if err != nil {
		return nil, err
	}
	respNsSyncReplVolStatus := &NsSyncReplVolStatus{}
	err = json.Unmarshal(byte, respNsSyncReplVolStatus)
	return respNsSyncReplVolStatus, err
}
