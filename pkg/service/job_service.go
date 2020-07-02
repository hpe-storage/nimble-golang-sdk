// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Job Service - Jobs are operations in progress in the system.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// JobService type 
type JobService struct {
	objectSet *client.JobObjectSet
}

// NewJobService - method to initialize "JobService" 
func NewJobService(gs *NsGroupService) (*JobService) {
	objectSet := gs.client.GetJobObjectSet()
	return &JobService{objectSet: objectSet}
}

// GetJobs - method returns a array of pointers of type "Jobs"
func (svc *JobService) GetJobs(params *util.GetParams) ([]*model.Job, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateJob - method creates a "Job"
func (svc *JobService) CreateJob(obj *model.Job) (*model.Job, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateJob - method modifies  the "Job" 
func (svc *JobService) UpdateJob(id string, obj *model.Job) (*model.Job, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetJobById - method returns a pointer to "Job"
func (svc *JobService) GetJobById(id string) (*model.Job, error) {
	return svc.objectSet.GetObject(id)
}

// GetJobByName - method returns a pointer "Job" 
func (svc *JobService) GetJobByName(name string) (*model.Job, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteJob - deletes the "Job"
func (svc *JobService) DeleteJob(id string) error {
	return svc.objectSet.DeleteObject(id)
}
