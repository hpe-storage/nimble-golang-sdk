/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsJobType.</p>
 */

type NsJobType string 

const (
	NSJOBTYPE_BULK_CHILD NsJobType = "bulk_child"
	NSJOBTYPE_BULK_PARENT NsJobType = "bulk_parent"
	NSJOBTYPE_SIMPLE NsJobType = "simple"

) 
