// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
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
func (objectSet *SoftwareVersionObjectSet) CreateObject(payload *nimbleos.SoftwareVersion) (*nimbleos.SoftwareVersion, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on SoftwareVersion")
}

// UpdateObject Modify existing SoftwareVersion object
func (objectSet *SoftwareVersionObjectSet) UpdateObject(id string, payload *nimbleos.SoftwareVersion) (*nimbleos.SoftwareVersion, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on SoftwareVersion")
}

// DeleteObject deletes the SoftwareVersion object with the specified ID
func (objectSet *SoftwareVersionObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on SoftwareVersion")
}

// GetObject returns a SoftwareVersion object with the given ID
func (objectSet *SoftwareVersionObjectSet) GetObject(id string) (*nimbleos.SoftwareVersion, error) {
	resp, err := objectSet.Client.Get(softwareVersionPath, id, &nimbleos.SoftwareVersion{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.SoftwareVersion), err
}

// GetObjectList returns the list of SoftwareVersion objects
func (objectSet *SoftwareVersionObjectSet) GetObjectList() ([]*nimbleos.SoftwareVersion, error) {
	resp, err := objectSet.Client.List(softwareVersionPath)
	if err != nil {
		return nil, err
	}
	return buildSoftwareVersionObjectSet(resp), err
}

// GetObjectListFromParams returns the list of SoftwareVersion objects using the given params query info
func (objectSet *SoftwareVersionObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.SoftwareVersion, error) {
	softwareVersionObjectSetResp, err := objectSet.Client.ListFromParams(softwareVersionPath, params)
	if err != nil {
		return nil, err
	}
	return buildSoftwareVersionObjectSet(softwareVersionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSoftwareVersionObjectSet(response interface{}) []*nimbleos.SoftwareVersion {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.SoftwareVersion, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.SoftwareVersion{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
