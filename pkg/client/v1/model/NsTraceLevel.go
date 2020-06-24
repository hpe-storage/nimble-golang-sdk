/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsTraceLevel.</p>
 */

type NsTraceLevel string 

const (
	NSTRACELEVEL_NOTE NsTraceLevel = "note"
	NSTRACELEVEL_WARN NsTraceLevel = "warn"
	NSTRACELEVEL_CRITICAL NsTraceLevel = "critical"
	NSTRACELEVEL_ERROR NsTraceLevel = "error"
	NSTRACELEVEL_INFO NsTraceLevel = "info"

) 
