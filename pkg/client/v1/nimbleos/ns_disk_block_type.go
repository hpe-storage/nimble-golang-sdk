// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsDiskBlockType Enum.

type NsDiskBlockType string

const (
	cNsDiskBlockTypeBlock512e NsDiskBlockType = "block_512e"
	cNsDiskBlockTypeBlock4kn  NsDiskBlockType = "block_4Kn"
	cNsDiskBlockTypeBlockNone NsDiskBlockType = "block_none"
	cNsDiskBlockTypeBlock512n NsDiskBlockType = "block_512n"
)

var pNsDiskBlockTypeBlock512e NsDiskBlockType
var pNsDiskBlockTypeBlock4kn NsDiskBlockType
var pNsDiskBlockTypeBlockNone NsDiskBlockType
var pNsDiskBlockTypeBlock512n NsDiskBlockType

// NsDiskBlockTypeBlock512e enum exports
var NsDiskBlockTypeBlock512e *NsDiskBlockType

// NsDiskBlockTypeBlock4kn enum exports
var NsDiskBlockTypeBlock4kn *NsDiskBlockType

// NsDiskBlockTypeBlockNone enum exports
var NsDiskBlockTypeBlockNone *NsDiskBlockType

// NsDiskBlockTypeBlock512n enum exports
var NsDiskBlockTypeBlock512n *NsDiskBlockType

func init() {
	pNsDiskBlockTypeBlock512e = cNsDiskBlockTypeBlock512e
	NsDiskBlockTypeBlock512e = &pNsDiskBlockTypeBlock512e

	pNsDiskBlockTypeBlock4kn = cNsDiskBlockTypeBlock4kn
	NsDiskBlockTypeBlock4kn = &pNsDiskBlockTypeBlock4kn

	pNsDiskBlockTypeBlockNone = cNsDiskBlockTypeBlockNone
	NsDiskBlockTypeBlockNone = &pNsDiskBlockTypeBlockNone

	pNsDiskBlockTypeBlock512n = cNsDiskBlockTypeBlock512n
	NsDiskBlockTypeBlock512n = &pNsDiskBlockTypeBlock512n

}
