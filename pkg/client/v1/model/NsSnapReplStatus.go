/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsSnapReplStatus.</p>
 */

type NsSnapReplStatus string 

const (
	NSSNAPREPLSTATUS_FAIL NsSnapReplStatus = "fail"
	NSSNAPREPLSTATUS_IN_PROGRESS NsSnapReplStatus = "in_progress"
	NSSNAPREPLSTATUS_PENDING NsSnapReplStatus = "pending"
	NSSNAPREPLSTATUS_COMPLETE NsSnapReplStatus = "complete"

) 
