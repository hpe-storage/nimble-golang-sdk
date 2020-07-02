// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Token Service - Manage user's session information.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// TokenService type 
type TokenService struct {
	objectSet *client.TokenObjectSet
}

// NewTokenService - method to initialize "TokenService" 
func NewTokenService(gs *NsGroupService) (*TokenService) {
	objectSet := gs.client.GetTokenObjectSet()
	return &TokenService{objectSet: objectSet}
}

// GetTokens - method returns a array of pointers of type "Tokens"
func (svc *TokenService) GetTokens(params *util.GetParams) ([]*model.Token, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateToken - method creates a "Token"
func (svc *TokenService) CreateToken(obj *model.Token) (*model.Token, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateToken - method modifies  the "Token" 
func (svc *TokenService) UpdateToken(id string, obj *model.Token) (*model.Token, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetTokenById - method returns a pointer to "Token"
func (svc *TokenService) GetTokenById(id string) (*model.Token, error) {
	return svc.objectSet.GetObject(id)
}


// DeleteToken - deletes the "Token"
func (svc *TokenService) DeleteToken(id string) error {
	return svc.objectSet.DeleteObject(id)
}
