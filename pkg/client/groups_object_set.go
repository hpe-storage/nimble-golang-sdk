// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Group is a collection of arrays operating together organized into storage pools.
const (
    groupPath = "groups"
)

// GroupObjectSet
type GroupObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Group object
func (objectSet *GroupObjectSet) CreateObject(payload *model.Group) (*model.Group, error) {
	groupObjectSetResp, err := objectSet.Client.Post(groupPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if groupObjectSetResp == nil {
		return nil,nil
	}
	return groupObjectSetResp.(*model.Group), err
}

// UpdateObject Modify existing Group object
func (objectSet *GroupObjectSet) UpdateObject(id string, payload *model.Group) (*model.Group, error) {
	groupObjectSetResp, err := objectSet.Client.Put(groupPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if groupObjectSetResp == nil {
		return nil,nil
	}
	return groupObjectSetResp.(*model.Group), err
}

// DeleteObject deletes the Group object with the specified ID
func (objectSet *GroupObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(groupPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a Group object with the given ID
func (objectSet *GroupObjectSet) GetObject(id string) (*model.Group, error) {
	groupObjectSetResp, err := objectSet.Client.Get(groupPath, id, model.Group{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if groupObjectSetResp == nil {
		return nil,nil
	}
	return groupObjectSetResp.(*model.Group), err
}

// GetObjectList returns the list of Group objects
func (objectSet *GroupObjectSet) GetObjectList() ([]*model.Group, error) {
	groupObjectSetResp, err := objectSet.Client.List(groupPath)
	if err != nil {
		return nil, err
	}
	return buildGroupObjectSet(groupObjectSetResp), err
}

// GetObjectListFromParams returns the list of Group objects using the given params query info
func (objectSet *GroupObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Group, error) {
	groupObjectSetResp, err := objectSet.Client.ListFromParams(groupPath, params)
	if err != nil {
		return nil, err
	}
	return buildGroupObjectSet(groupObjectSetResp), err
}

// generated function to build the appropriate response types
func buildGroupObjectSet(response interface{}) ([]*model.Group) {
	values := reflect.ValueOf(response)
	results := make([]*model.Group, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Group{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}