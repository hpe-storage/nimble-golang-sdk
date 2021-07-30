// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsDiskSetAttrFields provides field names to use in filter parameters, for example.
var NsDiskSetAttrFields *NsDiskSetAttrFieldHandles

func init() {
	NsDiskSetAttrFields = &NsDiskSetAttrFieldHandles{
		Driveset:            "driveset",
		SwState:             "sw_state",
		IsFlashShelf:        "is_flash_shelf",
		IsCapacityValid:     "is_capacity_valid",
		UsableCapacity:      "usable_capacity",
		RawCapacity:         "raw_capacity",
		UsableCacheCapacity: "usable_cache_capacity",
		RawCacheCapacity:    "raw_cache_capacity",
		AveMbPs:             "ave_mb_ps",
		AveSegmentPs:        "ave_segment_ps",
		AveTtc:              "ave_ttc",
		PctCompletion:       "pct_completion",
		PauseState:          "pause_state",
	}
}

// NsDiskSetAttr - A shelf logical attributes.
type NsDiskSetAttr struct {
	// Driveset - Driveset index for this shelf.
	Driveset *int64 `json:"driveset,omitempty"`
	// SwState - Software state.
	SwState *NsShelfSwState `json:"sw_state,omitempty"`
	// IsFlashShelf - Is this a all-flash-shelf.
	IsFlashShelf *bool `json:"is_flash_shelf,omitempty"`
	// IsCapacityValid - Is the capacity fields in this data struct valid.
	IsCapacityValid *bool `json:"is_capacity_valid,omitempty"`
	// UsableCapacity - Estimated usable capacity for this shelf.
	UsableCapacity *int64 `json:"usable_capacity,omitempty"`
	// RawCapacity - Hdd raw capacity for this shelf.
	RawCapacity *int64 `json:"raw_capacity,omitempty"`
	// UsableCacheCapacity - Estimated usable cache capacity for this shelf.
	UsableCacheCapacity *int64 `json:"usable_cache_capacity,omitempty"`
	// RawCacheCapacity - Raw cache capacity for this shelf.
	RawCacheCapacity *int64 `json:"raw_cache_capacity,omitempty"`
	// AveMbPs - Average evacuation speed in MB/s; valid only when sw_state is evacuating, ie. evacuation is underway.
	AveMbPs *int64 `json:"ave_mb_ps,omitempty"`
	// AveSegmentPs - Average evacuation speed in segments/sec; valid only when sw_state is evacuating, ie. evacuation is underway.
	AveSegmentPs *int64 `json:"ave_segment_ps,omitempty"`
	// AveTtc - Average time to complete in seconds; valid only when sw_state is evacuating, ie. evacuation is underway.
	AveTtc *int64 `json:"ave_ttc,omitempty"`
	// PctCompletion - Evacuation percent completion; valid only when sw_state is evacuating, ie. evacuation is underway.
	PctCompletion *int64 `json:"pct_completion,omitempty"`
	// PauseState - State of evacuation, paused or in-progress; valid only when sw_state is evacuating, ie. evacuation is underway.
	PauseState *int64 `json:"pause_state,omitempty"`
}

// NsDiskSetAttrFieldHandles provides a string representation for each AccessControlRecord field.
type NsDiskSetAttrFieldHandles struct {
	Driveset            string
	SwState             string
	IsFlashShelf        string
	IsCapacityValid     string
	UsableCapacity      string
	RawCapacity         string
	UsableCacheCapacity string
	RawCacheCapacity    string
	AveMbPs             string
	AveSegmentPs        string
	AveTtc              string
	PctCompletion       string
	PauseState          string
}
