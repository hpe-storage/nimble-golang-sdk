// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArrayDetail - Detailed array information.
// Export NsArrayDetailFields for advance operations like search filter etc.
var NsArrayDetailFields *NsArrayDetail

func init() {

	NsArrayDetailFields = &NsArrayDetail{
		ID:        "id",
		ArrayId:   "array_id",
		Name:      "name",
		ArrayName: "array_name",
	}
}

type NsArrayDetail struct {
	// ID - Array API ID.
	ID string `json:"id,omitempty"`
	// ArrayId - Array API ID.
	ArrayId string `json:"array_id,omitempty"`
	// Name - Unique name of array.
	Name string `json:"name,omitempty"`
	// ArrayName - Unique name of array.
	ArrayName string `json:"array_name,omitempty"`
	// EvacTime - Start time of array evacuation.
	EvacTime int64 `json:"evac_time,omitempty"`
	// EvacUsage - Initial data in the array.
	EvacUsage int64 `json:"evac_usage,omitempty"`
	// UsableCapacity - Usable capacity of the array.
	UsableCapacity int64 `json:"usable_capacity,omitempty"`
	// Usage - Usage of the array.
	Usage int64 `json:"usage,omitempty"`
	// VolUsageCompressedBytes - Usage of volumes in the array.
	VolUsageCompressedBytes int64 `json:"vol_usage_compressed_bytes,omitempty"`
	// SnapUsageCompressedBytes - Usage of snapshots in the array.
	SnapUsageCompressedBytes int64 `json:"snap_usage_compressed_bytes,omitempty"`
	// UsageValid - Indicate whether usage of the array is valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// Migrate - Migrate status of array.
	Migrate *NsPoolMigrate `json:"migrate,omitempty"`
}

// sdk internal struct
type nsArrayDetail struct {
	ID                       *string        `json:"id,omitempty"`
	ArrayId                  *string        `json:"array_id,omitempty"`
	Name                     *string        `json:"name,omitempty"`
	ArrayName                *string        `json:"array_name,omitempty"`
	EvacTime                 *int64         `json:"evac_time,omitempty"`
	EvacUsage                *int64         `json:"evac_usage,omitempty"`
	UsableCapacity           *int64         `json:"usable_capacity,omitempty"`
	Usage                    *int64         `json:"usage,omitempty"`
	VolUsageCompressedBytes  *int64         `json:"vol_usage_compressed_bytes,omitempty"`
	SnapUsageCompressedBytes *int64         `json:"snap_usage_compressed_bytes,omitempty"`
	UsageValid               *bool          `json:"usage_valid,omitempty"`
	Migrate                  *NsPoolMigrate `json:"migrate,omitempty"`
}

// EncodeNsArrayDetail - Transform NsArrayDetail to nsArrayDetail type
func EncodeNsArrayDetail(request interface{}) (*nsArrayDetail, error) {
	reqNsArrayDetail := request.(*NsArrayDetail)
	byte, err := json.Marshal(reqNsArrayDetail)
	if err != nil {
		return nil, err
	}
	respNsArrayDetailPtr := &nsArrayDetail{}
	err = json.Unmarshal(byte, respNsArrayDetailPtr)
	return respNsArrayDetailPtr, err
}

// DecodeNsArrayDetail Transform nsArrayDetail to NsArrayDetail type
func DecodeNsArrayDetail(request interface{}) (*NsArrayDetail, error) {
	reqNsArrayDetail := request.(*nsArrayDetail)
	byte, err := json.Marshal(reqNsArrayDetail)
	if err != nil {
		return nil, err
	}
	respNsArrayDetail := &NsArrayDetail{}
	err = json.Unmarshal(byte, respNsArrayDetail)
	return respNsArrayDetail, err
}
