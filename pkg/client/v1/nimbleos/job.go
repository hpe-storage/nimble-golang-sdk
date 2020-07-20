// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Job - Jobs are operations in progress in the system.
// Export JobFields for advance operations like search filter etc.
var JobFields *Job

func init() {

	JobFields = &Job{
		CurrentPhaseDescription: "current_phase_description",
		Description:             "description",
		ID:                      "id",
		Name:                    "name",
		ObjectId:                "object_id",
		ParentJobId:             "parent_job_id",
	}
}

type Job struct {
	// CompletionTime - Completion time of the job.
	CompletionTime int64 `json:"completion_time,omitempty"`
	// CreationTime - Time when this job was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// CurrentPhase - Phase number of the job in progress.
	CurrentPhase int64 `json:"current_phase,omitempty"`
	// CurrentPhaseDescription - Description of the current phase of the job.
	CurrentPhaseDescription string `json:"current_phase_description,omitempty"`
	// Description - Description of the job.
	Description string `json:"description,omitempty"`
	// ID - Identifier for job.
	ID string `json:"id,omitempty"`
	// Name - Name of the job.
	Name string `json:"name,omitempty"`
	// LastModified - Time of the last update from the job.
	LastModified int64 `json:"last_modified,omitempty"`
	// ObjectId - Identifier for object being acted upon.
	ObjectId string `json:"object_id,omitempty"`
	// OpType - Type of operation.
	OpType *NsOperationType `json:"op_type,omitempty"`
	// Type - Job type.
	Type *NsJobType `json:"type,omitempty"`
	// ParentJobId - Identifier of parent job.
	ParentJobId string `json:"parent_job_id,omitempty"`
	// PercentComplete - Progress of the job as a percentage.
	PercentComplete int64 `json:"percent_complete,omitempty"`
	// Request - Original request that the job is responsible for.
	Request *NsRequest `json:"request,omitempty"`
	// Response - Response from the operation as the job executes.
	Response *NsResponse `json:"response,omitempty"`
	// State - Status of the job.
	State *NsJobStatus `json:"state,omitempty"`
	// Result - Result of the job.
	Result *NsJobResult `json:"result,omitempty"`
	// TotalPhases - Total number of phases of the job.
	TotalPhases int64 `json:"total_phases,omitempty"`
}

// sdk internal struct
type job struct {
	CompletionTime          *int64           `json:"completion_time,omitempty"`
	CreationTime            *int64           `json:"creation_time,omitempty"`
	CurrentPhase            *int64           `json:"current_phase,omitempty"`
	CurrentPhaseDescription *string          `json:"current_phase_description,omitempty"`
	Description             *string          `json:"description,omitempty"`
	ID                      *string          `json:"id,omitempty"`
	Name                    *string          `json:"name,omitempty"`
	LastModified            *int64           `json:"last_modified,omitempty"`
	ObjectId                *string          `json:"object_id,omitempty"`
	OpType                  *NsOperationType `json:"op_type,omitempty"`
	Type                    *NsJobType       `json:"type,omitempty"`
	ParentJobId             *string          `json:"parent_job_id,omitempty"`
	PercentComplete         *int64           `json:"percent_complete,omitempty"`
	Request                 *NsRequest       `json:"request,omitempty"`
	Response                *NsResponse      `json:"response,omitempty"`
	State                   *NsJobStatus     `json:"state,omitempty"`
	Result                  *NsJobResult     `json:"result,omitempty"`
	TotalPhases             *int64           `json:"total_phases,omitempty"`
}

// EncodeJob - Transform Job to job type
func EncodeJob(request interface{}) (*job, error) {
	reqJob := request.(*Job)
	byte, err := json.Marshal(reqJob)
	if err != nil {
		return nil, err
	}
	respJobPtr := &job{}
	err = json.Unmarshal(byte, respJobPtr)
	return respJobPtr, err
}

// DecodeJob Transform job to Job type
func DecodeJob(request interface{}) (*Job, error) {
	reqJob := request.(*job)
	byte, err := json.Marshal(reqJob)
	if err != nil {
		return nil, err
	}
	respJob := &Job{}
	err = json.Unmarshal(byte, respJob)
	return respJob, err
}
