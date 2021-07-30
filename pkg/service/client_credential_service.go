// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package service

// ClientCredential Service - Credential that this device will trust.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ClientCredentialService type
type ClientCredentialService struct {
	objectSet *client.ClientCredentialObjectSet
}

// NewClientCredentialService - method to initialize "ClientCredentialService"
func NewClientCredentialService(gs *NsGroupService) *ClientCredentialService {
	objectSet := gs.client.GetClientCredentialObjectSet()
	return &ClientCredentialService{objectSet: objectSet}
}

// GetClientCredentials - method returns a array of pointers of type "ClientCredentials"
func (svc *ClientCredentialService) GetClientCredentials(params *param.GetParams) ([]*nimbleos.ClientCredential, error) {
	clientCredentialResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return clientCredentialResp, nil
}

// CreateClientCredential - method creates a "ClientCredential"
func (svc *ClientCredentialService) CreateClientCredential(obj *nimbleos.ClientCredential) (*nimbleos.ClientCredential, error) {
	if obj == nil {
		return nil, fmt.Errorf("CreateClientCredential: invalid parameter specified, %v", obj)
	}

	clientCredentialResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return clientCredentialResp, nil
}

// UpdateClientCredential - method modifies  the "ClientCredential"
func (svc *ClientCredentialService) UpdateClientCredential(id string, obj *nimbleos.ClientCredential) (*nimbleos.ClientCredential, error) {
	if obj == nil {
		return nil, fmt.Errorf("UpdateClientCredential: invalid parameter specified, %v", obj)
	}

	clientCredentialResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return clientCredentialResp, nil
}

// GetClientCredentialById - method returns a pointer to "ClientCredential"
func (svc *ClientCredentialService) GetClientCredentialById(id string) (*nimbleos.ClientCredential, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("GetClientCredentialById: invalid parameter specified, %v", id)
	}

	clientCredentialResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return clientCredentialResp, nil
}

// GetClientCredentialByName - method returns a pointer "ClientCredential"
func (svc *ClientCredentialService) GetClientCredentialByName(name string) (*nimbleos.ClientCredential, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: param.NewString(nimbleos.VolumeFields.Name),
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	clientCredentialResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(clientCredentialResp) == 0 {
		return nil, nil
	}

	return clientCredentialResp[0], nil
}

// DeleteClientCredential - deletes the "ClientCredential"
func (svc *ClientCredentialService) DeleteClientCredential(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("DeleteClientCredential: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
