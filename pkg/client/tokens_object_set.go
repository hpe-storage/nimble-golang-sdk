// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
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
func (objectSet *TokenObjectSet) CreateObject(payload *nimbleos.Token) (*nimbleos.Token, error) {
	newPayload, err := nimbleos.EncodeToken(payload)
	resp, err := objectSet.Client.Post(tokenPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return nimbleos.DecodeToken(resp)
}

// UpdateObject Modify existing Token object
func (objectSet *TokenObjectSet) UpdateObject(id string, payload *nimbleos.Token) (*nimbleos.Token, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on Token")
}

// DeleteObject deletes the Token object with the specified ID
func (objectSet *TokenObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(tokenPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Token object with the given ID
func (objectSet *TokenObjectSet) GetObject(id string) (*nimbleos.Token, error) {
	tokenObjectSetResp, err := objectSet.Client.Get(tokenPath, id, nimbleos.Token{})
	if err != nil {
		return nil, err
	}

	// null check
	if tokenObjectSetResp == nil {
		return nil, nil
	}
	return tokenObjectSetResp.(*nimbleos.Token), err
}

// GetObjectList returns the list of Token objects
func (objectSet *TokenObjectSet) GetObjectList() ([]*nimbleos.Token, error) {
	tokenObjectSetResp, err := objectSet.Client.List(tokenPath)
	if err != nil {
		return nil, err
	}
	return buildTokenObjectSet(tokenObjectSetResp), err
}

// GetObjectListFromParams returns the list of Token objects using the given params query info
func (objectSet *TokenObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Token, error) {
	tokenObjectSetResp, err := objectSet.Client.ListFromParams(tokenPath, params)
	if err != nil {
		return nil, err
	}
	return buildTokenObjectSet(tokenObjectSetResp), err
}

// generated function to build the appropriate response types
func buildTokenObjectSet(response interface{}) []*nimbleos.Token {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Token, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Token{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
