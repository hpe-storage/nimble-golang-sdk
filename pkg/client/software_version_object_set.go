// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Show the software version.
const (
    softwareVersionPath = "software_versions"
)

// SoftwareVersionObjectSet
type SoftwareVersionObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new SoftwareVersion object
func (objectSet *SoftwareVersionObjectSet) CreateObject(payload *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	response, err := objectSet.Client.Post(softwareVersionPath, payload)
	return response.(*model.SoftwareVersion), err
}

// UpdateObject Modify existing SoftwareVersion object
func (objectSet *SoftwareVersionObjectSet) UpdateObject(id string, payload *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	response, err := objectSet.Client.Put(softwareVersionPath, id, payload)
	return response.(*model.SoftwareVersion), err
}

// DeleteObject deletes the SoftwareVersion object with the specified ID
func (objectSet *SoftwareVersionObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(softwareVersionPath, id)
}

// GetObject returns a SoftwareVersion object with the given ID
func (objectSet *SoftwareVersionObjectSet) GetObject(id string) (*model.SoftwareVersion, error) {
	response, err := objectSet.Client.Get(softwareVersionPath, id, model.SoftwareVersion{})
	if response == nil {
		return nil, err
	}
	return response.(*model.SoftwareVersion), err
}

// GetObjectList returns the list of SoftwareVersion objects
func (objectSet *SoftwareVersionObjectSet) GetObjectList() ([]*model.SoftwareVersion, error) {
	response, err := objectSet.Client.List(softwareVersionPath)
	return buildSoftwareVersionObjectSet(response), err
}

// GetObjectListFromParams returns the list of SoftwareVersion objects using the given params query info
func (objectSet *SoftwareVersionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.SoftwareVersion, error) {
	response, err := objectSet.Client.ListFromParams(softwareVersionPath, params)
	return buildSoftwareVersionObjectSet(response), err
}

// generated function to build the appropriate response types
func buildSoftwareVersionObjectSet(response interface{}) ([]*model.SoftwareVersion) {
	values := reflect.ValueOf(response)
	results := make([]*model.SoftwareVersion, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.SoftwareVersion{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}