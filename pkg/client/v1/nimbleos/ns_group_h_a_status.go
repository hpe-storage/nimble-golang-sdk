// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsGroupHAStatusInSync enum exports
var NsGroupHAStatusInSync *NsGroupHAStatus

// NsGroupHAStatusRemoveInProgress enum exports
var NsGroupHAStatusRemoveInProgress *NsGroupHAStatus

// NsGroupHAStatusYetToSetup enum exports
var NsGroupHAStatusYetToSetup *NsGroupHAStatus

// NsGroupHAStatusUnsetupInProgress enum exports
var NsGroupHAStatusUnsetupInProgress *NsGroupHAStatus

// NsGroupHAStatusSetupInProgress enum exports
var NsGroupHAStatusSetupInProgress *NsGroupHAStatus

// NsGroupHAStatusOutOfSync enum exports
var NsGroupHAStatusOutOfSync *NsGroupHAStatus

// NsGroupHAStatusAddInProgress enum exports
var NsGroupHAStatusAddInProgress *NsGroupHAStatus

// NsGroupHAStatusSetupFailed enum exports
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
