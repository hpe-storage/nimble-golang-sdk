/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsJobResult.</p>
 */

type NsJobResult string 

const (
	NSJOBRESULT_FAILED NsJobResult = "failed"
	NSJOBRESULT_PARTIAL NsJobResult = "partial"
	NSJOBRESULT_UNKNOWN NsJobResult = "unknown"
	NSJOBRESULT_SUCCEEDED NsJobResult = "succeeded"

) 
