// Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *SpaceDomainObjectSet) CreateObject(payload *nimbleos.SpaceDomain) (*nimbleos.SpaceDomain, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on SpaceDomain")
}

// UpdateObject Modify existing SpaceDomain object
func (objectSet *SpaceDomainObjectSet) UpdateObject(id string, payload *nimbleos.SpaceDomain) (*nimbleos.SpaceDomain, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on SpaceDomain")
}

// DeleteObject deletes the SpaceDomain object with the specified ID
func (objectSet *SpaceDomainObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on SpaceDomain")
}

// GetObject returns a SpaceDomain object with the given ID
func (objectSet *SpaceDomainObjectSet) GetObject(id string) (*nimbleos.SpaceDomain, error) {
	resp, err := objectSet.Client.Get(spaceDomainPath, id, &nimbleos.SpaceDomain{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.SpaceDomain), err
}

// GetObjectList returns the list of SpaceDomain objects
func (objectSet *SpaceDomainObjectSet) GetObjectList() ([]*nimbleos.SpaceDomain, error) {
	resp, err := objectSet.Client.List(spaceDomainPath)
	if err != nil {
		return nil, err
	}
	return buildSpaceDomainObjectSet(resp), err
}

// GetObjectListFromParams returns the list of SpaceDomain objects using the given params query info
func (objectSet *SpaceDomainObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.SpaceDomain, error) {
	spaceDomainObjectSetResp, err := objectSet.Client.ListFromParams(spaceDomainPath, params)
	if err != nil {
		return nil, err
	}
	return buildSpaceDomainObjectSet(spaceDomainObjectSetResp), err
}

// generated function to build the appropriate response types
func buildSpaceDomainObjectSet(response interface{}) []*nimbleos.SpaceDomain {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.SpaceDomain, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.SpaceDomain{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
