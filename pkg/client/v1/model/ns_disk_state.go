// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsDiskState Enum.</p>
 */

type NsDiskState string 

const (
	NSDISKSTATE_VALID NsDiskState = "valid"
	NSDISKSTATE_VOID NsDiskState = "void"
	NSDISKSTATE_REMOVED NsDiskState = "removed"
	NSDISKSTATE_T_FAIL NsDiskState = "t_fail"
	NSDISKSTATE_IN_USE NsDiskState = "in use"
	NSDISKSTATE_ABSENT NsDiskState = "absent"
	NSDISKSTATE_FAILED NsDiskState = "failed"
	NSDISKSTATE_FOREIGN NsDiskState = "foreign"

) 
