// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsOperationStatus Enum.

type NsOperationStatus string

const (
	cNsOperationStatusInprogress NsOperationStatus = "inprogress"
	cNsOperationStatusFailed     NsOperationStatus = "failed"
	cNsOperationStatusUnknown    NsOperationStatus = "unknown"
	cNsOperationStatusSucceeded  NsOperationStatus = "succeeded"
)

var pNsOperationStatusInprogress NsOperationStatus
var pNsOperationStatusFailed NsOperationStatus
var pNsOperationStatusUnknown NsOperationStatus
var pNsOperationStatusSucceeded NsOperationStatus

// NsOperationStatusInprogress enum exports
var NsOperationStatusInprogress *NsOperationStatus

// NsOperationStatusFailed enum exports
var NsOperationStatusFailed *NsOperationStatus

// NsOperationStatusUnknown enum exports
var NsOperationStatusUnknown *NsOperationStatus

// NsOperationStatusSucceeded enum exports
var NsOperationStatusSucceeded *NsOperationStatus

func init() {
	pNsOperationStatusInprogress = cNsOperationStatusInprogress
	NsOperationStatusInprogress = &pNsOperationStatusInprogress

	pNsOperationStatusFailed = cNsOperationStatusFailed
	NsOperationStatusFailed = &pNsOperationStatusFailed

	pNsOperationStatusUnknown = cNsOperationStatusUnknown
	NsOperationStatusUnknown = &pNsOperationStatusUnknown

	pNsOperationStatusSucceeded = cNsOperationStatusSucceeded
	NsOperationStatusSucceeded = &pNsOperationStatusSucceeded

}
