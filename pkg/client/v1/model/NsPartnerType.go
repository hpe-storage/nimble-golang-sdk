/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsPartnerType.</p>
 */

type NsPartnerType string 

const (
	NSPARTNERTYPE_TUNNEL_ENDPOINT NsPartnerType = "tunnel_endpoint"
	NSPARTNERTYPE_POOL NsPartnerType = "pool"
	NSPARTNERTYPE_TUNNEL_INITIATOR NsPartnerType = "tunnel_initiator"
	NSPARTNERTYPE_GROUP NsPartnerType = "group"

) 
