/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArrayUnassignMigStatus - Data migration status for array being unassigned from its pool.
// Export NsArrayUnassignMigStatusFields for advance operations like search filter etc.
var NsArrayUnassignMigStatusFields *NsArrayUnassignMigStatus

func init(){
	IDfield:= "id"
	Namefield:= "name"
		
	NsArrayUnassignMigStatusFields= &NsArrayUnassignMigStatus{
		ID: &IDfield,
		Name: &Namefield,
		
	}
}

type NsArrayUnassignMigStatus struct {
   
    // Unique identifier of the array being unassigned.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the array being unassigned.
    
 	Name *string `json:"name,omitempty"`
   
    // List of arrays to which data is being migrated.
    
   	DestinationArrays []*NsArraySummary `json:"destination_arrays,omitempty"`
   
    // Number of bytes already migrated to the destination arrays.
    
   	BytesMigrated *int64 `json:"bytes_migrated,omitempty"`
   
    // Number of bytes yet to be migrated to the destination arrays.
    
   	BytesRemaining *int64 `json:"bytes_remaining,omitempty"`
   
    // Time when array unassign operation started.
    
   	StartTime *int64 `json:"start_time,omitempty"`
   
    // Estimated completion time. 0 if unknown.
    
   	EstimatedCompletionTime *int64 `json:"estimated_completion_time,omitempty"`
}
