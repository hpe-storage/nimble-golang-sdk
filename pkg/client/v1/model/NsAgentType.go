/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsAgentType.</p>
 */

type NsAgentType string 

const (
	NSAGENTTYPE_SMIS NsAgentType = "smis"
	NSAGENTTYPE_VVOL NsAgentType = "vvol"
	NSAGENTTYPE_OPENSTACK NsAgentType = "openstack"
	NSAGENTTYPE_OPENSTACKV2 NsAgentType = "openstackv2"
	NSAGENTTYPE_NONE NsAgentType = "none"

) 
