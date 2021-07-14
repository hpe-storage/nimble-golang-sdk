// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// User Service - Represents users configured to manage the system.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// UserService type
type UserService struct {
    objectSet *client.UserObjectSet
}

// NewUserService - method to initialize "UserService"
func NewUserService(gs *NsGroupService) (*UserService) {
    objectSet := gs.client.GetUserObjectSet()
    return &UserService{objectSet: objectSet}
}

// GetUsers - method returns a array of pointers of type "Users"
func (svc *UserService) GetUsers(params *param.GetParams) ([]*nimbleos.User, error) {
    userResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return userResp,nil
}

// CreateUser - method creates a "User"
func (svc *UserService) CreateUser(obj *nimbleos.User) (*nimbleos.User, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateUser: invalid parameter specified, %v",obj)
    }

    userResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return userResp,nil
}

// UpdateUser - method modifies  the "User"
func (svc *UserService) UpdateUser(id string, obj *nimbleos.User) (*nimbleos.User, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateUser: invalid parameter specified, %v",obj)
    }

    userResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return userResp,nil
}

// GetUserById - method returns a pointer to "User"
func (svc *UserService) GetUserById(id string) (*nimbleos.User, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetUserById: invalid parameter specified, %v",id)
    }

    userResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return userResp,nil
}

// GetUserByName - method returns a pointer "User"
func (svc *UserService) GetUserByName(name string) (*nimbleos.User, error) {
    params := &param.GetParams{
        Filter: &param.SearchFilter{
            FieldName: nimbleos.VolumeFields.Name,
            Operator:  param.EQUALS.String(),
            Value:     name,
        },
    }
    userResp, err := svc.objectSet.GetObjectListFromParams(params)
    if err != nil {
        return nil, err
    }

    if len(userResp) == 0 {
        return nil, nil
    }

    return userResp[0],nil
}
// DeleteUser - deletes the "User"
func (svc *UserService) DeleteUser(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteUser: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

// UnlockUser - unlocks user account locked due to failed logins.
//   Required parameters:
//       id - ID for the user.


func (svc *UserService) UnlockUser( id string ) ( *nimbleos.NsUserLockStatus , error) {


     if  len(id) == 0  {
return nil,fmt.Errorf("UnlockUser: invalid parameter specified id: %v ", id )
     }

resp, err := svc.objectSet.Unlock( &id )
     return resp, err
}
