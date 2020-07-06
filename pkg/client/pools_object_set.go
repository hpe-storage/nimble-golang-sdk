// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manage pools. Pools are an aggregation of arrays.
const (
    poolPath = "pools"
)

// PoolObjectSet
type PoolObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Pool object
func (objectSet *PoolObjectSet) CreateObject(payload *model.Pool) (*model.Pool, error) {
	poolObjectSetResp, err := objectSet.Client.Post(poolPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if poolObjectSetResp == nil {
		return nil,nil
	}
	return poolObjectSetResp.(*model.Pool), err
}

// UpdateObject Modify existing Pool object
func (objectSet *PoolObjectSet) UpdateObject(id string, payload *model.Pool) (*model.Pool, error) {
	poolObjectSetResp, err := objectSet.Client.Put(poolPath, id, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if poolObjectSetResp == nil {
		return nil,nil
	}
	return poolObjectSetResp.(*model.Pool), err
}

// DeleteObject deletes the Pool object with the specified ID
func (objectSet *PoolObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(poolPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a Pool object with the given ID
func (objectSet *PoolObjectSet) GetObject(id string) (*model.Pool, error) {
	poolObjectSetResp, err := objectSet.Client.Get(poolPath, id, model.Pool{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if poolObjectSetResp == nil {
		return nil,nil
	}
	return poolObjectSetResp.(*model.Pool), err
}

// GetObjectList returns the list of Pool objects
func (objectSet *PoolObjectSet) GetObjectList() ([]*model.Pool, error) {
	poolObjectSetResp, err := objectSet.Client.List(poolPath)
	if err != nil {
		return nil, err
	}
	return buildPoolObjectSet(poolObjectSetResp), err
}

// GetObjectListFromParams returns the list of Pool objects using the given params query info
func (objectSet *PoolObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Pool, error) {
	poolObjectSetResp, err := objectSet.Client.ListFromParams(poolPath, params)
	if err != nil {
		return nil, err
	}
	return buildPoolObjectSet(poolObjectSetResp), err
}
// generated function to build the appropriate response types
func buildPoolObjectSet(response interface{}) ([]*model.Pool) {
	values := reflect.ValueOf(response)
	results := make([]*model.Pool, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Pool{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
