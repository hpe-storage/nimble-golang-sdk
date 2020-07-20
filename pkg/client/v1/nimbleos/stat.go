// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Stat - Access generic stats interface via REST for internal testing.
// Export StatFields for advance operations like search filter etc.
var StatFields *Stat

func init() {

	StatFields = &Stat{
		Scope:     "scope",
		DomainId:  "domain_id",
		SetId:     "set_id",
		VolIds:    "vol_ids",
		Sensors:   "sensors",
		PoolId:    "pool_id",
		ArrayName: "array_name",
		VolId:     "vol_id",
	}
}

type Stat struct {
	// Scope - Stat scope parameter to pass to stats reader.  Optional input attribute.
	Scope string `json:"scope,omitempty"`
	// DomainId - Stat domain_id to pass to stats reader.  Required input attribute with set_id if vol_id not used.
	DomainId string `json:"domain_id,omitempty"`
	// SetId - Stat set_id to pass to stats reader.  Required input attribute with domain_id if vol_id not used.
	SetId string `json:"set_id,omitempty"`
	// VolIds - A CSV list of REST vol-id to pass to stats reader.  Required input attribute if domain_id/set_id not used.
	VolIds string `json:"vol_ids,omitempty"`
	// Sensors - Requested stats sensor list.  Optional input attribute.  Defaults to all sensors if not supplied.
	Sensors string `json:"sensors,omitempty"`
	// Starttime - Retrieve stats with createTime greater than or equal to this time.  Optional input attribute.  Defaults to now if not supplied.
	Starttime int64 `json:"starttime,omitempty"`
	// Endtime - Retrieve stats with createTime less than or equal to this time.  Optional input attribute.  Defaults to now if not supplied.
	Endtime int64 `json:"endtime,omitempty"`
	// Interval - Retrieve stats with this time interval between values.  Typical use is 1, 60, or 3600 for per-second, per-minute, and per-hour stats.  Optional.  Defaults to 1 if not supplied.
	Interval int64 `json:"interval,omitempty"`
	// Cumulative - Retrieve stat sensor cumulative data.
	Cumulative *bool `json:"cumulative,omitempty"`
	// PoolId - Stats summed for pool specified by this ID.  Optional input attribute.  Defaults to whole group if pool_id is not supplied.
	PoolId string `json:"pool_id,omitempty"`
	// ArrayName - Retrieve stats from this array only.  Optional parameter.
	ArrayName string `json:"array_name,omitempty"`
	// SensorData - Response sensor data.
	SensorData []*NsSensorData `json:"sensor_data,omitempty"`
	// SensorCumulativeData - Response sensor cumulative data.
	SensorCumulativeData []*NsSensorCumulativeData `json:"sensor_cumulative_data,omitempty"`
	// VolId - The REST ID of a volume.
	VolId string `json:"vol_id,omitempty"`
}

// sdk internal struct
type stat struct {
	Scope                *string                   `json:"scope,omitempty"`
	DomainId             *string                   `json:"domain_id,omitempty"`
	SetId                *string                   `json:"set_id,omitempty"`
	VolIds               *string                   `json:"vol_ids,omitempty"`
	Sensors              *string                   `json:"sensors,omitempty"`
	Starttime            *int64                    `json:"starttime,omitempty"`
	Endtime              *int64                    `json:"endtime,omitempty"`
	Interval             *int64                    `json:"interval,omitempty"`
	Cumulative           *bool                     `json:"cumulative,omitempty"`
	PoolId               *string                   `json:"pool_id,omitempty"`
	ArrayName            *string                   `json:"array_name,omitempty"`
	SensorData           []*NsSensorData           `json:"sensor_data,omitempty"`
	SensorCumulativeData []*NsSensorCumulativeData `json:"sensor_cumulative_data,omitempty"`
	VolId                *string                   `json:"vol_id,omitempty"`
}

// EncodeStat - Transform Stat to stat type
func EncodeStat(request interface{}) (*stat, error) {
	reqStat := request.(*Stat)
	byte, err := json.Marshal(reqStat)
	if err != nil {
		return nil, err
	}
	respStatPtr := &stat{}
	err = json.Unmarshal(byte, respStatPtr)
	return respStatPtr, err
}

// DecodeStat Transform stat to Stat type
func DecodeStat(request interface{}) (*Stat, error) {
	reqStat := request.(*stat)
	byte, err := json.Marshal(reqStat)
	if err != nil {
		return nil, err
	}
	respStat := &Stat{}
	err = json.Unmarshal(byte, respStat)
	return respStat, err
}
