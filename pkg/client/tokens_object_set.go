// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
	"strings"
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
	resp, err := objectSet.Client.Post(tokenPath, payload, &nimbleos.Token{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)
			}
		}
		return nil, err
	}

	return resp.(*nimbleos.Token), err
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
	resp, err := objectSet.Client.Get(tokenPath, id, &nimbleos.Token{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Token), err
}

// GetObjectList returns the list of Token objects
func (objectSet *TokenObjectSet) GetObjectList() ([]*nimbleos.Token, error) {
	resp, err := objectSet.Client.List(tokenPath)
	if err != nil {
		return nil, err
	}
	return buildTokenObjectSet(resp), err
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

// List of supported actions on object sets

// ReportUserDetails - Reports the user details for this token.
func (objectSet *TokenObjectSet) ReportUserDetails(id *string) (*nimbleos.NsTokenReportUserDetailsReturn, error) {

	reportUserDetailsUri := tokenPath
	reportUserDetailsUri = reportUserDetailsUri + "/" + *id
	reportUserDetailsUri = reportUserDetailsUri + "/actions/" + "report_user_details"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	resp, err := objectSet.Client.Post(reportUserDetailsUri, payload, &nimbleos.NsTokenReportUserDetailsReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsTokenReportUserDetailsReturn), err
}
