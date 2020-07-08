// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
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
	return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelInitiatorAlias")
}

// UpdateObject Modify existing FibreChannelInitiatorAlias object
func (objectSet *FibreChannelInitiatorAliasObjectSet) UpdateObject(id string, payload *model.FibreChannelInitiatorAlias) (*model.FibreChannelInitiatorAlias, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on FibreChannelInitiatorAlias")
}

// DeleteObject deletes the FibreChannelInitiatorAlias object with the specified ID
func (objectSet *FibreChannelInitiatorAliasObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on FibreChannelInitiatorAlias")
}

// GetObject returns a FibreChannelInitiatorAlias object with the given ID
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObject(id string) (*model.FibreChannelInitiatorAlias, error) {
	fibreChannelInitiatorAliasObjectSetResp, err := objectSet.Client.Get(fibreChannelInitiatorAliasPath, id, model.FibreChannelInitiatorAlias{})
	if err != nil {
		return nil, err
	}

	// null check
	if fibreChannelInitiatorAliasObjectSetResp == nil {
		return nil, nil
	}
	return fibreChannelInitiatorAliasObjectSetResp.(*model.FibreChannelInitiatorAlias), err
}

// GetObjectList returns the list of FibreChannelInitiatorAlias objects
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObjectList() ([]*model.FibreChannelInitiatorAlias, error) {
	fibreChannelInitiatorAliasObjectSetResp, err := objectSet.Client.List(fibreChannelInitiatorAliasPath)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelInitiatorAliasObjectSet(fibreChannelInitiatorAliasObjectSetResp), err
}

// GetObjectListFromParams returns the list of FibreChannelInitiatorAlias objects using the given params query info
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.FibreChannelInitiatorAlias, error) {
	fibreChannelInitiatorAliasObjectSetResp, err := objectSet.Client.ListFromParams(fibreChannelInitiatorAliasPath, params)
	if err != nil {
		return nil, err
	}
	return buildFibreChannelInitiatorAliasObjectSet(fibreChannelInitiatorAliasObjectSetResp), err
}

// generated function to build the appropriate response types
func buildFibreChannelInitiatorAliasObjectSet(response interface{}) []*model.FibreChannelInitiatorAlias {
	values := reflect.ValueOf(response)
	results := make([]*model.FibreChannelInitiatorAlias, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.FibreChannelInitiatorAlias{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
