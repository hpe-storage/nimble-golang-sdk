// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsSnapStatus Enum.

type NsSnapStatus string

const (
	cNsSnapStatusOkay           NsSnapStatus = "okay"
	cNsSnapStatusDeleted        NsSnapStatus = "deleted"
	cNsSnapStatusDeferredDelete NsSnapStatus = "deferred_delete"
	cNsSnapStatusFailedDelete   NsSnapStatus = "failed_delete"
	cNsSnapStatusForceDelete    NsSnapStatus = "force_delete"
	cNsSnapStatusCreateRetry    NsSnapStatus = "create_retry"
	cNsSnapStatusInvalid        NsSnapStatus = "invalid"
	cNsSnapStatusCreating       NsSnapStatus = "creating"
)

var pNsSnapStatusOkay NsSnapStatus
var pNsSnapStatusDeleted NsSnapStatus
var pNsSnapStatusDeferredDelete NsSnapStatus
var pNsSnapStatusFailedDelete NsSnapStatus
var pNsSnapStatusForceDelete NsSnapStatus
var pNsSnapStatusCreateRetry NsSnapStatus
var pNsSnapStatusInvalid NsSnapStatus
var pNsSnapStatusCreating NsSnapStatus

// NsSnapStatusOkay enum exports
var NsSnapStatusOkay *NsSnapStatus

// NsSnapStatusDeleted enum exports
var NsSnapStatusDeleted *NsSnapStatus

// NsSnapStatusDeferredDelete enum exports
var NsSnapStatusDeferredDelete *NsSnapStatus

// NsSnapStatusFailedDelete enum exports
var NsSnapStatusFailedDelete *NsSnapStatus

// NsSnapStatusForceDelete enum exports
var NsSnapStatusForceDelete *NsSnapStatus

// NsSnapStatusCreateRetry enum exports
var NsSnapStatusCreateRetry *NsSnapStatus

// NsSnapStatusInvalid enum exports
var NsSnapStatusInvalid *NsSnapStatus

// NsSnapStatusCreating enum exports
var NsSnapStatusCreating *NsSnapStatus

func init() {
	pNsSnapStatusOkay = cNsSnapStatusOkay
	NsSnapStatusOkay = &pNsSnapStatusOkay

	pNsSnapStatusDeleted = cNsSnapStatusDeleted
	NsSnapStatusDeleted = &pNsSnapStatusDeleted

	pNsSnapStatusDeferredDelete = cNsSnapStatusDeferredDelete
	NsSnapStatusDeferredDelete = &pNsSnapStatusDeferredDelete

	pNsSnapStatusFailedDelete = cNsSnapStatusFailedDelete
	NsSnapStatusFailedDelete = &pNsSnapStatusFailedDelete

	pNsSnapStatusForceDelete = cNsSnapStatusForceDelete
	NsSnapStatusForceDelete = &pNsSnapStatusForceDelete

	pNsSnapStatusCreateRetry = cNsSnapStatusCreateRetry
	NsSnapStatusCreateRetry = &pNsSnapStatusCreateRetry

	pNsSnapStatusInvalid = cNsSnapStatusInvalid
	NsSnapStatusInvalid = &pNsSnapStatusInvalid

	pNsSnapStatusCreating = cNsSnapStatusCreating
	NsSnapStatusCreating = &pNsSnapStatusCreating

}
