// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsPartnerTypeTunnelEndpoint *NsPartnerType
var NsPartnerTypePool *NsPartnerType
var NsPartnerTypeTunnelInitiator *NsPartnerType
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
