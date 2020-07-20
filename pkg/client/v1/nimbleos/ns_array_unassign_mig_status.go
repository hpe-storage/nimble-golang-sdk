// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArrayUnassignMigStatus - Data migration status for array being unassigned from its pool.
// Export NsArrayUnassignMigStatusFields for advance operations like search filter etc.
var NsArrayUnassignMigStatusFields *NsArrayUnassignMigStatus

func init() {

	NsArrayUnassignMigStatusFields = &NsArrayUnassignMigStatus{
		ID:   "id",
		Name: "name",
	}
}

type NsArrayUnassignMigStatus struct {
	// ID - Unique identifier of the array being unassigned.
	ID string `json:"id,omitempty"`
	// Name - Name of the array being unassigned.
	Name string `json:"name,omitempty"`
	// DestinationArrays - List of arrays to which data is being migrated.
	DestinationArrays []*NsArraySummary `json:"destination_arrays,omitempty"`
	// BytesMigrated - Number of bytes already migrated to the destination arrays.
	BytesMigrated int64 `json:"bytes_migrated,omitempty"`
	// BytesRemaining - Number of bytes yet to be migrated to the destination arrays.
	BytesRemaining int64 `json:"bytes_remaining,omitempty"`
	// StartTime - Time when array unassign operation started.
	StartTime int64 `json:"start_time,omitempty"`
	// EstimatedCompletionTime - Estimated completion time. 0 if unknown.
	EstimatedCompletionTime int64 `json:"estimated_completion_time,omitempty"`
}

// sdk internal struct
type nsArrayUnassignMigStatus struct {
	ID                      *string           `json:"id,omitempty"`
	Name                    *string           `json:"name,omitempty"`
	DestinationArrays       []*NsArraySummary `json:"destination_arrays,omitempty"`
	BytesMigrated           *int64            `json:"bytes_migrated,omitempty"`
	BytesRemaining          *int64            `json:"bytes_remaining,omitempty"`
	StartTime               *int64            `json:"start_time,omitempty"`
	EstimatedCompletionTime *int64            `json:"estimated_completion_time,omitempty"`
}

// EncodeNsArrayUnassignMigStatus - Transform NsArrayUnassignMigStatus to nsArrayUnassignMigStatus type
func EncodeNsArrayUnassignMigStatus(request interface{}) (*nsArrayUnassignMigStatus, error) {
	reqNsArrayUnassignMigStatus := request.(*NsArrayUnassignMigStatus)
	byte, err := json.Marshal(reqNsArrayUnassignMigStatus)
	if err != nil {
		return nil, err
	}
	respNsArrayUnassignMigStatusPtr := &nsArrayUnassignMigStatus{}
	err = json.Unmarshal(byte, respNsArrayUnassignMigStatusPtr)
	return respNsArrayUnassignMigStatusPtr, err
}

// DecodeNsArrayUnassignMigStatus Transform nsArrayUnassignMigStatus to NsArrayUnassignMigStatus type
func DecodeNsArrayUnassignMigStatus(request interface{}) (*NsArrayUnassignMigStatus, error) {
	reqNsArrayUnassignMigStatus := request.(*nsArrayUnassignMigStatus)
	byte, err := json.Marshal(reqNsArrayUnassignMigStatus)
	if err != nil {
		return nil, err
	}
	respNsArrayUnassignMigStatus := &NsArrayUnassignMigStatus{}
	err = json.Unmarshal(byte, respNsArrayUnassignMigStatus)
	return respNsArrayUnassignMigStatus, err
}
