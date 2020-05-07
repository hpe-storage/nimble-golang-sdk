// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

import (
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
)

type TokenService struct {
	objectSet *client.TokenObjectSet
}

func NewTokenService(gs *GroupService) (vs *TokenService) {
	objectSet := gs.client.GetTokenObjectSet()
	return &TokenService{objectSet: objectSet}
}
