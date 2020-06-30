// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Job Service - Jobs are operations in progress in the system.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetJobsWithFields - method returns a array of pointers of type "Job" 
func (svc *JobService) GetJobsWithFields(fields []string) ([]*model.Job, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateJob - method creates a "Job"
func (svc *JobService) CreateJob(obj *model.Job) (*model.Job, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditJob - method modifies  the "Job" 
func (svc *JobService) EditJob(id string, obj *model.Job) (*model.Job, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyJob - private method for more than one element check. 
func onlyJob(objs []*model.Job) (*model.Job, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Job found with the given filter")
	}

	return objs[0], nil
}

 
// GetJobsByID - method returns associative a array of pointers of type "Job", filter by Id
func (svc *JobService) GetJobsByID(pool *model.Pool, fields []string) (map[string]*model.Job, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetJobs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Job)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetJobById - method returns a pointer to "Job"
func (svc *JobService) GetJobById(id string) (*model.Job, error) {
	return svc.objectSet.GetObject(id)
}

// GetJobsByName - method returns a associative array of pointers of type "Job", filter by name 
func (svc *JobService) GetJobsByName(pool *model.Pool, fields []string) (map[string]*model.Job, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetJobs(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Job)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
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
	return onlyJob(objs)
}	

