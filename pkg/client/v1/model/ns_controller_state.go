// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsControllerState Enum.</p>
 */

type NsControllerState string 

const (
	NSCONTROLLERSTATE_START_ACTIVE NsControllerState = "start_active"
	NSCONTROLLERSTATE_START_STANDBY NsControllerState = "start_standby"
	NSCONTROLLERSTATE_STALE NsControllerState = "stale"
	NSCONTROLLERSTATE_STANDBY NsControllerState = "standby"
	NSCONTROLLERSTATE_ACTIVE NsControllerState = "active"
	NSCONTROLLERSTATE_SOLO NsControllerState = "solo"
	NSCONTROLLERSTATE_NONE NsControllerState = "none"

) 
