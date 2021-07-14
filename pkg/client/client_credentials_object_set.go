// Copyright 2020 Hewlett Packard Enterprise Development LP


package client

import (

        "github.com/hpe-storage/common-host-libs/jsonutil"
        "reflect"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"

)

// Credential that this device will trust.
const (
    clientCredentialPath = "client_credentials"
)

// ClientCredentialObjectSet
type ClientCredentialObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new ClientCredential object
func (objectSet *ClientCredentialObjectSet) CreateObject(payload *nimbleos.ClientCredential) (*nimbleos.ClientCredential, error) {
    resp, err := objectSet.Client.Post(clientCredentialPath, payload, &nimbleos.ClientCredential{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.ClientCredential), err
}

// UpdateObject Modify existing ClientCredential object
func (objectSet *ClientCredentialObjectSet) UpdateObject(id string, payload *nimbleos.ClientCredential) (*nimbleos.ClientCredential, error) {
    resp, err:= objectSet.Client.Put(clientCredentialPath, id, payload, &nimbleos.ClientCredential{})
    if err != nil {
        return nil, err
    }

    return resp.(*nimbleos.ClientCredential), err
}

// DeleteObject deletes the ClientCredential object with the specified ID
func (objectSet *ClientCredentialObjectSet) DeleteObject(id string) error {
    err := objectSet.Client.Delete(clientCredentialPath, id)
    if err != nil {
        return err
    }
    return nil
}

// GetObject returns a ClientCredential object with the given ID
func (objectSet *ClientCredentialObjectSet) GetObject(id string) (*nimbleos.ClientCredential, error) {
    resp, err:= objectSet.Client.Get(clientCredentialPath, id, &nimbleos.ClientCredential{})
    if err != nil {
        return nil, err
    }

    // null check
    if resp == nil {
        return nil, nil
    }
    return resp.(*nimbleos.ClientCredential), err
}

// GetObjectList returns the list of ClientCredential objects
func (objectSet *ClientCredentialObjectSet) GetObjectList() ([]*nimbleos.ClientCredential, error) {
    resp, err:= objectSet.Client.List(clientCredentialPath)
    if err != nil {
        return nil, err
    }
    return buildClientCredentialObjectSet(resp), err
}

// GetObjectListFromParams returns the list of ClientCredential objects using the given params query info
func (objectSet *ClientCredentialObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ClientCredential, error) {
    clientCredentialObjectSetResp,err:= objectSet.Client.ListFromParams(clientCredentialPath, params)
    if err != nil {
        return nil, err
    }
    return buildClientCredentialObjectSet(clientCredentialObjectSetResp), err
}
// generated function to build the appropriate response types
func buildClientCredentialObjectSet(response interface{}) ([]*nimbleos.ClientCredential) {
    values := reflect.ValueOf(response)
    results := make([]*nimbleos.ClientCredential, values.Len())

    for i := 0; i < values.Len(); i++ {
        value := &nimbleos.ClientCredential{}
        jsonutil.Decode(values.Index(i).Interface(), value)
        results[i] = value
    }

    return results
}

