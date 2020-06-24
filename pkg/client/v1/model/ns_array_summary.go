/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArraySummary - Array summary information.
// Export NsArraySummaryFields for advance operations like search filter etc.
var NsArraySummaryFields *NsArraySummary

func init(){
	IDfield:= "id"
	ArrayIDfield:= "array_id"
	Namefield:= "name"
	ArrayNamefield:= "array_name"
		
	NsArraySummaryFields= &NsArraySummary{
		ID: &IDfield,
		ArrayID: &ArrayIDfield,
		Name: &Namefield,
		ArrayName: &ArrayNamefield,
		
	}
}

type NsArraySummary struct {
   
    // Array API ID.
    
 	ID *string `json:"id,omitempty"`
   
    // Array API ID.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Unique name of array.
    
 	Name *string `json:"name,omitempty"`
   
    // Unique name of array.
    
 	ArrayName *string `json:"array_name,omitempty"`
}
