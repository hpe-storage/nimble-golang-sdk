// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Token Service - Manage user's session information.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// TokenService type
type TokenService struct {
	objectSet *client.TokenObjectSet
}

// NewTokenService - method to initialize "TokenService"
func NewTokenService(gs *NsGroupService) *TokenService {
	objectSet := gs.client.GetTokenObjectSet()
	return &TokenService{objectSet: objectSet}
}

// GetTokens - method returns a array of pointers of type "Tokens"
func (svc *TokenService) GetTokens(params *param.GetParams) ([]*nimbleos.Token, error) {
	if params == nil {
		return nil, fmt.Errorf("GetTokens: invalid parameter specified, %v", params)
	}

	tokenResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}

// CreateToken - method creates a "Token"
func (svc *TokenService) CreateToken(obj *nimbleos.Token) (*nimbleos.Token, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateToken: invalid parameter specified, %v", obj)
	}

	tokenResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}

// UpdateToken - method modifies  the "Token"
func (svc *TokenService) UpdateToken(id string, obj *nimbleos.Token) (*nimbleos.Token, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateToken: invalid parameter specified, %v", obj)
	}

	tokenResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}

// GetTokenById - method returns a pointer to "Token"
func (svc *TokenService) GetTokenById(id string) (*nimbleos.Token, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetTokenById: invalid parameter specified, %v", id)
	}

	tokenResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}

// DeleteToken - deletes the "Token"
func (svc *TokenService) DeleteToken(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteToken: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}

// ReportUserDetailsToken - reports the user details for this token.
//   Required parameters:
//       id - ID for the session token.

func (svc *TokenService) ReportUserDetailsToken(id string) (*nimbleos.NsTokenReportUserDetailsReturn, error) {

	if len(id) == 0 {
		return nil, fmt.Errorf("ReportUserDetailsToken: invalid parameter specified id: %v ", id)
	}

	resp, err := svc.objectSet.ReportUserDetails(&id)
	return resp, err
}
