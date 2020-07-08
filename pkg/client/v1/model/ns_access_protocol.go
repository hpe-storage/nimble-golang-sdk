// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsAccessProtocol Enum.

type NsAccessProtocol string

const (
	cNsAccessProtocolIscsi NsAccessProtocol = "iscsi"
	cNsAccessProtocolFc    NsAccessProtocol = "fc"
)

var pNsAccessProtocolIscsi NsAccessProtocol
var pNsAccessProtocolFc NsAccessProtocol

// Export
var NsAccessProtocolIscsi *NsAccessProtocol
var NsAccessProtocolFc *NsAccessProtocol

func init() {
	pNsAccessProtocolIscsi = cNsAccessProtocolIscsi
	NsAccessProtocolIscsi = &pNsAccessProtocolIscsi

	pNsAccessProtocolFc = cNsAccessProtocolFc
	NsAccessProtocolFc = &pNsAccessProtocolFc

}
