// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsUpdateState Enum.
 
type NsUpdateState string 

const (
	NSUPDATESTATE_NORMAL NsUpdateState = "normal"
	NSUPDATESTATE_PAUSED NsUpdateState = "paused"
	NSUPDATESTATE_UPDATING NsUpdateState = "updating"
	NSUPDATESTATE_INVALID NsUpdateState = "invalid"
	NSUPDATESTATE_TIMED_OUT NsUpdateState = "timed_out"
	NSUPDATESTATE_FAILED NsUpdateState = "failed"

) 
