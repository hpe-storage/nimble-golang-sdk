// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsCachePolicy Enum.</p>
 */

type NsCachePolicy string 

const (
	NSCACHEPOLICY_NORMAL NsCachePolicy = "normal"
	NSCACHEPOLICY_NO_WRITE NsCachePolicy = "no_write"
	NSCACHEPOLICY_AGGRESSIVE_READ_NO_WRITE NsCachePolicy = "aggressive_read_no_write"
	NSCACHEPOLICY_DISABLED NsCachePolicy = "disabled"
	NSCACHEPOLICY_AGGRESSIVE NsCachePolicy = "aggressive"

) 
