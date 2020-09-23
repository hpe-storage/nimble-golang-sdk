// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Show the API version.
const (
	versionPath = "versions"
)

// VersionObjectSet
type VersionObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Version object
func (objectSet *VersionObjectSet) CreateObject(payload *nimbleos.Version) (*nimbleos.Version, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Version")
}

// UpdateObject Modify existing Version object
func (objectSet *VersionObjectSet) UpdateObject(id string, payload *nimbleos.Version) (*nimbleos.Version, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Version")
}

// DeleteObject deletes the Version object with the specified ID
func (objectSet *VersionObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Version")
}

// GetObject returns a Version object with the given ID
func (objectSet *VersionObjectSet) GetObject(id string) (*nimbleos.Version, error) {
	resp, err := objectSet.Client.Get(versionPath, id, &nimbleos.Version{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Version), err
}

// GetObjectList returns the list of Version objects
func (objectSet *VersionObjectSet) GetObjectList() ([]*nimbleos.Version, error) {
	resp, err := objectSet.Client.List(versionPath)
	if err != nil {
		return nil, err
	}
	return buildVersionObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Version objects using the given params query info
func (objectSet *VersionObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Version, error) {
	versionObjectSetResp, err := objectSet.Client.ListFromParams(versionPath, params)
	if err != nil {
		return nil, err
	}
	return buildVersionObjectSet(versionObjectSetResp), err
}

// generated function to build the appropriate response types
func buildVersionObjectSet(response interface{}) []*nimbleos.Version {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Version, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Version{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
