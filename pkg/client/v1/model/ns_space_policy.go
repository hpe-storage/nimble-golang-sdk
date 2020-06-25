// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsSpacePolicy Enum.
 
type NsSpacePolicy string 

const (
	NSSPACEPOLICY_OFFLINE NsSpacePolicy = "offline"
	NSSPACEPOLICY_LOGIN_ONLY NsSpacePolicy = "login_only"
	NSSPACEPOLICY_NON_WRITABLE NsSpacePolicy = "non_writable"
	NSSPACEPOLICY_READ_ONLY NsSpacePolicy = "read_only"
	NSSPACEPOLICY_INVALID NsSpacePolicy = "invalid"

) 
