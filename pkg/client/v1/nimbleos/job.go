// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export JobFields provides field names to use in filter parameters, for example.
var JobFields *JobFieldHandles

func init() {
	JobFields = &JobFieldHandles{
		CompletionTime:          "completion_time",
		CreationTime:            "creation_time",
		CurrentPhase:            "current_phase",
		CurrentPhaseDescription: "current_phase_description",
		Description:             "description",
		ID:                      "id",
		Name:                    "name",
		LastModified:            "last_modified",
		ObjectId:                "object_id",
		OpType:                  "op_type",
		Type:                    "type",
		ParentJobId:             "parent_job_id",
		PercentComplete:         "percent_complete",
		Request:                 "request",
		Response:                "response",
		State:                   "state",
		Result:                  "result",
		TotalPhases:             "total_phases",
	}
}

// Job - Jobs are operations in progress in the system.
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

// JobFieldHandles provides a string representation for each AccessControlRecord field.
type JobFieldHandles struct {
	CompletionTime          string
	CreationTime            string
	CurrentPhase            string
	CurrentPhaseDescription string
	Description             string
	ID                      string
	Name                    string
	LastModified            string
	ObjectId                string
	OpType                  string
	Type                    string
	ParentJobId             string
	PercentComplete         string
	Request                 string
	Response                string
	State                   string
	Result                  string
	TotalPhases             string
}
