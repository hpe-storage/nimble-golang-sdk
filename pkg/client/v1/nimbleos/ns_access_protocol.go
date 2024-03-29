// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsAccessProtocol Enum.

type NsAccessProtocol string

const (
	cNsAccessProtocolIscsi NsAccessProtocol = "iscsi"
	cNsAccessProtocolFc    NsAccessProtocol = "fc"
)

var pNsAccessProtocolIscsi NsAccessProtocol
var pNsAccessProtocolFc NsAccessProtocol

// NsAccessProtocolIscsi enum exports
var NsAccessProtocolIscsi *NsAccessProtocol

// NsAccessProtocolFc enum exports
var NsAccessProtocolFc *NsAccessProtocol

func init() {
	pNsAccessProtocolIscsi = cNsAccessProtocolIscsi
	NsAccessProtocolIscsi = &pNsAccessProtocolIscsi

	pNsAccessProtocolFc = cNsAccessProtocolFc
	NsAccessProtocolFc = &pNsAccessProtocolFc

}
