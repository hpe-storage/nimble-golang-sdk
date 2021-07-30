// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsDiskSetAttr - A shelf logical attributes.
// Export NsDiskSetAttrFields for advance operations like search filter etc.
var NsDiskSetAttrFields *NsDiskSetAttrStringFields

func init() {
	Drivesetfield := "driveset"
	SwStatefield := "sw_state"
	IsFlashShelffield := "is_flash_shelf"
	IsCapacityValidfield := "is_capacity_valid"
	UsableCapacityfield := "usable_capacity"
	RawCapacityfield := "raw_capacity"
	UsableCacheCapacityfield := "usable_cache_capacity"
	RawCacheCapacityfield := "raw_cache_capacity"
	AveMbPsfield := "ave_mb_ps"
	AveSegmentPsfield := "ave_segment_ps"
	AveTtcfield := "ave_ttc"
	PctCompletionfield := "pct_completion"
	PauseStatefield := "pause_state"

	NsDiskSetAttrFields = &NsDiskSetAttrStringFields{
		Driveset:            &Drivesetfield,
		SwState:             &SwStatefield,
		IsFlashShelf:        &IsFlashShelffield,
		IsCapacityValid:     &IsCapacityValidfield,
		UsableCapacity:      &UsableCapacityfield,
		RawCapacity:         &RawCapacityfield,
		UsableCacheCapacity: &UsableCacheCapacityfield,
		RawCacheCapacity:    &RawCacheCapacityfield,
		AveMbPs:             &AveMbPsfield,
		AveSegmentPs:        &AveSegmentPsfield,
		AveTtc:              &AveTtcfield,
		PctCompletion:       &PctCompletionfield,
		PauseState:          &PauseStatefield,
	}
}

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

// Struct for NsDiskSetAttrFields
type NsDiskSetAttrStringFields struct {
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
