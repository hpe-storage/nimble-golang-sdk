// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ApplicationServer Service - An application server is an external agent that collaborates with an array to manage storage resources; for example, Volume Shadow Copy Service (VSS) or VMware.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// ApplicationServerService type
type ApplicationServerService struct {
	objectSet *client.ApplicationServerObjectSet
}

// NewApplicationServerService - method to initialize "ApplicationServerService"
func NewApplicationServerService(gs *NsGroupService) *ApplicationServerService {
	objectSet := gs.client.GetApplicationServerObjectSet()
	return &ApplicationServerService{objectSet: objectSet}
}

// GetApplicationServers - method returns a array of pointers of type "ApplicationServers"
func (svc *ApplicationServerService) GetApplicationServers(params *param.GetParams) ([]*nimbleos.ApplicationServer, error) {
	if params == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", params)
	}

	applicationServerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	return applicationServerResp, nil
}

// CreateApplicationServer - method creates a "ApplicationServer"
func (svc *ApplicationServerService) CreateApplicationServer(obj *nimbleos.ApplicationServer) (*nimbleos.ApplicationServer, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	applicationServerResp, err := svc.objectSet.CreateObject(obj)
	if err != nil {
		return nil, err
	}
	return applicationServerResp, nil
}

// UpdateApplicationServer - method modifies  the "ApplicationServer"
func (svc *ApplicationServerService) UpdateApplicationServer(id string, obj *nimbleos.ApplicationServer) (*nimbleos.ApplicationServer, error) {
	if obj == nil {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", obj)
	}

	applicationServerResp, err := svc.objectSet.UpdateObject(id, obj)
	if err != nil {
		return nil, err
	}
	return applicationServerResp, nil
}

// GetApplicationServerById - method returns a pointer to "ApplicationServer"
func (svc *ApplicationServerService) GetApplicationServerById(id string) (*nimbleos.ApplicationServer, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("error: invalid parameter specified, %v", id)
	}

	applicationServerResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil, err
	}
	return applicationServerResp, nil
}

// GetApplicationServerByName - method returns a pointer "ApplicationServer"
func (svc *ApplicationServerService) GetApplicationServerByName(name string) (*nimbleos.ApplicationServer, error) {
	params := &param.GetParams{
		Filter: &param.SearchFilter{
			FieldName: &nimbleos.VolumeFields.Name,
			Operator:  param.EQUALS.String(),
			Value:     name,
		},
	}
	applicationServerResp, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}

	if len(applicationServerResp) == 0 {
		return nil, nil
	}

	return applicationServerResp[0], nil
}

// DeleteApplicationServer - deletes the "ApplicationServer"
func (svc *ApplicationServerService) DeleteApplicationServer(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s", id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}
