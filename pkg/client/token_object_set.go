// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Manage user&#39;s session information.
const (
    tokenPath = "tokens"
)


// TokenObjectSet
type TokenObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Token object
func (objectSet *TokenObjectSet) CreateObject(payload *model.Token) (*model.Token, error) {
	response, err := objectSet.Client.Post(tokenPath, payload)
	return response.(*model.Token), err
}

// UpdateObject Modify existing Token object
func (objectSet *TokenObjectSet) UpdateObject(id string, payload *model.Token) (*model.Token, error) {
	response, err := objectSet.Client.Put(tokenPath, id, payload)
	return response.(*model.Token), err
}

// DeleteObject deletes the Token object with the specified ID
func (objectSet *TokenObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(tokenPath, id)
}

// GetObject returns a Token object with the given ID
func (objectSet *TokenObjectSet) GetObject(id string) (*model.Token, error) {
	response, err := objectSet.Client.Get(tokenPath, id, model.Token{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Token), err
}

// GetObjectList returns the list of Token objects
func (objectSet *TokenObjectSet) GetObjectList() ([]*model.Token, error) {
	response, err := objectSet.Client.List(tokenPath)
	return buildTokenObjectSet(response), err
}

// GetObjectListFromParams returns the list of Token objects using the given params query info
func (objectSet *TokenObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Token, error) {
	response, err := objectSet.Client.ListFromParams(tokenPath, params)
	return buildTokenObjectSet(response), err
}

// generated function to build the appropriate response types
func buildTokenObjectSet(response interface{}) ([]*model.Token) {
	values := reflect.ValueOf(response)
	results := make([]*model.Token, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Token{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}