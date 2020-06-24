/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsSubnetType.</p>
 */

type NsSubnetType string 

const (
	NSSUBNETTYPE_MGMT NsSubnetType = "mgmt"
	NSSUBNETTYPE_UNCONFIGURED NsSubnetType = "unconfigured"
	NSSUBNETTYPE_DATA NsSubnetType = "data"
	NSSUBNETTYPE_MGMT_DATA NsSubnetType = "mgmt_data"
	NSSUBNETTYPE_INVALID NsSubnetType = "invalid"

) 
