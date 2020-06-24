/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// HealthCheck - View the health of the group of arrays.
// Export HealthCheckFields for advance operations like search filter etc.
var HealthCheckFields *HealthCheck

func init(){
	IDfield:= "id"
	ArrayIDfield:= "array_id"
	CtrlrIDfield:= "ctrlr_id"
		
	HealthCheckFields= &HealthCheck{
		ID: &IDfield,
		ArrayID: &ArrayIDfield,
		CtrlrID: &CtrlrIDfield,
		
	}
}

type HealthCheck struct {
   
    // Identifier for the health check.
    
 	ID *string `json:"id,omitempty"`
   
    // Scope at which the health check is to be run.
    
   	Scope *NsHcfScope `json:"scope,omitempty"`
   
    // Context for aggregating health check results.
    
   	Context *NsHcfContext `json:"context,omitempty"`
   
    // Flag to indicate running the health checks and then report results.
    
 	OnDemand *bool `json:"on_demand,omitempty"`
   
    // Identifier of the array to which this result belongs.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Identifier of the controller to which this result belongs.
    
 	CtrlrID *string `json:"ctrlr_id,omitempty"`
   
    // List of health check errors for a particular element.
    
	ElementResult *NsHcfResult `json:"element_result,omitempty"`
}
