// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// LdapDomain Service - Manages the storage array's membership with LDAP servers.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// LdapDomainService type
type LdapDomainService struct {
    objectSet *client.LdapDomainObjectSet
}

// NewLdapDomainService - method to initialize "LdapDomainService"
func NewLdapDomainService(gs *NsGroupService) (*LdapDomainService) {
    objectSet := gs.client.GetLdapDomainObjectSet()
    return &LdapDomainService{objectSet: objectSet}
}

// GetLdapDomains - method returns a array of pointers of type "LdapDomains"
func (svc *LdapDomainService) GetLdapDomains(params *param.GetParams) ([]*nimbleos.LdapDomain, error) {
    ldapDomainResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return ldapDomainResp,nil
}

// CreateLdapDomain - method creates a "LdapDomain"
func (svc *LdapDomainService) CreateLdapDomain(obj *nimbleos.LdapDomain) (*nimbleos.LdapDomain, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateLdapDomain: invalid parameter specified, %v",obj)
    }

    ldapDomainResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return ldapDomainResp,nil
}

// UpdateLdapDomain - method modifies  the "LdapDomain"
func (svc *LdapDomainService) UpdateLdapDomain(id string, obj *nimbleos.LdapDomain) (*nimbleos.LdapDomain, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateLdapDomain: invalid parameter specified, %v",obj)
    }

    ldapDomainResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return ldapDomainResp,nil
}

// GetLdapDomainById - method returns a pointer to "LdapDomain"
func (svc *LdapDomainService) GetLdapDomainById(id string) (*nimbleos.LdapDomain, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetLdapDomainById: invalid parameter specified, %v",id)
    }

    ldapDomainResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return ldapDomainResp,nil
}

// DeleteLdapDomain - deletes the "LdapDomain"
func (svc *LdapDomainService) DeleteLdapDomain(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteLdapDomain: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

// TestUserLdapDomain - tests whether the user exist in the given LDAP Domain.
//   Required parameters:
//       id - Unique identifier for the LDAP Domain.
//       user - Name of the LDAP Domain user.


func (svc *LdapDomainService) TestUserLdapDomain( id string , user string ) ( *nimbleos.NsLdapUser , error) {


     if  len(id) == 0  ||  len(user) == 0  {
return nil,fmt.Errorf("TestUserLdapDomain: invalid parameter specified id: %v, user: %v ", id, user )
     }

resp, err := svc.objectSet.TestUser( &id , &user )
     return resp, err
}
// TestGroupLdapDomain - tests whether the user group exist in the given LDAP Domain.
//   Required parameters:
//       id - Unique identifier for the LDAP Domain.
//       group - Name of the group.


func (svc *LdapDomainService) TestGroupLdapDomain( id string , group string )(error) {

     if  len(id) == 0  ||  len(group) == 0  {
return fmt.Errorf("TestGroupLdapDomain: invalid parameter specified id: %v, group: %v ", id , group )
     }

err := svc.objectSet.TestGroup( &id , &group )
     return err
}
// ReportStatusLdapDomain - reports the LDAP connectivity status of the given LDAP ID.
//   Required parameters:
//       id - Unique identifier for the LDAP Domain.


func (svc *LdapDomainService) ReportStatusLdapDomain( id string ) ( *nimbleos.NsLdapReportStatusReturn , error) {


     if  len(id) == 0  {
return nil,fmt.Errorf("ReportStatusLdapDomain: invalid parameter specified id: %v ", id )
     }

resp, err := svc.objectSet.ReportStatus( &id )
     return resp, err
}
