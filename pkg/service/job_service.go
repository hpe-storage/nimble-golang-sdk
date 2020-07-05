// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Job Service - Jobs are operations in progress in the system.

import (
	"fmt"
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
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	jobResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return jobResp,nil
}

// CreateJob - method creates a "Job"
func (svc *JobService) CreateJob(obj *model.Job) (*model.Job, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	jobResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return jobResp,nil
}

// UpdateJob - method modifies  the "Job" 
func (svc *JobService) UpdateJob(id string, obj *model.Job) (*model.Job, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	jobResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return jobResp,nil
}

// GetJobById - method returns a pointer to "Job"
func (svc *JobService) GetJobById(id string) (*model.Job, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	jobResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return jobResp,nil
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
	jobResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(jobResp) == 0 {
    	return nil, nil
    }
    
	return jobResp[0],nil
}	

// DeleteJob - deletes the "Job"
func (svc *JobService) DeleteJob(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
