// Copyright 2020 Hewlett Packard Enterprise Development LP

// update true
package client

import (
	"fmt"
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
	"strings"
)

// Group is a collection of arrays operating together organized into storage pools.
const (
	groupPath = "groups"
)

// GroupObjectSet
type GroupObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Group object
func (objectSet *GroupObjectSet) CreateObject(payload *nimbleos.Group) (*nimbleos.Group, error) {
	return nil, fmt.Errorf("Unsupported operation 'create' on Group")
}

// UpdateObject Modify existing Group object
func (objectSet *GroupObjectSet) UpdateObject(id string, payload *nimbleos.Group) (*nimbleos.Group, error) {
	resp, err := objectSet.Client.Put(groupPath, id, payload, &nimbleos.Group{})
	if err != nil {
		//process http code 202
		if strings.Contains(err.Error(), "status (202)") {
			if resp != nil {
				ID := resp.(string)
				// Get object
				return objectSet.GetObject(ID)

			}
		}
		return nil, err
	}

	return resp.(*nimbleos.Group), err
}

// DeleteObject deletes the Group object with the specified ID
func (objectSet *GroupObjectSet) DeleteObject(id string) error {
	return fmt.Errorf("Unsupported operation 'delete' on Group")
}

// GetObject returns a Group object with the given ID
func (objectSet *GroupObjectSet) GetObject(id string) (*nimbleos.Group, error) {
	resp, err := objectSet.Client.Get(groupPath, id, &nimbleos.Group{})
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return resp.(*nimbleos.Group), err
}

// GetObjectList returns the list of Group objects
func (objectSet *GroupObjectSet) GetObjectList() ([]*nimbleos.Group, error) {
	resp, err := objectSet.Client.List(groupPath)
	if err != nil {
		return nil, err
	}
	return buildGroupObjectSet(resp), err
}

// GetObjectListFromParams returns the list of Group objects using the given params query info
func (objectSet *GroupObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.Group, error) {
	groupObjectSetResp, err := objectSet.Client.ListFromParams(groupPath, params)
	if err != nil {
		return nil, err
	}
	return buildGroupObjectSet(groupObjectSetResp), err
}

