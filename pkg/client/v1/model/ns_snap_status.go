// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsSnapStatus Enum.</p>
 */

type NsSnapStatus string 

const (
	NSSNAPSTATUS_OKAY NsSnapStatus = "okay"
	NSSNAPSTATUS_DELETED NsSnapStatus = "deleted"
	NSSNAPSTATUS_DEFERRED_DELETE NsSnapStatus = "deferred_delete"
	NSSNAPSTATUS_FAILED_DELETE NsSnapStatus = "failed_delete"
	NSSNAPSTATUS_FORCE_DELETE NsSnapStatus = "force_delete"
	NSSNAPSTATUS_CREATE_RETRY NsSnapStatus = "create_retry"
	NSSNAPSTATUS_INVALID NsSnapStatus = "invalid"
	NSSNAPSTATUS_CREATING NsSnapStatus = "creating"

) 
