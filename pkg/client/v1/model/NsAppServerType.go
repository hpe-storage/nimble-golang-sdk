/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsAppServerType.</p>
 */

type NsAppServerType string 

const (
	NSAPPSERVERTYPE_VSS NsAppServerType = "vss"
	NSAPPSERVERTYPE_VMWARE NsAppServerType = "vmware"
	NSAPPSERVERTYPE_CISCO NsAppServerType = "cisco"
	NSAPPSERVERTYPE_CONTAINER_NODE NsAppServerType = "container_node"
	NSAPPSERVERTYPE_STACK_VISION NsAppServerType = "stack_vision"

) 
