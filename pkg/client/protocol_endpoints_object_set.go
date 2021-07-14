// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

    "fmt"
        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

)

// Protocol endpoints are administrative logical units (LUs) in an LU conglomerate to be used with VMware Virtual Volumes.
const (
    protocolEndpointPath = "protocol_endpoints"
)

// ProtocolEndpointObjectSet
type ProtocolEndpointObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ProtocolEndpoint object
func (objectSet *ProtocolEndpointObjectSet) CreateObject(payload *nimbleos.ProtocolEndpoint) (*nimbleos.ProtocolEndpoint, error) {
    return nil, fmt.Errorf("Unsupported operation 'create' on ProtocolEndpoint")
}

// UpdateObject Modify existing ProtocolEndpoint object
func (objectSet *ProtocolEndpointObjectSet) UpdateObject(id string, payload *nimbleos.ProtocolEndpoint) (*nimbleos.ProtocolEndpoint, error) {
    return nil, fmt.Errorf("Unsupported operation 'update' on ProtocolEndpoint")
}

// DeleteObject deletes the ProtocolEndpoint object with the specified ID
func (objectSet *ProtocolEndpointObjectSet) DeleteObject(id string) error {
    return fmt.Errorf("Unsupported operation 'delete' on ProtocolEndpoint")
}

// GetObject returns a ProtocolEndpoint object with the given ID
func (objectSet *ProtocolEndpointObjectSet) GetObject(id string) (*nimbleos.ProtocolEndpoint, error) {
    resp, err:= objectSet.Client.Get(protocolEndpointPath, id, &nimbleos.ProtocolEndpoint{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.ProtocolEndpoint), err
}

// GetObjectList returns the list of ProtocolEndpoint objects
func (objectSet *ProtocolEndpointObjectSet) GetObjectList() ([]*nimbleos.ProtocolEndpoint, error) {
    resp, err:= objectSet.Client.List(protocolEndpointPath)
    if err != nil {
        return nil, err
    }
    return buildProtocolEndpointObjectSet(resp), err
}

// GetObjectListFromParams returns the list of ProtocolEndpoint objects using the given params query info
func (objectSet *ProtocolEndpointObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ProtocolEndpoint, error) {
    protocolEndpointObjectSetResp,err:= objectSet.Client.ListFromParams(protocolEndpointPath, params)
    if err != nil {
        return nil, err
    }
    return buildProtocolEndpointObjectSet(protocolEndpointObjectSetResp), err
}
// generated function to build the appropriate response types
func buildProtocolEndpointObjectSet(response interface{}) ([]*nimbleos.ProtocolEndpoint) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.ProtocolEndpoint, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.ProtocolEndpoint{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

