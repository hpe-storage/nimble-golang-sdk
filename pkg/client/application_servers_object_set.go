// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

)

// An application server is an external agent that collaborates with an array to manage storage resources; for example, Volume Shadow Copy Service (VSS) or VMware.
const (
    applicationServerPath = "application_servers"
)

// ApplicationServerObjectSet
type ApplicationServerObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ApplicationServer object
func (objectSet *ApplicationServerObjectSet) CreateObject(payload *nimbleos.ApplicationServer) (*nimbleos.ApplicationServer, error) {
    resp, err := objectSet.Client.Post(applicationServerPath, payload, &nimbleos.ApplicationServer{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.ApplicationServer), err
}

// UpdateObject Modify existing ApplicationServer object
func (objectSet *ApplicationServerObjectSet) UpdateObject(id string, payload *nimbleos.ApplicationServer) (*nimbleos.ApplicationServer, error) {
    resp, err:= objectSet.Client.Put(applicationServerPath, id, payload, &nimbleos.ApplicationServer{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.ApplicationServer), err
}

// DeleteObject deletes the ApplicationServer object with the specified ID
func (objectSet *ApplicationServerObjectSet) DeleteObject(id string) error {
    err := objectSet.Client.Delete(applicationServerPath, id)
    if err != nil {
        return err
    }
    return nil
}

// GetObject returns a ApplicationServer object with the given ID
func (objectSet *ApplicationServerObjectSet) GetObject(id string) (*nimbleos.ApplicationServer, error) {
    resp, err:= objectSet.Client.Get(applicationServerPath, id, &nimbleos.ApplicationServer{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.ApplicationServer), err
}

// GetObjectList returns the list of ApplicationServer objects
func (objectSet *ApplicationServerObjectSet) GetObjectList() ([]*nimbleos.ApplicationServer, error) {
    resp, err:= objectSet.Client.List(applicationServerPath)
    if err != nil {
        return nil, err
    }
    return buildApplicationServerObjectSet(resp), err
}

// GetObjectListFromParams returns the list of ApplicationServer objects using the given params query info
func (objectSet *ApplicationServerObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ApplicationServer, error) {
    applicationServerObjectSetResp,err:= objectSet.Client.ListFromParams(applicationServerPath, params)
    if err != nil {
        return nil, err
    }
    return buildApplicationServerObjectSet(applicationServerObjectSetResp), err
}
// generated function to build the appropriate response types
func buildApplicationServerObjectSet(response interface{}) ([]*nimbleos.ApplicationServer) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.ApplicationServer, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.ApplicationServer{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

