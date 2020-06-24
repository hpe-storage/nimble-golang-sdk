/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Job - Jobs are operations in progress in the system.
// Export JobFields for advance operations like search filter etc.
var JobFields *Job

func init(){
	CurrentPhaseDescriptionfield:= "current_phase_description"
	Descriptionfield:= "description"
	IDfield:= "id"
	Namefield:= "name"
	ObjectIDfield:= "object_id"
	ParentJobIDfield:= "parent_job_id"
		
	JobFields= &Job{
		CurrentPhaseDescription: &CurrentPhaseDescriptionfield,
		Description: &Descriptionfield,
		ID: &IDfield,
		Name: &Namefield,
		ObjectID: &ObjectIDfield,
		ParentJobID: &ParentJobIDfield,
		
	}
}

type Job struct {
   
    // Completion time of the job.
    
   	CompletionTime *int64 `json:"completion_time,omitempty"`
   
    // Time when this job was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Phase number of the job in progress.
    
   	CurrentPhase *int64 `json:"current_phase,omitempty"`
   
    // Description of the current phase of the job.
    
 	CurrentPhaseDescription *string `json:"current_phase_description,omitempty"`
   
    // Description of the job.
    
 	Description *string `json:"description,omitempty"`
   
    // Identifier for job.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the job.
    
 	Name *string `json:"name,omitempty"`
   
    // Time of the last update from the job.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Identifier for object being acted upon.
    
 	ObjectID *string `json:"object_id,omitempty"`
   
    // Type of operation.
    
   	OpType *NsOperationType `json:"op_type,omitempty"`
   
    // Job type.
    
   	Type *NsJobType `json:"type,omitempty"`
   
    // Identifier of parent job.
    
 	ParentJobID *string `json:"parent_job_id,omitempty"`
   
    // Progress of the job as a percentage.
    
   	PercentComplete *int64 `json:"percent_complete,omitempty"`
   
    // Original request that the job is responsible for.
    
	Request *NsRequest `json:"request,omitempty"`
   
    // Response from the operation as the job executes.
    
	Response *NsResponse `json:"response,omitempty"`
   
    // Status of the job.
    
   	State *NsJobStatus `json:"state,omitempty"`
   
    // Result of the job.
    
   	Result *NsJobResult `json:"result,omitempty"`
   
    // Total number of phases of the job.
    
   	TotalPhases *int64 `json:"total_phases,omitempty"`
}
