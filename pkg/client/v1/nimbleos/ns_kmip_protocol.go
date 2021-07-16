// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsKmipProtocol Enum.

type NsKmipProtocol string

const (
	cNsKmipProtocolKmip11 NsKmipProtocol = "KMIP1_1"
	cNsKmipProtocolKmip12 NsKmipProtocol = "KMIP1_2"
	cNsKmipProtocolKmip10 NsKmipProtocol = "KMIP1_0"
	cNsKmipProtocolKmip13 NsKmipProtocol = "KMIP1_3"
)

var pNsKmipProtocolKmip11 NsKmipProtocol
var pNsKmipProtocolKmip12 NsKmipProtocol
var pNsKmipProtocolKmip10 NsKmipProtocol
var pNsKmipProtocolKmip13 NsKmipProtocol

// NsKmipProtocolKmip11 enum exports
var NsKmipProtocolKmip11 *NsKmipProtocol

// NsKmipProtocolKmip12 enum exports
var NsKmipProtocolKmip12 *NsKmipProtocol

// NsKmipProtocolKmip10 enum exports
var NsKmipProtocolKmip10 *NsKmipProtocol

// NsKmipProtocolKmip13 enum exports
var NsKmipProtocolKmip13 *NsKmipProtocol

func init() {
	pNsKmipProtocolKmip11 = cNsKmipProtocolKmip11
	NsKmipProtocolKmip11 = &pNsKmipProtocolKmip11

	pNsKmipProtocolKmip12 = cNsKmipProtocolKmip12
	NsKmipProtocolKmip12 = &pNsKmipProtocolKmip12

	pNsKmipProtocolKmip10 = cNsKmipProtocolKmip10
	NsKmipProtocolKmip10 = &pNsKmipProtocolKmip10

	pNsKmipProtocolKmip13 = cNsKmipProtocolKmip13
	NsKmipProtocolKmip13 = &pNsKmipProtocolKmip13

}
