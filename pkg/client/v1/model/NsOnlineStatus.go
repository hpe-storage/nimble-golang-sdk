/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsOnlineStatus.</p>
 */

type NsOnlineStatus string 

const (
	NSONLINESTATUS_OFFLINE NsOnlineStatus = "offline"
	NSONLINESTATUS_ONLINE NsOnlineStatus = "online"
	NSONLINESTATUS_PARTIAL NsOnlineStatus = "partial"

) 
