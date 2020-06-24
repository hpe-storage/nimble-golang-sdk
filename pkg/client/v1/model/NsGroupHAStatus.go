/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsGroupHAStatus.</p>
 */

type NsGroupHAStatus string 

const (
	NSGROUPHASTATUS_IN_SYNC NsGroupHAStatus = "in_sync"
	NSGROUPHASTATUS_REMOVE_IN_PROGRESS NsGroupHAStatus = "remove_in_progress"
	NSGROUPHASTATUS_YET_TO_SETUP NsGroupHAStatus = "yet_to_setup"
	NSGROUPHASTATUS_UNSETUP_IN_PROGRESS NsGroupHAStatus = "unsetup_in_progress"
	NSGROUPHASTATUS_SETUP_IN_PROGRESS NsGroupHAStatus = "setup_in_progress"
	NSGROUPHASTATUS_OUT_OF_SYNC NsGroupHAStatus = "out_of_sync"
	NSGROUPHASTATUS_ADD_IN_PROGRESS NsGroupHAStatus = "add_in_progress"
	NSGROUPHASTATUS_SETUP_FAILED NsGroupHAStatus = "setup_failed"

) 
