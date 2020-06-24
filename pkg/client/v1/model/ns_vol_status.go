// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsVolStatus Enum.</p>
 */

type NsVolStatus string 

const (
	NSVOLSTATUS_OFFLINE NsVolStatus = "offline"
	NSVOLSTATUS_LOGIN_ONLY NsVolStatus = "login_only"
	NSVOLSTATUS_NON_WRITABLE NsVolStatus = "non_writable"
	NSVOLSTATUS_READ_ONLY NsVolStatus = "read_only"
	NSVOLSTATUS_ONLINE NsVolStatus = "online"

) 
