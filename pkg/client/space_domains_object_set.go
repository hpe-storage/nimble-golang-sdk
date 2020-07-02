// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// A space domain is created for each application category and block size for each each pool.
const (
    spaceDomainPath = "space_domains"
)

// SpaceDomainObjectSet
type SpaceDomainObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new SpaceDomain object
func (objectSet *SpaceDomainObjectSet) CreateObject(payload *model.SpaceDomain) (*model.SpaceDomain, error) {
	response, err := objectSet.Client.Post(spaceDomainPath, payload)
	return response.(*model.SpaceDomain), err
}

// UpdateObject Modify existing SpaceDomain object
func (objectSet *SpaceDomainObjectSet) UpdateObject(id string, payload *model.SpaceDomain) (*model.SpaceDomain, error) {
	response, err := objectSet.Client.Put(spaceDomainPath, id, payload)
	return response.(*model.SpaceDomain), err
}

// DeleteObject deletes the SpaceDomain object with the specified ID
func (objectSet *SpaceDomainObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(spaceDomainPath, id)
}

// GetObject returns a SpaceDomain object with the given ID
func (objectSet *SpaceDomainObjectSet) GetObject(id string) (*model.SpaceDomain, error) {
	response, err := objectSet.Client.Get(spaceDomainPath, id, model.SpaceDomain{})
	if response == nil {
		return nil, err
	}
	return response.(*model.SpaceDomain), err
}

// GetObjectList returns the list of SpaceDomain objects
func (objectSet *SpaceDomainObjectSet) GetObjectList() ([]*model.SpaceDomain, error) {
	response, err := objectSet.Client.List(spaceDomainPath)
	return buildSpaceDomainObjectSet(response), err
}

// GetObjectListFromParams returns the list of SpaceDomain objects using the given params query info
func (objectSet *SpaceDomainObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.SpaceDomain, error) {
	response, err := objectSet.Client.ListFromParams(spaceDomainPath, params)
	return buildSpaceDomainObjectSet(response), err
}

// generated function to build the appropriate response types
func buildSpaceDomainObjectSet(response interface{}) ([]*model.SpaceDomain) {
	values := reflect.ValueOf(response)
	results := make([]*model.SpaceDomain, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.SpaceDomain{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}