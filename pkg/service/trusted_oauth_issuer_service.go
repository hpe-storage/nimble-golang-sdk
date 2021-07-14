// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// TrustedOauthIssuer Service - Oauth Credential Issuers that this device will trust.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// TrustedOauthIssuerService type
type TrustedOauthIssuerService struct {
    objectSet *client.TrustedOauthIssuerObjectSet
}

// NewTrustedOauthIssuerService - method to initialize "TrustedOauthIssuerService"
func NewTrustedOauthIssuerService(gs *NsGroupService) (*TrustedOauthIssuerService) {
    objectSet := gs.client.GetTrustedOauthIssuerObjectSet()
    return &TrustedOauthIssuerService{objectSet: objectSet}
}

// GetTrustedOauthIssuers - method returns a array of pointers of type "TrustedOauthIssuers"
func (svc *TrustedOauthIssuerService) GetTrustedOauthIssuers(params *param.GetParams) ([]*nimbleos.TrustedOauthIssuer, error) {
    trustedOauthIssuerResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return trustedOauthIssuerResp,nil
}

// CreateTrustedOauthIssuer - method creates a "TrustedOauthIssuer"
func (svc *TrustedOauthIssuerService) CreateTrustedOauthIssuer(obj *nimbleos.TrustedOauthIssuer) (*nimbleos.TrustedOauthIssuer, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateTrustedOauthIssuer: invalid parameter specified, %v",obj)
    }

    trustedOauthIssuerResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return trustedOauthIssuerResp,nil
}

// UpdateTrustedOauthIssuer - method modifies  the "TrustedOauthIssuer"
func (svc *TrustedOauthIssuerService) UpdateTrustedOauthIssuer(id string, obj *nimbleos.TrustedOauthIssuer) (*nimbleos.TrustedOauthIssuer, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateTrustedOauthIssuer: invalid parameter specified, %v",obj)
    }

    trustedOauthIssuerResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return trustedOauthIssuerResp,nil
}

// GetTrustedOauthIssuerById - method returns a pointer to "TrustedOauthIssuer"
func (svc *TrustedOauthIssuerService) GetTrustedOauthIssuerById(id string) (*nimbleos.TrustedOauthIssuer, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetTrustedOauthIssuerById: invalid parameter specified, %v",id)
    }

    trustedOauthIssuerResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return trustedOauthIssuerResp,nil
}

// GetTrustedOauthIssuerByName - method returns a pointer "TrustedOauthIssuer"
func (svc *TrustedOauthIssuerService) GetTrustedOauthIssuerByName(name string) (*nimbleos.TrustedOauthIssuer, error) {
    params := &param.GetParams{
        Filter: &param.SearchFilter{
            FieldName: nimbleos.VolumeFields.Name,
            Operator:  param.EQUALS.String(),
            Value:     name,
        },
    }
    trustedOauthIssuerResp, err := svc.objectSet.GetObjectListFromParams(params)
    if err != nil {
        return nil, err
    }

    if len(trustedOauthIssuerResp) == 0 {
        return nil, nil
    }

    return trustedOauthIssuerResp[0],nil
}
// DeleteTrustedOauthIssuer - deletes the "TrustedOauthIssuer"
func (svc *TrustedOauthIssuerService) DeleteTrustedOauthIssuer(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteTrustedOauthIssuer: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

