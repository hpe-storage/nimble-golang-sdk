// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Token Service - Manage user's session information.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
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

// GetTokensWithFields - method returns a array of pointers of type "Token" 
func (svc *TokenService) GetTokensWithFields(fields []string) ([]*model.Token, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateToken - method creates a "Token"
func (svc *TokenService) CreateToken(obj *model.Token) (*model.Token, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditToken - method modifies  the "Token" 
func (svc *TokenService) EditToken(id string, obj *model.Token) (*model.Token, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyToken - private method for more than one element check. 
func onlyToken(objs []*model.Token) (*model.Token, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Token found with the given filter")
	}

	return objs[0], nil
}

 
// GetTokensByID - method returns associative a array of pointers of type "Token", filter by Id
func (svc *TokenService) GetTokensByID(pool *model.Pool, fields []string) (map[string]*model.Token, error) {
	params := &util.GetParams{}

	// make sure ID field is selected
	if _, found := params.FindField("id"); !found {
		fields = append(fields, "id")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetTokens(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Token)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetTokenById - method returns a pointer to "Token"
func (svc *TokenService) GetTokenById(id string) (*model.Token, error) {
	return svc.objectSet.GetObject(id)
}


