// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)


// Manage user's session information.
const (
    tokenPath = "tokens"
)

// TokenObjectSet
type TokenObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Token object
func (objectSet *TokenObjectSet) CreateObject(payload *model.Token) (*model.Token, error) {
	tokenObjectSetResp, err := objectSet.Client.Post(tokenPath, payload)
	if err !=nil {
		return nil,err
	}
	
	// null check
	if tokenObjectSetResp == nil {
		return nil,nil
	}
	return tokenObjectSetResp.(*model.Token), err
}

// UpdateObject Modify existing Token object
func (objectSet *TokenObjectSet) UpdateObject(id string, payload *model.Token) (*model.Token, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Token")
}

// DeleteObject deletes the Token object with the specified ID
func (objectSet *TokenObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(tokenPath, id)
	if err !=nil {
		return err
	}
	return nil
}

// GetObject returns a Token object with the given ID
func (objectSet *TokenObjectSet) GetObject(id string) (*model.Token, error) {
	tokenObjectSetResp, err := objectSet.Client.Get(tokenPath, id, model.Token{})
	if err != nil {
		return nil, err
	}
	
	// null check
	if tokenObjectSetResp == nil {
		return nil,nil
	}
	return tokenObjectSetResp.(*model.Token), err
}

// GetObjectList returns the list of Token objects
func (objectSet *TokenObjectSet) GetObjectList() ([]*model.Token, error) {
	tokenObjectSetResp, err := objectSet.Client.List(tokenPath)
	if err != nil {
		return nil, err
	}
	return buildTokenObjectSet(tokenObjectSetResp), err
}

// GetObjectListFromParams returns the list of Token objects using the given params query info
func (objectSet *TokenObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Token, error) {
	tokenObjectSetResp, err := objectSet.Client.ListFromParams(tokenPath, params)
	if err != nil {
		return nil, err
	}
	return buildTokenObjectSet(tokenObjectSetResp), err
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
