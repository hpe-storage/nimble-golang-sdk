// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Token Service - Manage user's session information.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
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
func (svc *TokenService) GetTokens(params *util.GetParams) ([]*model.Token, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	tokenResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}

// CreateToken - method creates a "Token"
func (svc *TokenService) CreateToken(obj *model.Token) (*model.Token, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	tokenResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}

// UpdateToken - method modifies  the "Token"
func (svc *TokenService) UpdateToken(id string, obj *model.Token) (*model.Token, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	tokenResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}

// GetTokenById - method returns a pointer to "Token"
func (svc *TokenService) GetTokenById(id string) (*model.Token, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
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
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
