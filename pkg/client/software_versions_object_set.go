// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
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
	return nil, fmt.Errorf("Unsupported operation 'create' on SoftwareVersion")
}

// UpdateObject Modify existing SoftwareVersion object
func (objectSet *SoftwareVersionObjectSet) UpdateObject(id string, payload *model.SoftwareVersion) (*model.SoftwareVersion, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on SoftwareVersion")
}

// DeleteObject deletes the SoftwareVersion object with the specified ID
func (objectSet *SoftwareVersionObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on SoftwareVersion")
}

// GetObject returns a SoftwareVersion object with the given ID
func (objectSet *SoftwareVersionObjectSet) GetObject(id string) (*model.SoftwareVersion, error) {
	softwareVersionObjectSetResp, err := objectSet.Client.Get(softwareVersionPath, id, model.SoftwareVersion{})
	if err != nil {
		return nil, err
	}

	// null check
	if softwareVersionObjectSetResp == nil {
		return nil, nil
	}
	return softwareVersionObjectSetResp.(*model.SoftwareVersion), err
}

// GetObjectList returns the list of SoftwareVersion objects
func (objectSet *SoftwareVersionObjectSet) GetObjectList() ([]*model.SoftwareVersion, error) {
	softwareVersionObjectSetResp, err := objectSet.Client.List(softwareVersionPath)
	if err != nil {
		return nil, err
	}
	return buildSoftwareVersionObjectSet(softwareVersionObjectSetResp), err
}

// GetObjectListFromParams returns the list of SoftwareVersion objects using the given params query info
func (objectSet *SoftwareVersionObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.SoftwareVersion, error) {
	softwareVersionObjectSetResp, err := objectSet.Client.ListFromParams(softwareVersionPath, params)
	if err != nil {
		return nil, err
	}
	return buildSoftwareVersionObjectSet(softwareVersionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSoftwareVersionObjectSet(response interface{}) []*model.SoftwareVersion {
	values := reflect.ValueOf(response)
	results := make([]*model.SoftwareVersion, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.SoftwareVersion{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
