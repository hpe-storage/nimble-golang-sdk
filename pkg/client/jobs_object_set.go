// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
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
	return nil, fmt.Errorf("Unsupported operation 'create' on Job")
}

// UpdateObject Modify existing Job object
func (objectSet *JobObjectSet) UpdateObject(id string, payload *model.Job) (*model.Job, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Job")
}

// DeleteObject deletes the Job object with the specified ID
func (objectSet *JobObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Job")
}

// GetObject returns a Job object with the given ID
func (objectSet *JobObjectSet) GetObject(id string) (*model.Job, error) {
	jobObjectSetResp, err := objectSet.Client.Get(jobPath, id, model.Job{})
	if err != nil {
		return nil, err
	}

	// null check
	if jobObjectSetResp == nil {
		return nil, nil
	}
	return jobObjectSetResp.(*model.Job), err
}

// GetObjectList returns the list of Job objects
func (objectSet *JobObjectSet) GetObjectList() ([]*model.Job, error) {
	jobObjectSetResp, err := objectSet.Client.List(jobPath)
	if err != nil {
		return nil, err
	}
	return buildJobObjectSet(jobObjectSetResp), err
}

// GetObjectListFromParams returns the list of Job objects using the given params query info
func (objectSet *JobObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Job, error) {
	jobObjectSetResp, err := objectSet.Client.ListFromParams(jobPath, params)
	if err != nil {
		return nil, err
	}
	return buildJobObjectSet(jobObjectSetResp), err
}

// generated function to build the appropriate response types
func buildJobObjectSet(response interface{}) []*model.Job {
	values := reflect.ValueOf(response)
	results := make([]*model.Job, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Job{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
