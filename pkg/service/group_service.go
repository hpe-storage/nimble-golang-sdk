// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// Group Service - Group is a collection of arrays operating together organized into storage pools.

import (
    "fmt"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
    "github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// GroupService type
type GroupService struct {
    objectSet *client.GroupObjectSet
}

// NewGroupService - method to initialize "GroupService"
func NewGroupService(gs *NsGroupService) (*GroupService) {
    objectSet := gs.client.GetGroupObjectSet()
    return &GroupService{objectSet: objectSet}
}

// GetGroups - method returns a array of pointers of type "Groups"
func (svc *GroupService) GetGroups(params *param.GetParams) ([]*nimbleos.Group, error) {
    groupResp,err := svc.objectSet.GetObjectListFromParams(params)
    if err !=nil {
        return nil,err
    }
    return groupResp,nil
}

// CreateGroup - method creates a "Group"
func (svc *GroupService) CreateGroup(obj *nimbleos.Group) (*nimbleos.Group, error) {
    if obj == nil {
        return nil,fmt.Errorf("CreateGroup: invalid parameter specified, %v",obj)
    }

    groupResp,err := svc.objectSet.CreateObject(obj)
    if err !=nil {
        return nil,err
    }
    return groupResp,nil
}

// UpdateGroup - method modifies  the "Group"
func (svc *GroupService) UpdateGroup(id string, obj *nimbleos.Group) (*nimbleos.Group, error) {
    if obj == nil {
        return nil,fmt.Errorf("UpdateGroup: invalid parameter specified, %v",obj)
    }

    groupResp,err :=svc.objectSet.UpdateObject(id, obj)
    if err !=nil {
        return nil,err
    }
    return groupResp,nil
}

// GetGroupById - method returns a pointer to "Group"
func (svc *GroupService) GetGroupById(id string) (*nimbleos.Group, error) {
    if len(id) == 0 {
        return nil,fmt.Errorf("GetGroupById: invalid parameter specified, %v",id)
    }

    groupResp, err := svc.objectSet.GetObject(id)
    if err != nil {
        return nil,err
    }
    return groupResp,nil
}

// GetGroupByName - method returns a pointer "Group"
func (svc *GroupService) GetGroupByName(name string) (*nimbleos.Group, error) {
    params := &param.GetParams{
        Filter: &param.SearchFilter{
            FieldName: nimbleos.VolumeFields.Name,
            Operator:  param.EQUALS.String(),
            Value:     name,
        },
    }
    groupResp, err := svc.objectSet.GetObjectListFromParams(params)
    if err != nil {
        return nil, err
    }

    if len(groupResp) == 0 {
        return nil, nil
    }

    return groupResp[0],nil
}
// DeleteGroup - deletes the "Group"
func (svc *GroupService) DeleteGroup(id string) error {
    if len(id) == 0 {
        return fmt.Errorf("DeleteGroup: invalid parameter specified, %s",id)
    }
    err := svc.objectSet.DeleteObject(id)
    if err != nil {
        return err
    }
    return nil
}

// RebootGroup - reboot all arrays in the group.
//   Required parameters:
//       id - ID of the group to reboot.


func (svc *GroupService) RebootGroup( id string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("RebootGroup: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.Reboot( &id )
     return err
}
// HaltGroup - halt all arrays in the group.
//   Required parameters:
//       id - ID of the group to halt.

//   Optional parameters:
//       force - Halt remaining arrays when one or more is unreachable.

func (svc *GroupService) HaltGroup( id string , force *bool )(error) {

     if  len(id) == 0  {
return fmt.Errorf("HaltGroup: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.Halt( &id , force )
     return err
}
// TestAlertGroup - generate a test alert.
//   Required parameters:
//       id - ID of the group.
//       level - Level of the test alert.


func (svc *GroupService) TestAlertGroup( id string , level *nimbleos.NsSeverityLevel )(error) {

     if  len(id) == 0  {
return fmt.Errorf("TestAlertGroup: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.TestAlert( &id , level )
     return err
}
// SoftwareUpdatePrecheckGroup - run software update precheck.
//   Required parameters:
//       id - ID of the group.

//   Optional parameters:
//       skipPrecheckMask - Flag to allow skipping certain types of prechecks.

func (svc *GroupService) SoftwareUpdatePrecheckGroup( id string , skipPrecheckMask *uint64 ) ( *nimbleos.NsSoftwareUpdateReturn , error) {


     if  len(id) == 0  {
return nil,fmt.Errorf("SoftwareUpdatePrecheckGroup: invalid parameter specified id: %v ", id )
     }

resp, err := svc.objectSet.SoftwareUpdatePrecheck( &id , skipPrecheckMask )
     return resp, err
}
// SoftwareUpdateStartGroup - update the group software to the downloaded version.
//   Required parameters:
//       id - ID of the group.

//   Optional parameters:
//       skipStartCheckMask - Flag to allow skipping certain types of checks.

func (svc *GroupService) SoftwareUpdateStartGroup( id string , skipStartCheckMask *uint64 ) ( *nimbleos.NsSoftwareUpdateReturn , error) {


     if  len(id) == 0  {
return nil,fmt.Errorf("SoftwareUpdateStartGroup: invalid parameter specified id: %v ", id )
     }

resp, err := svc.objectSet.SoftwareUpdateStart( &id , skipStartCheckMask )
     return resp, err
}
// SoftwareDownloadGroup - download software update package.
//   Required parameters:
//       id - ID of the group.
//       version - Version string to download.

//   Optional parameters:
//       force - Flag to force download.

func (svc *GroupService) SoftwareDownloadGroup( id string , version string , force *bool )(error) {

     if  len(id) == 0  ||  len(version) == 0  {
return fmt.Errorf("SoftwareDownloadGroup: invalid parameter specified id: %v, version: %v ", id , version )
     }

err := svc.objectSet.SoftwareDownload( &id , &version , force )
     return err
}
// SoftwareCancelDownloadGroup - cancel ongoing download of software.
//   Required parameters:
//       id - ID of the group.


func (svc *GroupService) SoftwareCancelDownloadGroup( id string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("SoftwareCancelDownloadGroup: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.SoftwareCancelDownload( &id )
     return err
}
// SoftwareUpdateResumeGroup - resume stopped software update.
//   Required parameters:
//       id - ID of the group.


func (svc *GroupService) SoftwareUpdateResumeGroup( id string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("SoftwareUpdateResumeGroup: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.SoftwareUpdateResume( &id )
     return err
}
// GetGroupDiscoveredListGroup - get list of discovered groups with arrays that are initialized.
//   Required parameters:
//       id - ID of the group.

//   Optional parameters:
//       groupName - Name of the group requested to be discovered.

func (svc *GroupService) GetGroupDiscoveredListGroup( id string , groupName *string ) ( *nimbleos.NsDiscoveredGroupListReturn , error) {


     if  len(id) == 0  {
return nil,fmt.Errorf("GetGroupDiscoveredListGroup: invalid parameter specified id: %v ", id )
     }

resp, err := svc.objectSet.GetGroupDiscoveredList( &id , groupName )
     return resp, err
}
// ValidateMergeGroup - perform group merge validation.
//   Required parameters:
//       id - ID of the group.
//       srcGroupName - Name of the source group.
//       srcGroupIp - IP address of the source group.
//       srcUsername - Username of the source group.
//       srcPassword - Password of the source group.

//   Optional parameters:
//       srcPassphrase - Source group encryption passphrase.
//       skipSecondaryMgmtIp - Skip check for secondary management IP address.

func (svc *GroupService) ValidateMergeGroup( id string , srcGroupName string , srcGroupIp string , srcUsername string , srcPassword string , srcPassphrase *string , skipSecondaryMgmtIp *bool ) ( *nimbleos.NsGroupMergeReturn , error) {


     if  len(id) == 0  ||  len(srcGroupName) == 0  ||  len(srcGroupIp) == 0  ||  len(srcUsername) == 0  ||  len(srcPassword) == 0  {
return nil,fmt.Errorf("ValidateMergeGroup: invalid parameter specified id: %v, srcGroupName: %v, srcGroupIp: %v, srcUsername: %v, srcPassword: %v ", id, srcGroupName, srcGroupIp, srcUsername, srcPassword )
     }

resp, err := svc.objectSet.ValidateMerge( &id , &srcGroupName , &srcGroupIp , &srcUsername , &srcPassword , srcPassphrase , skipSecondaryMgmtIp )
     return resp, err
}
// MergeGroup - perform group merge with the specified group.
//   Required parameters:
//       id - ID of the group.
//       srcGroupName - Name of the source group.
//       srcGroupIp - IP address of the source group.
//       srcUsername - Username of the source group.
//       srcPassword - Password of the source group.

//   Optional parameters:
//       srcPassphrase - Source group encryption passphrase.
//       force - Ignore warnings and forcibly merge specified group with this group.
//       skipSecondaryMgmtIp - Skip check for secondary management IP address.

func (svc *GroupService) MergeGroup( id string , srcGroupName string , srcGroupIp string , srcUsername string , srcPassword string , srcPassphrase *string , force *bool , skipSecondaryMgmtIp *bool ) ( *nimbleos.NsGroupMergeReturn , error) {


     if  len(id) == 0  ||  len(srcGroupName) == 0  ||  len(srcGroupIp) == 0  ||  len(srcUsername) == 0  ||  len(srcPassword) == 0  {
return nil,fmt.Errorf("MergeGroup: invalid parameter specified id: %v, srcGroupName: %v, srcGroupIp: %v, srcUsername: %v, srcPassword: %v ", id, srcGroupName, srcGroupIp, srcUsername, srcPassword )
     }

resp, err := svc.objectSet.Merge( &id , &srcGroupName , &srcGroupIp , &srcUsername , &srcPassword , srcPassphrase , force , skipSecondaryMgmtIp )
     return resp, err
}
// GetEulaGroup - get URL to download EULA contents.
//   Required parameters:
//       id - ID of the group.

//   Optional parameters:
//       locale - Locale of EULA contents. Default is en.
//       format - Format of EULA contents. Default is HTML.
//       phase - Phase of EULA contents. Default is setup.
//       force - Flag to force EULA.

func (svc *GroupService) GetEulaGroup( id string , locale *nimbleos.NsEulaLocale , format *nimbleos.NsEulaFormat , phase *nimbleos.NsEulaPhase , force *bool ) ( *nimbleos.NsEulaReturn , error) {


     if  len(id) == 0  {
return nil,fmt.Errorf("GetEulaGroup: invalid parameter specified id: %v ", id )
     }

resp, err := svc.objectSet.GetEula( &id , locale , format , phase , force )
     return resp, err
}
// CheckMigrateGroup - check if the group Management Service can be migrated to the group Management Service backup array.
//   Required parameters:
//       id - ID of the group.


func (svc *GroupService) CheckMigrateGroup( id string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("CheckMigrateGroup: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.CheckMigrate( &id )
     return err
}
// MigrateGroup - migrate the group Management Service to the current group Management Service backup array.
//   Required parameters:
//       id - ID of the group.


func (svc *GroupService) MigrateGroup( id string )(error) {

     if  len(id) == 0  {
return fmt.Errorf("MigrateGroup: invalid parameter specified id: %v ", id )
     }

err := svc.objectSet.Migrate( &id )
     return err
}
// GetTimezoneListGroup - get list of group timezones.
//   Required parameters:
//       id - ID of the group.


func (svc *GroupService) GetTimezoneListGroup( id string ) ( *nimbleos.NsTimezonesReturn , error) {


     if  len(id) == 0  {
return nil,fmt.Errorf("GetTimezoneListGroup: invalid parameter specified id: %v ", id )
     }

resp, err := svc.objectSet.GetTimezoneList( &id )
     return resp, err
}
