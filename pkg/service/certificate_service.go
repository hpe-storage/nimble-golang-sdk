// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Certificate Service - Manage certificates used by SSL/TLS.

import (
	"fmt"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)

// CertificateService type 
type CertificateService struct {
	objectSet *client.CertificateObjectSet
}

// NewCertificateService - method to initialize "CertificateService" 
func NewCertificateService(gs *NsGroupService) (*CertificateService) {
	objectSet := gs.client.GetCertificateObjectSet()
	return &CertificateService{objectSet: objectSet}
}

// GetCertificates - method returns a array of pointers of type "Certificates"
func (svc *CertificateService) GetCertificates(params *util.GetParams) ([]*model.Certificate, error) {
	return svc.objectSet.GetObjectListFromParams(params)
}

// GetCertificatesWithFields - method returns a array of pointers of type "Certificate" 
func (svc *CertificateService) GetCertificatesWithFields(fields []string) ([]*model.Certificate, error) {
	params := &util.GetParams{}
	params.WithFields(fields)
	return svc.objectSet.GetObjectListFromParams(params)
}

// CreateCertificate - method creates a "Certificate"
func (svc *CertificateService) CreateCertificate(obj *model.Certificate) (*model.Certificate, error) {
	// TODO: validate parameters
	return svc.objectSet.CreateObject(obj)
}

// EditCertificate - method modifies  the "Certificate" 
func (svc *CertificateService) EditCertificate(id string, obj *model.Certificate) (*model.Certificate, error) {
	return svc.objectSet.UpdateObject(id, obj)
}

// onlyCertificate - private method for more than one element check. 
func onlyCertificate(objs []*model.Certificate) (*model.Certificate, error) {
	if len(objs) == 0 {
		return nil, nil
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("More than one Certificate found with the given filter")
	}

	return objs[0], nil
}

 
// GetCertificatesByID - method returns associative a array of pointers of type "Certificate", filter by Id
func (svc *CertificateService) GetCertificatesByID(pool *model.Pool, fields []string) (map[string]*model.Certificate, error) {
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
	objs, err := svc.GetCertificates(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Certificate)
	for _, obj := range objs {
		objMap[*obj.ID] = obj
	}
	return objMap, nil
}

// GetCertificateById - method returns a pointer to "Certificate"
func (svc *CertificateService) GetCertificateById(id string) (*model.Certificate, error) {
	return svc.objectSet.GetObject(id)
}

// GetCertificatesByName - method returns a associative array of pointers of type "Certificate", filter by name 
func (svc *CertificateService) GetCertificatesByName(pool *model.Pool, fields []string) (map[string]*model.Certificate, error) {
	params := &util.GetParams{}

	// make sure ID and Name fields are always selected
	if _, found := params.FindField("name"); !found {
		fields = append(fields, "name")
	}
	params.WithFields(fields)

	// check if requested to filter by pool
	if pool != nil {
		filter := &util.SearchFilter{}
		filter.Init("pool_id", util.EQUALS, *pool.ID, false)
		params.WithSearchFilter(filter)
	}
	objs, err := svc.GetCertificates(params)
	if err != nil {
		return nil, err
	}
	objMap := make(map[string]*model.Certificate)
	for _, obj := range objs {
		objMap[*obj.Name] = obj
	}
	return objMap, nil
}

// GetCertificateByName - method returns a pointer "Certificate" 
func (svc *CertificateService) GetCertificateByName(name string) (*model.Certificate, error) {
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
	return onlyCertificate(objs)
}	

