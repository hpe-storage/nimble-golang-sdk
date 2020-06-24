/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsDiskSetAttr - A shelf logical attributes.
// Export NsDiskSetAttrFields for advance operations like search filter etc.
var NsDiskSetAttrFields *NsDiskSetAttr

func init(){
		
	NsDiskSetAttrFields= &NsDiskSetAttr{
		
	}
}

type NsDiskSetAttr struct {
   
    // Driveset index for this shelf.
    
   	Driveset *int64 `json:"driveset,omitempty"`
   
    // Software state.
    
   	SwState *NsShelfSwState `json:"sw_state,omitempty"`
   
    // Is this a all-flash-shelf.
    
 	IsFlashShelf *bool `json:"is_flash_shelf,omitempty"`
   
    // Is the capacity fields in this data struct valid.
    
 	IsCapacityValID *bool `json:"is_capacity_valid,omitempty"`
   
    // Estimated usable capacity for this shelf.
    
   	UsableCapacity *int64 `json:"usable_capacity,omitempty"`
   
    // Hdd raw capacity for this shelf.
    
   	RawCapacity *int64 `json:"raw_capacity,omitempty"`
   
    // Estimated usable cache capacity for this shelf.
    
   	UsableCacheCapacity *int64 `json:"usable_cache_capacity,omitempty"`
   
    // Raw cache capacity for this shelf.
    
   	RawCacheCapacity *int64 `json:"raw_cache_capacity,omitempty"`
   
    // Average evacuation speed in MB/s; valid only when sw_state is evacuating, ie. evacuation is underway.
    
   	AveMbPs *int64 `json:"ave_mb_ps,omitempty"`
   
    // Average evacuation speed in segments/sec; valid only when sw_state is evacuating, ie. evacuation is underway.
    
   	AveSegmentPs *int64 `json:"ave_segment_ps,omitempty"`
   
    // Average time to complete in seconds; valid only when sw_state is evacuating, ie. evacuation is underway.
    
   	AveTtc *int64 `json:"ave_ttc,omitempty"`
   
    // Evacuation percent completion; valid only when sw_state is evacuating, ie. evacuation is underway.
    
   	PctCompletion *int64 `json:"pct_completion,omitempty"`
   
    // State of evacuation, paused or in-progress; valid only when sw_state is evacuating, ie. evacuation is underway.
    
   	PauseState *int64 `json:"pause_state,omitempty"`
}
