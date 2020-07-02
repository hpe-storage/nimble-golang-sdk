// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// ApplicationServer Service - An application server is an external agent that collaborates with an array to manage storage resources; for example, Volume Shadow Copy Service (VSS) or VMware.

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// ApplicationServerService type 
type ApplicationServerService struct {
	objectSet *client.ApplicationServerObjectSet
}

// NewApplicationServerService - method to initialize "ApplicationServerService" 
func NewApplicationServerService(gs *NsGroupService) (*ApplicationServerService) {
	objectSet := gs.client.GetApplicationServerObjectSet()
	return &ApplicationServerService{objectSet: objectSet}
}

// GetApplicationServers - method returns a array of pointers of type "ApplicationServers"
func (svc *ApplicationServerService) GetApplicationServers(params *util.GetParams) ([]*model.ApplicationServer, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateApplicationServer - method creates a "ApplicationServer"
func (svc *ApplicationServerService) CreateApplicationServer(obj *model.ApplicationServer) (*model.ApplicationServer, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// UpdateApplicationServer - method modifies  the "ApplicationServer" 
func (svc *ApplicationServerService) UpdateApplicationServer(id string, obj *model.ApplicationServer) (*model.ApplicationServer, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// GetApplicationServerById - method returns a pointer to "ApplicationServer"
func (svc *ApplicationServerService) GetApplicationServerById(id string) (*model.ApplicationServer, error) {
	return svc.objectSet.GetObject(id)
}

// GetApplicationServerByName - method returns a pointer "ApplicationServer" 
func (svc *ApplicationServerService) GetApplicationServerByName(name string) (*model.ApplicationServer, error) {
	params := &util.GetParams{
		Filter: &util.SearchFilter{
			FieldName: model.VolumeFields.Name,
			Operator:  util.EQUALS.String(),
			Value:     name,
		},
	}
	objs, err := svc.objectSet.GetObjectListFromParams(params)
	if err != nil {
		return nil, err
	}
	
	if len(objs) == 0 {
    	return nil, nil
    }
    
	return objs[0],nil
}	

// DeleteApplicationServer - deletes the "ApplicationServer"
func (svc *ApplicationServerService) DeleteApplicationServer(id string) error {
	return svc.objectSet.DeleteObject(id)
}
