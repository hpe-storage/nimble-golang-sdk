/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsAccessApplyTo.</p>
 */

type NsAccessApplyTo string 

const (
	NSACCESSAPPLYTO_VOLUME NsAccessApplyTo = "volume"
	NSACCESSAPPLYTO_PE NsAccessApplyTo = "pe"
	NSACCESSAPPLYTO_VVOL_VOLUME NsAccessApplyTo = "vvol_volume"
	NSACCESSAPPLYTO_VVOL_SNAPSHOT NsAccessApplyTo = "vvol_snapshot"
	NSACCESSAPPLYTO_SNAPSHOT NsAccessApplyTo = "snapshot"
	NSACCESSAPPLYTO_BOTH NsAccessApplyTo = "both"

) 
