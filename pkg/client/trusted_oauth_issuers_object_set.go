// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Oauth Credential Issuers that this device will trust.
const (
	trustedOauthIssuerPath = "trusted_oauth_issuers"
)

// TrustedOauthIssuerObjectSet
type TrustedOauthIssuerObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new TrustedOauthIssuer object
func (objectSet *TrustedOauthIssuerObjectSet) CreateObject(payload *nimbleos.TrustedOauthIssuer) (*nimbleos.TrustedOauthIssuer, error) {
	resp, err := objectSet.Client.Post(trustedOauthIssuerPath, payload, &nimbleos.TrustedOauthIssuer{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.TrustedOauthIssuer), err
}

// UpdateObject Modify existing TrustedOauthIssuer object
func (objectSet *TrustedOauthIssuerObjectSet) UpdateObject(id string, payload *nimbleos.TrustedOauthIssuer) (*nimbleos.TrustedOauthIssuer, error) {
	return nil, fmt.Errorf("Unsupported operation 'update' on TrustedOauthIssuer")
}

// DeleteObject deletes the TrustedOauthIssuer object with the specified ID
func (objectSet *TrustedOauthIssuerObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on TrustedOauthIssuer")
}

// GetObject returns a TrustedOauthIssuer object with the given ID
func (objectSet *TrustedOauthIssuerObjectSet) GetObject(id string) (*nimbleos.TrustedOauthIssuer, error) {
	resp, err := objectSet.Client.Get(trustedOauthIssuerPath, id, &nimbleos.TrustedOauthIssuer{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.TrustedOauthIssuer), err
}

// GetObjectList returns the list of TrustedOauthIssuer objects
func (objectSet *TrustedOauthIssuerObjectSet) GetObjectList() ([]*nimbleos.TrustedOauthIssuer, error) {
	resp, err := objectSet.Client.List(trustedOauthIssuerPath)
	if err != nil {
		return nil, err
	}
	return buildTrustedOauthIssuerObjectSet(resp), err
}

// GetObjectListFromParams returns the list of TrustedOauthIssuer objects using the given params query info
func (objectSet *TrustedOauthIssuerObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.TrustedOauthIssuer, error) {
	trustedOauthIssuerObjectSetResp, err := objectSet.Client.ListFromParams(trustedOauthIssuerPath, params)
	if err != nil {
		return nil, err
	}
	return buildTrustedOauthIssuerObjectSet(trustedOauthIssuerObjectSetResp), err
}

// generated function to build the appropriate response types
func buildTrustedOauthIssuerObjectSet(response interface{}) []*nimbleos.TrustedOauthIssuer {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.TrustedOauthIssuer, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.TrustedOauthIssuer{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
