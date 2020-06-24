/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsUpdateState.</p>
 */

type NsUpdateState string 

const (
	NSUPDATESTATE_NORMAL NsUpdateState = "normal"
	NSUPDATESTATE_PAUSED NsUpdateState = "paused"
	NSUPDATESTATE_UPDATING NsUpdateState = "updating"
	NSUPDATESTATE_INVALID NsUpdateState = "invalid"
	NSUPDATESTATE_TIMED_OUT NsUpdateState = "timed_out"
	NSUPDATESTATE_FAILED NsUpdateState = "failed"

) 
