// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsGroupHAStatus Enum.

type NsGroupHAStatus string

const (
	cNsGroupHAStatusInSync            NsGroupHAStatus = "in_sync"
	cNsGroupHAStatusRemoveInProgress  NsGroupHAStatus = "remove_in_progress"
	cNsGroupHAStatusYetToSetup        NsGroupHAStatus = "yet_to_setup"
	cNsGroupHAStatusUnsetupInProgress NsGroupHAStatus = "unsetup_in_progress"
	cNsGroupHAStatusSetupInProgress   NsGroupHAStatus = "setup_in_progress"
	cNsGroupHAStatusOutOfSync         NsGroupHAStatus = "out_of_sync"
	cNsGroupHAStatusAddInProgress     NsGroupHAStatus = "add_in_progress"
	cNsGroupHAStatusSetupFailed       NsGroupHAStatus = "setup_failed"
)

var pNsGroupHAStatusInSync NsGroupHAStatus
var pNsGroupHAStatusRemoveInProgress NsGroupHAStatus
var pNsGroupHAStatusYetToSetup NsGroupHAStatus
var pNsGroupHAStatusUnsetupInProgress NsGroupHAStatus
var pNsGroupHAStatusSetupInProgress NsGroupHAStatus
var pNsGroupHAStatusOutOfSync NsGroupHAStatus
var pNsGroupHAStatusAddInProgress NsGroupHAStatus
var pNsGroupHAStatusSetupFailed NsGroupHAStatus

// Export
var NsGroupHAStatusInSync *NsGroupHAStatus
var NsGroupHAStatusRemoveInProgress *NsGroupHAStatus
var NsGroupHAStatusYetToSetup *NsGroupHAStatus
var NsGroupHAStatusUnsetupInProgress *NsGroupHAStatus
var NsGroupHAStatusSetupInProgress *NsGroupHAStatus
var NsGroupHAStatusOutOfSync *NsGroupHAStatus
var NsGroupHAStatusAddInProgress *NsGroupHAStatus
var NsGroupHAStatusSetupFailed *NsGroupHAStatus

func init() {
	pNsGroupHAStatusInSync = cNsGroupHAStatusInSync
	NsGroupHAStatusInSync = &pNsGroupHAStatusInSync

	pNsGroupHAStatusRemoveInProgress = cNsGroupHAStatusRemoveInProgress
	NsGroupHAStatusRemoveInProgress = &pNsGroupHAStatusRemoveInProgress

	pNsGroupHAStatusYetToSetup = cNsGroupHAStatusYetToSetup
	NsGroupHAStatusYetToSetup = &pNsGroupHAStatusYetToSetup

	pNsGroupHAStatusUnsetupInProgress = cNsGroupHAStatusUnsetupInProgress
	NsGroupHAStatusUnsetupInProgress = &pNsGroupHAStatusUnsetupInProgress

	pNsGroupHAStatusSetupInProgress = cNsGroupHAStatusSetupInProgress
	NsGroupHAStatusSetupInProgress = &pNsGroupHAStatusSetupInProgress

	pNsGroupHAStatusOutOfSync = cNsGroupHAStatusOutOfSync
	NsGroupHAStatusOutOfSync = &pNsGroupHAStatusOutOfSync

	pNsGroupHAStatusAddInProgress = cNsGroupHAStatusAddInProgress
	NsGroupHAStatusAddInProgress = &pNsGroupHAStatusAddInProgress

	pNsGroupHAStatusSetupFailed = cNsGroupHAStatusSetupFailed
	NsGroupHAStatusSetupFailed = &pNsGroupHAStatusSetupFailed

}
