/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsPerfPolicySummary - Select fields containing performance policy.
// Export NsPerfPolicySummaryFields for advance operations like search filter etc.
var NsPerfPolicySummaryFields *NsPerfPolicySummary

func init(){
	Namefield:= "name"
		
	NsPerfPolicySummaryFields= &NsPerfPolicySummary{
		Name: &Namefield,
		
	}
}

type NsPerfPolicySummary struct {
   
    // Name of performance policy.
    
 	Name *string `json:"name,omitempty"`
}
