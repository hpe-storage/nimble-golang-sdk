// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

    "fmt"
        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

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
func (objectSet *FibreChannelInitiatorAliasObjectSet) CreateObject(payload *nimbleos.FibreChannelInitiatorAlias) (*nimbleos.FibreChannelInitiatorAlias, error) {
    return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelInitiatorAlias")
}

// UpdateObject Modify existing FibreChannelInitiatorAlias object
func (objectSet *FibreChannelInitiatorAliasObjectSet) UpdateObject(id string, payload *nimbleos.FibreChannelInitiatorAlias) (*nimbleos.FibreChannelInitiatorAlias, error) {
    return nil, fmt.Errorf("Unsupported operation 'update' on FibreChannelInitiatorAlias")
}

// DeleteObject deletes the FibreChannelInitiatorAlias object with the specified ID
func (objectSet *FibreChannelInitiatorAliasObjectSet) DeleteObject(id string) error {
    return fmt.Errorf("Unsupported operation 'delete' on FibreChannelInitiatorAlias")
}

// GetObject returns a FibreChannelInitiatorAlias object with the given ID
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObject(id string) (*nimbleos.FibreChannelInitiatorAlias, error) {
    resp, err:= objectSet.Client.Get(fibreChannelInitiatorAliasPath, id, &nimbleos.FibreChannelInitiatorAlias{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.FibreChannelInitiatorAlias), err
}

// GetObjectList returns the list of FibreChannelInitiatorAlias objects
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObjectList() ([]*nimbleos.FibreChannelInitiatorAlias, error) {
    resp, err:= objectSet.Client.List(fibreChannelInitiatorAliasPath)
    if err != nil {
        return nil, err
    }
    return buildFibreChannelInitiatorAliasObjectSet(resp), err
}

// GetObjectListFromParams returns the list of FibreChannelInitiatorAlias objects using the given params query info
func (objectSet *FibreChannelInitiatorAliasObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.FibreChannelInitiatorAlias, error) {
    fibreChannelInitiatorAliasObjectSetResp,err:= objectSet.Client.ListFromParams(fibreChannelInitiatorAliasPath, params)
    if err != nil {
        return nil, err
    }
    return buildFibreChannelInitiatorAliasObjectSet(fibreChannelInitiatorAliasObjectSetResp), err
}
// generated function to build the appropriate response types
func buildFibreChannelInitiatorAliasObjectSet(response interface{}) ([]*nimbleos.FibreChannelInitiatorAlias) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.FibreChannelInitiatorAlias, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.FibreChannelInitiatorAlias{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

