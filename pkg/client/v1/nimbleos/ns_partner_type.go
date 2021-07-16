// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPartnerType Enum.

type NsPartnerType string

const (
	cNsPartnerTypeTunnelEndpoint  NsPartnerType = "tunnel_endpoint"
	cNsPartnerTypePool            NsPartnerType = "pool"
	cNsPartnerTypeTunnelInitiator NsPartnerType = "tunnel_initiator"
	cNsPartnerTypeGroup           NsPartnerType = "group"
)

var pNsPartnerTypeTunnelEndpoint NsPartnerType
var pNsPartnerTypePool NsPartnerType
var pNsPartnerTypeTunnelInitiator NsPartnerType
var pNsPartnerTypeGroup NsPartnerType

// NsPartnerTypeTunnelEndpoint enum exports
var NsPartnerTypeTunnelEndpoint *NsPartnerType

// NsPartnerTypePool enum exports
var NsPartnerTypePool *NsPartnerType

// NsPartnerTypeTunnelInitiator enum exports
var NsPartnerTypeTunnelInitiator *NsPartnerType

// NsPartnerTypeGroup enum exports
var NsPartnerTypeGroup *NsPartnerType

func init() {
	pNsPartnerTypeTunnelEndpoint = cNsPartnerTypeTunnelEndpoint
	NsPartnerTypeTunnelEndpoint = &pNsPartnerTypeTunnelEndpoint

	pNsPartnerTypePool = cNsPartnerTypePool
	NsPartnerTypePool = &pNsPartnerTypePool

	pNsPartnerTypeTunnelInitiator = cNsPartnerTypeTunnelInitiator
	NsPartnerTypeTunnelInitiator = &pNsPartnerTypeTunnelInitiator

	pNsPartnerTypeGroup = cNsPartnerTypeGroup
	NsPartnerTypeGroup = &pNsPartnerTypeGroup

}
