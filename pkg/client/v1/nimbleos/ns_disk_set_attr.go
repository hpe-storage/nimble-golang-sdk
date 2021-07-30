// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsDiskSetAttrFields provides field names to use in filter parameters, for example.
var NsDiskSetAttrFields *NsDiskSetAttrFieldHandles

func init() {
	fieldDriveset := "driveset"
	fieldSwState := "sw_state"
	fieldIsFlashShelf := "is_flash_shelf"
	fieldIsCapacityValid := "is_capacity_valid"
	fieldUsableCapacity := "usable_capacity"
	fieldRawCapacity := "raw_capacity"
	fieldUsableCacheCapacity := "usable_cache_capacity"
	fieldRawCacheCapacity := "raw_cache_capacity"
	fieldAveMbPs := "ave_mb_ps"
	fieldAveSegmentPs := "ave_segment_ps"
	fieldAveTtc := "ave_ttc"
	fieldPctCompletion := "pct_completion"
	fieldPauseState := "pause_state"

	NsDiskSetAttrFields = &NsDiskSetAttrFieldHandles{
		Driveset:            &fieldDriveset,
		SwState:             &fieldSwState,
		IsFlashShelf:        &fieldIsFlashShelf,
		IsCapacityValid:     &fieldIsCapacityValid,
		UsableCapacity:      &fieldUsableCapacity,
		RawCapacity:         &fieldRawCapacity,
		UsableCacheCapacity: &fieldUsableCacheCapacity,
		RawCacheCapacity:    &fieldRawCacheCapacity,
		AveMbPs:             &fieldAveMbPs,
		AveSegmentPs:        &fieldAveSegmentPs,
		AveTtc:              &fieldAveTtc,
		PctCompletion:       &fieldPctCompletion,
		PauseState:          &fieldPauseState,
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
	Driveset            *string
	SwState             *string
	IsFlashShelf        *string
	IsCapacityValid     *string
	UsableCapacity      *string
	RawCapacity         *string
	UsableCacheCapacity *string
	RawCacheCapacity    *string
	AveMbPs             *string
	AveSegmentPs        *string
	AveTtc              *string
	PctCompletion       *string
	PauseState          *string
}
