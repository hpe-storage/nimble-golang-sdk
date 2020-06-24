// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

/**
 * Fibre Channel ports provide data access. This API provides the list of all Fibre Channel ports configured on the arrays.
 *
 */
const (
    fibreChannelPortPath = "fibre_channel_ports"
)

/**
 * FibreChannelPortObjectSet
*/
type FibreChannelPortObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new FibreChannelPort object
func (objectSet *FibreChannelPortObjectSet) CreateObject(payload *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	response, err := objectSet.Client.Post(fibreChannelPortPath, payload)
	return response.(*model.FibreChannelPort), err
}

// UpdateObject Modify existing FibreChannelPort object
func (objectSet *FibreChannelPortObjectSet) UpdateObject(id string, payload *model.FibreChannelPort) (*model.FibreChannelPort, error) {
	response, err := objectSet.Client.Put(fibreChannelPortPath, id, payload)
	return response.(*model.FibreChannelPort), err
}

// DeleteObject deletes the FibreChannelPort object with the specified ID
func (objectSet *FibreChannelPortObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(fibreChannelPortPath, id)
}

// GetObject returns a FibreChannelPort object with the given ID
func (objectSet *FibreChannelPortObjectSet) GetObject(id string) (*model.FibreChannelPort, error) {
	response, err := objectSet.Client.Get(fibreChannelPortPath, id, model.FibreChannelPort{})
	if response == nil {
		return nil, err
	}
	return response.(*model.FibreChannelPort), err
}

// GetObjectList returns the list of FibreChannelPort objects
func (objectSet *FibreChannelPortObjectSet) GetObjectList() ([]*model.FibreChannelPort, error) {
	response, err := objectSet.Client.List(fibreChannelPortPath)
	return buildFibreChannelPortObjectSet(response), err
}

// GetObjectListFromParams returns the list of FibreChannelPort objects using the given params query info
func (objectSet *FibreChannelPortObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelPort, error) {
	response, err := objectSet.Client.ListFromParams(fibreChannelPortPath, params)
	return buildFibreChannelPortObjectSet(response), err
}

// generated function to build the appropriate response types
func buildFibreChannelPortObjectSet(response interface{}) ([]*model.FibreChannelPort) {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelPort, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelPort{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}