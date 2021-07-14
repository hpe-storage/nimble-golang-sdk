// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

    "fmt"
        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

)

// Fibre Channel session is created when Fibre Channel initiator connects to this group.
const (
    fibreChannelSessionPath = "fibre_channel_sessions"
)

// FibreChannelSessionObjectSet
type FibreChannelSessionObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new FibreChannelSession object
func (objectSet *FibreChannelSessionObjectSet) CreateObject(payload *nimbleos.FibreChannelSession) (*nimbleos.FibreChannelSession, error) {
    return nil, fmt.Errorf("Unsupported operation 'create' on FibreChannelSession")
}

// UpdateObject Modify existing FibreChannelSession object
func (objectSet *FibreChannelSessionObjectSet) UpdateObject(id string, payload *nimbleos.FibreChannelSession) (*nimbleos.FibreChannelSession, error) {
    return nil, fmt.Errorf("Unsupported operation 'update' on FibreChannelSession")
}

// DeleteObject deletes the FibreChannelSession object with the specified ID
func (objectSet *FibreChannelSessionObjectSet) DeleteObject(id string) error {
    return fmt.Errorf("Unsupported operation 'delete' on FibreChannelSession")
}

// GetObject returns a FibreChannelSession object with the given ID
func (objectSet *FibreChannelSessionObjectSet) GetObject(id string) (*nimbleos.FibreChannelSession, error) {
    resp, err:= objectSet.Client.Get(fibreChannelSessionPath, id, &nimbleos.FibreChannelSession{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.FibreChannelSession), err
}

// GetObjectList returns the list of FibreChannelSession objects
func (objectSet *FibreChannelSessionObjectSet) GetObjectList() ([]*nimbleos.FibreChannelSession, error) {
    resp, err:= objectSet.Client.List(fibreChannelSessionPath)
    if err != nil {
        return nil, err
    }
    return buildFibreChannelSessionObjectSet(resp), err
}

// GetObjectListFromParams returns the list of FibreChannelSession objects using the given params query info
func (objectSet *FibreChannelSessionObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.FibreChannelSession, error) {
    fibreChannelSessionObjectSetResp,err:= objectSet.Client.ListFromParams(fibreChannelSessionPath, params)
    if err != nil {
        return nil, err
    }
    return buildFibreChannelSessionObjectSet(fibreChannelSessionObjectSetResp), err
}
// generated function to build the appropriate response types
func buildFibreChannelSessionObjectSet(response interface{}) ([]*nimbleos.FibreChannelSession) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.FibreChannelSession, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.FibreChannelSession{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

