// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// This API provides the alias information for Fibre Channel initiators.
const (
    fibreChannelInitiatorAliasPath = "fibre_channel_initiator_aliases"
)

// FibreChannelInitiatorAliasObjectSet
type FibreChannelInitiatorAliasObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new FibreChannelInitiatorAlias object
func (objectSet *FibreChannelInitiatorAliasObjectSet) CreateObject(payload *model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	response, err := objectSet.Client.Post(fibreChannelInitiatorAliasPath, payload)
	return response.(*model.FibreChannelInitiatorAlias), err
}

// UpdateObject Modify existing FibreChannelInitiatorAlias object
func (objectSet *FibreChannelInitiatorAliasObjectSet) UpdateObject(id string, payload *model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	response, err := objectSet.Client.Put(fibreChannelInitiatorAliasPath, id, payload)
	return response.(*model.FibreChannelInitiatorAlias), err
}

// DeleteObject deletes the FibreChannelInitiatorAlias object with the specified ID
func (objectSet *FibreChannelInitiatorAliasObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(fibreChannelInitiatorAliasPath, id)
}

// GetObject returns a FibreChannelInitiatorAlias object with the given ID
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObject(id string) (*model.FibreChannelInitiatorAlias, error) {
	response, err := objectSet.Client.Get(fibreChannelInitiatorAliasPath, id, model.FibreChannelInitiatorAlias{})
	if response == nil {
		return nil, err
	}
	return response.(*model.FibreChannelInitiatorAlias), err
}

// GetObjectList returns the list of FibreChannelInitiatorAlias objects
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObjectList() ([]*model.FibreChannelInitiatorAlias, error) {
	response, err := objectSet.Client.List(fibreChannelInitiatorAliasPath)
	return buildFibreChannelInitiatorAliasObjectSet(response), err
}

// GetObjectListFromParams returns the list of FibreChannelInitiatorAlias objects using the given params query info
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelInitiatorAlias, error) {
	response, err := objectSet.Client.ListFromParams(fibreChannelInitiatorAliasPath, params)
	return buildFibreChannelInitiatorAliasObjectSet(response), err
}

// generated function to build the appropriate response types
func buildFibreChannelInitiatorAliasObjectSet(response interface{}) ([]*model.FibreChannelInitiatorAlias) {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelInitiatorAlias, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelInitiatorAlias{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}