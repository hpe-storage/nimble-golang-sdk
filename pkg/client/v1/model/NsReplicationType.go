/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsReplicationType.</p>
 */

type NsReplicationType string 

const (
	NSREPLICATIONTYPE_SYNCHRONOUS NsReplicationType = "synchronous"
	NSREPLICATIONTYPE_PERIODIC_SNAPSHOT NsReplicationType = "periodic_snapshot"

) 
