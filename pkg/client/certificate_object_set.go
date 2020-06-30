// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"reflect"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/client/v1/model"
	"github.hpe.com/nimble-dcs/golang-sdk/pkg/util"
)


// Manage certificates used by SSL/TLS.
const (
    certificatePath = "certificates"
)

// CertificateObjectSet
type CertificateObjectSet struct {
    Client *GroupMgmtClient
}

// CreateObject creates a new Certificate object
func (objectSet *CertificateObjectSet) CreateObject(payload *model.Certificate) (*model.Certificate, error) {
	response, err := objectSet.Client.Post(certificatePath, payload)
	return response.(*model.Certificate), err
}

// UpdateObject Modify existing Certificate object
func (objectSet *CertificateObjectSet) UpdateObject(id string, payload *model.Certificate) (*model.Certificate, error) {
	response, err := objectSet.Client.Put(certificatePath, id, payload)
	return response.(*model.Certificate), err
}

// DeleteObject deletes the Certificate object with the specified ID
func (objectSet *CertificateObjectSet) DeleteObject(id string) error {
	return objectSet.Client.Delete(certificatePath, id)
}

// GetObject returns a Certificate object with the given ID
func (objectSet *CertificateObjectSet) GetObject(id string) (*model.Certificate, error) {
	response, err := objectSet.Client.Get(certificatePath, id, model.Certificate{})
	if response == nil {
		return nil, err
	}
	return response.(*model.Certificate), err
}

// GetObjectList returns the list of Certificate objects
func (objectSet *CertificateObjectSet) GetObjectList() ([]*model.Certificate, error) {
	response, err := objectSet.Client.List(certificatePath)
	return buildCertificateObjectSet(response), err
}

// GetObjectListFromParams returns the list of Certificate objects using the given params query info
func (objectSet *CertificateObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Certificate, error) {
	response, err := objectSet.Client.ListFromParams(certificatePath, params)
	return buildCertificateObjectSet(response), err
}

// generated function to build the appropriate response types
func buildCertificateObjectSet(response interface{}) ([]*model.Certificate) {
	values := reflect.ValueOf(response)
	results := make([]*model.Certificate, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Certificate{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}