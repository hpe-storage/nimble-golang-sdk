// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsShelfState Enum.</p>
 */

type NsShelfState string 

const (
	NSSHELFSTATE_DISCOVERING NsShelfState = "discovering"
	NSSHELFSTATE_DISCONNECTED NsShelfState = "disconnected"
	NSSHELFSTATE_VOID NsShelfState = "void"
	NSSHELFSTATE_AVAILABLE NsShelfState = "available"
	NSSHELFSTATE_ONLINE NsShelfState = "online"
	NSSHELFSTATE_FAULTY NsShelfState = "faulty"
	NSSHELFSTATE_FOREIGN NsShelfState = "foreign"
	NSSHELFSTATE_UNKNOWN NsShelfState = "unknown"

) 
