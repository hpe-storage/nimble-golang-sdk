// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// Job Service - Jobs are operations in progress in the system.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// JobService type
type JobService struct {
	objectSet *client.JobObjectSet
}

// NewJobService - method to initialize "JobService"
func NewJobService(gs *NsGroupService) *JobService {
	objectSet := gs.client.GetJobObjectSet()
	return &JobService{objectSet: objectSet}
}

// GetJobs - method returns a array of pointers of type "Jobs"
func (svc *JobService) GetJobs(params *param.GetParams) ([]*nimbleos.Job, error) {
	jobResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return jobResp, nil
}

// CreateJob - method creates a "Job"
func (svc *JobService) CreateJob(obj *nimbleos.Job) (*nimbleos.Job, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateJob: invalid parameter specified, %v", obj)
	}

	jobResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return jobResp, nil
}

// UpdateJob - method modifies  the "Job"
func (svc *JobService) UpdateJob(id string, obj *nimbleos.Job) (*nimbleos.Job, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateJob: invalid parameter specified, %v", obj)
	}

	jobResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return jobResp, nil
}

// GetJobById - method returns a pointer to "Job"
func (svc *JobService) GetJobById(id string) (*nimbleos.Job, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetJobById: invalid parameter specified, %v", id)
	}

	jobResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return jobResp, nil
}

// GetJobByName - method returns a pointer "Job"
func (svc *JobService) GetJobByName(name string) (*nimbleos.Job, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: param.NewString(nimbleos.VolumeFields.Name),
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	jobResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(jobResp) == 0 {
		return nil, nil
	}

	return jobResp[0], nil
}

// DeleteJob - deletes the "Job"
func (svc *JobService) DeleteJob(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteJob: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
