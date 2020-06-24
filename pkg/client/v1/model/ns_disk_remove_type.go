// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsDiskRemoveType Enum.</p>
 */

type NsDiskRemoveType string 

const (
	NSDISKREMOVETYPE_TEMP_FAIL NsDiskRemoveType = "temp_fail"
	NSDISKREMOVETYPE_LDR_REMOVE NsDiskRemoveType = "ldr_remove"
	NSDISKREMOVETYPE_PERM_FAIL NsDiskRemoveType = "perm_fail"
	NSDISKREMOVETYPE_REMOVE NsDiskRemoveType = "remove"
	NSDISKREMOVETYPE_DIAG_REMOVE NsDiskRemoveType = "diag_remove"

) 
