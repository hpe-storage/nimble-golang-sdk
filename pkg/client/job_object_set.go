// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Jobs are operations in progress in the system.
const (
    jobPath = "jobs"
)


// JobObjectSet
type JobObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Job object
func (objectSet *JobObjectSet) CreateObject(payload *model.Job) (*model.Job, error) {
	response, err := objectSet.Client.Post(jobPath, payload)
	return response.(*model.Job), err
}

// UpdateObject Modify existing Job object
func (objectSet *JobObjectSet) UpdateObject(id string, payload *model.Job) (*model.Job, error) {
	response, err := objectSet.Client.Put(jobPath, id, payload)
	return response.(*model.Job), err
}

// DeleteObject deletes the Job object with the specified ID
func (objectSet *JobObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(jobPath, id)
}

// GetObject returns a Job object with the given ID
func (objectSet *JobObjectSet) GetObject(id string) (*model.Job, error) {
	response, err := objectSet.Client.Get(jobPath, id, model.Job{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Job), err
}

// GetObjectList returns the list of Job objects
func (objectSet *JobObjectSet) GetObjectList() ([]*model.Job, error) {
	response, err := objectSet.Client.List(jobPath)
	return buildJobObjectSet(response), err
}

// GetObjectListFromParams returns the list of Job objects using the given params query info
func (objectSet *JobObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Job, error) {
	response, err := objectSet.Client.ListFromParams(jobPath, params)
	return buildJobObjectSet(response), err
}

// generated function to build the appropriate response types
func buildJobObjectSet(response interface{}) ([]*model.Job) {
	values := reflect.ValueOf(response)
	results := make([]*model.Job, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Job{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}