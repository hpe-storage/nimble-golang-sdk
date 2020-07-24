// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// Job - Jobs are operations in progress in the system.
// Export JobFields for advance operations like search filter etc.
var JobFields *Job

func init() {
	CurrentPhaseDescriptionfield := "current_phase_description"
	Descriptionfield := "description"
	IDfield := "id"
	Namefield := "name"
	ObjectIdfield := "object_id"
	ParentJobIdfield := "parent_job_id"

	JobFields = &Job{
		CurrentPhaseDescription: &CurrentPhaseDescriptionfield,
		Description:             &Descriptionfield,
		ID:                      &IDfield,
		Name:                    &Namefield,
		ObjectId:                &ObjectIdfield,
		ParentJobId:             &ParentJobIdfield,
	}
}

type Job struct {
	// CompletionTime - Completion time of the job.
	CompletionTime *int64 `json:"completion_time,omitempty"`
	// CreationTime - Time when this job was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// CurrentPhase - Phase number of the job in progress.
	CurrentPhase *int64 `json:"current_phase,omitempty"`
	// CurrentPhaseDescription - Description of the current phase of the job.
	CurrentPhaseDescription *string `json:"current_phase_description,omitempty"`
	// Description - Description of the job.
	Description *string `json:"description,omitempty"`
	// ID - Identifier for job.
	ID *string `json:"id,omitempty"`
	// Name - Name of the job.
	Name *string `json:"name,omitempty"`
	// LastModified - Time of the last update from the job.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ObjectId - Identifier for object being acted upon.
	ObjectId *string `json:"object_id,omitempty"`
	// OpType - Type of operation.
	OpType *NsOperationType `json:"op_type,omitempty"`
	// Type - Job type.
	Type *NsJobType `json:"type,omitempty"`
	// ParentJobId - Identifier of parent job.
	ParentJobId *string `json:"parent_job_id,omitempty"`
	// PercentComplete - Progress of the job as a percentage.
	PercentComplete *int64 `json:"percent_complete,omitempty"`
	// Request - Original request that the job is responsible for.
	Request *NsRequest `json:"request,omitempty"`
	// Response - Response from the operation as the job executes.
	Response *NsResponse `json:"response,omitempty"`
	// State - Status of the job.
	State *NsJobStatus `json:"state,omitempty"`
	// Result - Result of the job.
	Result *NsJobResult `json:"result,omitempty"`
	// TotalPhases - Total number of phases of the job.
	TotalPhases *int64 `json:"total_phases,omitempty"`
}
