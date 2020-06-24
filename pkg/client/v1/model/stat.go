/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Stat - Access generic stats interface via REST for internal testing.
// Export StatFields for advance operations like search filter etc.
var StatFields *Stat

func init(){
	Scopefield:= "scope"
	DomainIDfield:= "domain_id"
	SetIDfield:= "set_id"
	VolIDsfield:= "vol_ids"
	Sensorsfield:= "sensors"
	PoolIDfield:= "pool_id"
	ArrayNamefield:= "array_name"
	VolIDfield:= "vol_id"
		
	StatFields= &Stat{
		Scope: &Scopefield,
		DomainID: &DomainIDfield,
		SetID: &SetIDfield,
		VolIDs: &VolIDsfield,
		Sensors: &Sensorsfield,
		PoolID: &PoolIDfield,
		ArrayName: &ArrayNamefield,
		VolID: &VolIDfield,
		
	}
}

type Stat struct {
   
    // Stat scope parameter to pass to stats reader.  Optional input attribute.
    
 	Scope *string `json:"scope,omitempty"`
   
    // Stat domain_id to pass to stats reader.  Required input attribute with set_id if vol_id not used.
    
 	DomainID *string `json:"domain_id,omitempty"`
   
    // Stat set_id to pass to stats reader.  Required input attribute with domain_id if vol_id not used.
    
 	SetID *string `json:"set_id,omitempty"`
   
    // A CSV list of REST vol-id to pass to stats reader.  Required input attribute if domain_id/set_id not used.
    
 	VolIDs *string `json:"vol_ids,omitempty"`
   
    // Requested stats sensor list.  Optional input attribute.  Defaults to all sensors if not supplied.
    
 	Sensors *string `json:"sensors,omitempty"`
   
    // Retrieve stats with createTime greater than or equal to this time.  Optional input attribute.  Defaults to now if not supplied.
    
   	Starttime *int64 `json:"starttime,omitempty"`
   
    // Retrieve stats with createTime less than or equal to this time.  Optional input attribute.  Defaults to now if not supplied.
    
   	Endtime *int64 `json:"endtime,omitempty"`
   
    // Retrieve stats with this time interval between values.  Typical use is 1, 60, or 3600 for per-second, per-minute, and per-hour stats.  Optional.  Defaults to 1 if not supplied.
    
   	Interval *int64 `json:"interval,omitempty"`
   
    // Retrieve stat sensor cumulative data.
    
 	Cumulative *bool `json:"cumulative,omitempty"`
   
    // Stats summed for pool specified by this ID.  Optional input attribute.  Defaults to whole group if pool_id is not supplied.
    
 	PoolID *string `json:"pool_id,omitempty"`
   
    // Retrieve stats from this array only.  Optional parameter.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // Response sensor data.
    
   	SensorData []*NsSensorData `json:"sensor_data,omitempty"`
   
    // Response sensor cumulative data.
    
   	SensorCumulativeData []*NsSensorCumulativeData `json:"sensor_cumulative_data,omitempty"`
   
    // The REST ID of a volume.
    
 	VolID *string `json:"vol_id,omitempty"`
}
