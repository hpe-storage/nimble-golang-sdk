// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsOperationStatusInprogress *NsOperationStatus
var NsOperationStatusFailed *NsOperationStatus
var NsOperationStatusUnknown *NsOperationStatus
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