// generated function to build the appropriate response types
func buildGroupObjectSet(response interface{}) []*nimbleos.Group {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.Group, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.Group{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}

// List of supported actions on object sets

// Reboot - Reboot all arrays in the group.
func (objectSet *GroupObjectSet) Reboot(id *string) error {

	rebootUri := groupPath
	rebootUri = rebootUri + "/" + *id
	rebootUri = rebootUri + "/actions/" + "reboot"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(rebootUri, payload, &emptyStruct)
	return err
}

// Halt - Halt all arrays in the group.
func (objectSet *GroupObjectSet) Halt(id *string, force *bool) error {

	haltUri := groupPath
	haltUri = haltUri + "/" + *id
	haltUri = haltUri + "/actions/" + "halt"

	payload := &struct {
		Id    *string `json:"id,omitempty"`
		Force *bool   `json:"force,omitempty"`
	}{
		id,
		force,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(haltUri, payload, &emptyStruct)
	return err
}

// TestAlert - Generate a test alert.
func (objectSet *GroupObjectSet) TestAlert(id *string, level *nimbleos.NsSeverityLevel) error {

	testAlertUri := groupPath
	testAlertUri = testAlertUri + "/" + *id
	testAlertUri = testAlertUri + "/actions/" + "test_alert"

	payload := &struct {
		Id    *string                   `json:"id,omitempty"`
		Level *nimbleos.NsSeverityLevel `json:"level,omitempty"`
	}{
		id,
		level,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(testAlertUri, payload, &emptyStruct)
	return err
}

// SoftwareUpdatePrecheck - Run software update precheck.
func (objectSet *GroupObjectSet) SoftwareUpdatePrecheck(id *string, skipPrecheckMask *uint64) (*nimbleos.NsSoftwareUpdateReturn, error) {

	softwareUpdatePrecheckUri := groupPath
	softwareUpdatePrecheckUri = softwareUpdatePrecheckUri + "/" + *id
	softwareUpdatePrecheckUri = softwareUpdatePrecheckUri + "/actions/" + "software_update_precheck"

	payload := &struct {
		Id               *string `json:"id,omitempty"`
		SkipPrecheckMask *uint64 `json:"skip_precheck_mask,omitempty"`
	}{
		id,
		skipPrecheckMask,
	}

	resp, err := objectSet.Client.Post(softwareUpdatePrecheckUri, payload, &nimbleos.NsSoftwareUpdateReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsSoftwareUpdateReturn), err
}

// SoftwareUpdateStart - Update the group software to the downloaded version.
func (objectSet *GroupObjectSet) SoftwareUpdateStart(id *string, skipStartCheckMask *uint64) (*nimbleos.NsSoftwareUpdateReturn, error) {

	softwareUpdateStartUri := groupPath
	softwareUpdateStartUri = softwareUpdateStartUri + "/" + *id
	softwareUpdateStartUri = softwareUpdateStartUri + "/actions/" + "software_update_start"

	payload := &struct {
		Id                 *string `json:"id,omitempty"`
		SkipStartCheckMask *uint64 `json:"skip_start_check_mask,omitempty"`
	}{
		id,
		skipStartCheckMask,
	}

	resp, err := objectSet.Client.Post(softwareUpdateStartUri, payload, &nimbleos.NsSoftwareUpdateReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsSoftwareUpdateReturn), err
}

// SoftwareDownload - Download software update package.
func (objectSet *GroupObjectSet) SoftwareDownload(id *string, version *string, force *bool) error {

	softwareDownloadUri := groupPath
	softwareDownloadUri = softwareDownloadUri + "/" + *id
	softwareDownloadUri = softwareDownloadUri + "/actions/" + "software_download"

	payload := &struct {
		Id      *string `json:"id,omitempty"`
		Version *string `json:"version,omitempty"`
		Force   *bool   `json:"force,omitempty"`
	}{
		id,
		version,
		force,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(softwareDownloadUri, payload, &emptyStruct)
	return err
}

// SoftwareCancelDownload - Cancel ongoing download of software.
func (objectSet *GroupObjectSet) SoftwareCancelDownload(id *string) error {

	softwareCancelDownloadUri := groupPath
	softwareCancelDownloadUri = softwareCancelDownloadUri + "/" + *id
	softwareCancelDownloadUri = softwareCancelDownloadUri + "/actions/" + "software_cancel_download"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(softwareCancelDownloadUri, payload, &emptyStruct)
	return err
}

// SoftwareUpdateResume - Resume stopped software update.
func (objectSet *GroupObjectSet) SoftwareUpdateResume(id *string) error {

	softwareUpdateResumeUri := groupPath
	softwareUpdateResumeUri = softwareUpdateResumeUri + "/" + *id
	softwareUpdateResumeUri = softwareUpdateResumeUri + "/actions/" + "software_update_resume"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(softwareUpdateResumeUri, payload, &emptyStruct)
	return err
}

// GetGroupDiscoveredList - Get list of discovered groups with arrays that are initialized.
func (objectSet *GroupObjectSet) GetGroupDiscoveredList(id *string, groupName *string) (*nimbleos.NsDiscoveredGroupListReturn, error) {

	getGroupDiscoveredListUri := groupPath
	getGroupDiscoveredListUri = getGroupDiscoveredListUri + "/" + *id
	getGroupDiscoveredListUri = getGroupDiscoveredListUri + "/actions/" + "get_group_discovered_list"

	payload := &struct {
		Id        *string `json:"id,omitempty"`
		GroupName *string `json:"group_name,omitempty"`
	}{
		id,
		groupName,
	}

	resp, err := objectSet.Client.Post(getGroupDiscoveredListUri, payload, &nimbleos.NsDiscoveredGroupListReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsDiscoveredGroupListReturn), err
}

// ValidateMerge - Perform group merge validation.
func (objectSet *GroupObjectSet) ValidateMerge(id *string, srcGroupName *string, srcGroupIp *string, srcUsername *string, srcPassword *string, srcPassphrase *string, skipSecondaryMgmtIp *bool) (*nimbleos.NsGroupMergeReturn, error) {

	validateMergeUri := groupPath
	validateMergeUri = validateMergeUri + "/" + *id
	validateMergeUri = validateMergeUri + "/actions/" + "validate_merge"

	payload := &struct {
		Id                  *string `json:"id,omitempty"`
		SrcGroupName        *string `json:"src_group_name,omitempty"`
		SrcGroupIp          *string `json:"src_group_ip,omitempty"`
		SrcUsername         *string `json:"src_username,omitempty"`
		SrcPassword         *string `json:"src_password,omitempty"`
		SrcPassphrase       *string `json:"src_passphrase,omitempty"`
		SkipSecondaryMgmtIp *bool   `json:"skip_secondary_mgmt_ip,omitempty"`
	}{
		id,
		srcGroupName,
		srcGroupIp,
		srcUsername,
		srcPassword,
		srcPassphrase,
		skipSecondaryMgmtIp,
	}

	resp, err := objectSet.Client.Post(validateMergeUri, payload, &nimbleos.NsGroupMergeReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsGroupMergeReturn), err
}

// Merge - Perform group merge with the specified group.
func (objectSet *GroupObjectSet) Merge(id *string, srcGroupName *string, srcGroupIp *string, srcUsername *string, srcPassword *string, srcPassphrase *string, force *bool, skipSecondaryMgmtIp *bool) (*nimbleos.NsGroupMergeReturn, error) {

	mergeUri := groupPath
	mergeUri = mergeUri + "/" + *id
	mergeUri = mergeUri + "/actions/" + "merge"

	payload := &struct {
		Id                  *string `json:"id,omitempty"`
		SrcGroupName        *string `json:"src_group_name,omitempty"`
		SrcGroupIp          *string `json:"src_group_ip,omitempty"`
		SrcUsername         *string `json:"src_username,omitempty"`
		SrcPassword         *string `json:"src_password,omitempty"`
		SrcPassphrase       *string `json:"src_passphrase,omitempty"`
		Force               *bool   `json:"force,omitempty"`
		SkipSecondaryMgmtIp *bool   `json:"skip_secondary_mgmt_ip,omitempty"`
	}{
		id,
		srcGroupName,
		srcGroupIp,
		srcUsername,
		srcPassword,
		srcPassphrase,
		force,
		skipSecondaryMgmtIp,
	}

	resp, err := objectSet.Client.Post(mergeUri, payload, &nimbleos.NsGroupMergeReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsGroupMergeReturn), err
}

// GetEula - Get URL to download EULA contents.
func (objectSet *GroupObjectSet) GetEula(id *string, locale *nimbleos.NsEulaLocale, format *nimbleos.NsEulaFormat, phase *nimbleos.NsEulaPhase, force *bool) (*nimbleos.NsEulaReturn, error) {

	getEulaUri := groupPath
	getEulaUri = getEulaUri + "/" + *id
	getEulaUri = getEulaUri + "/actions/" + "get_eula"

	payload := &struct {
		Id     *string                `json:"id,omitempty"`
		Locale *nimbleos.NsEulaLocale `json:"locale,omitempty"`
		Format *nimbleos.NsEulaFormat `json:"format,omitempty"`
		Phase  *nimbleos.NsEulaPhase  `json:"phase,omitempty"`
		Force  *bool                  `json:"force,omitempty"`
	}{
		id,
		locale,
		format,
		phase,
		force,
	}

	resp, err := objectSet.Client.Post(getEulaUri, payload, &nimbleos.NsEulaReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsEulaReturn), err
}

// CheckMigrate - Check if the group Management Service can be migrated to the group Management Service backup array.
func (objectSet *GroupObjectSet) CheckMigrate(id *string) error {

	checkMigrateUri := groupPath
	checkMigrateUri = checkMigrateUri + "/" + *id
	checkMigrateUri = checkMigrateUri + "/actions/" + "check_migrate"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(checkMigrateUri, payload, &emptyStruct)
	return err
}

// Migrate - Migrate the group Management Service to the current group Management Service backup array.
func (objectSet *GroupObjectSet) Migrate(id *string) error {

	migrateUri := groupPath
	migrateUri = migrateUri + "/" + *id
	migrateUri = migrateUri + "/actions/" + "migrate"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	var emptyStruct struct{}
	_, err := objectSet.Client.Post(migrateUri, payload, &emptyStruct)
	return err
}

// GetTimezoneList - Get list of group timezones.
func (objectSet *GroupObjectSet) GetTimezoneList(id *string) (*nimbleos.NsTimezonesReturn, error) {

	getTimezoneListUri := groupPath
	getTimezoneListUri = getTimezoneListUri + "/" + *id
	getTimezoneListUri = getTimezoneListUri + "/actions/" + "get_timezone_list"

	payload := &struct {
		Id *string `json:"id,omitempty"`
	}{
		id,
	}

	resp, err := objectSet.Client.Post(getTimezoneListUri, payload, &nimbleos.NsTimezonesReturn{})
	if err != nil {
		return nil, err
	}

	return resp.(*nimbleos.NsTimezonesReturn), err
}
