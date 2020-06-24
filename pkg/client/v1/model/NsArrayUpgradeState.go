/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsArrayUpgradeState.</p>
 */

type NsArrayUpgradeState string 

const (
	NSARRAYUPGRADESTATE_ABORT_FAILED NsArrayUpgradeState = "abort_failed"
	NSARRAYUPGRADESTATE_PRECHECK NsArrayUpgradeState = "precheck"
	NSARRAYUPGRADESTATE_IN_PROGRESS NsArrayUpgradeState = "in_progress"
	NSARRAYUPGRADESTATE_PAUSED NsArrayUpgradeState = "paused"
	NSARRAYUPGRADESTATE_ABORTED NsArrayUpgradeState = "aborted"
	NSARRAYUPGRADESTATE_ABORTING NsArrayUpgradeState = "aborting"
	NSARRAYUPGRADESTATE_STARTED NsArrayUpgradeState = "started"
	NSARRAYUPGRADESTATE_NONE NsArrayUpgradeState = "none"
	NSARRAYUPGRADESTATE_FAILED NsArrayUpgradeState = "failed"
	NSARRAYUPGRADESTATE_AWAITING_NEXT_STAGE NsArrayUpgradeState = "awaiting_next_stage"
	NSARRAYUPGRADESTATE_COMPLETE NsArrayUpgradeState = "complete"

) 
