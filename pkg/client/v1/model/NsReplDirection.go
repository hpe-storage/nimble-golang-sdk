/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsReplDirection.</p>
 */

type NsReplDirection string 

const (
	NSREPLDIRECTION_UPSTREAM NsReplDirection = "upstream"
	NSREPLDIRECTION_DOWNSTREAM NsReplDirection = "downstream"
	NSREPLDIRECTION_NONE NsReplDirection = "none"
	NSREPLDIRECTION_BI_DIRECTIONAL NsReplDirection = "bi_directional"

) 
