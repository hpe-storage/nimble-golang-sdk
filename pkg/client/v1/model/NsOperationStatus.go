/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsOperationStatus.</p>
 */

type NsOperationStatus string 

const (
	NSOPERATIONSTATUS_INPROGRESS NsOperationStatus = "inprogress"
	NSOPERATIONSTATUS_FAILED NsOperationStatus = "failed"
	NSOPERATIONSTATUS_UNKNOWN NsOperationStatus = "unknown"
	NSOPERATIONSTATUS_SUCCEEDED NsOperationStatus = "succeeded"

) 
