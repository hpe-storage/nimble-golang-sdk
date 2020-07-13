// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsDiskOp Enum.

type NsDiskOp string

const (
	cNsDiskOpAdd    NsDiskOp = "add"
	cNsDiskOpRemove NsDiskOp = "remove"
)

var pNsDiskOpAdd NsDiskOp
var pNsDiskOpRemove NsDiskOp

// NsDiskOpAdd enum exports
var NsDiskOpAdd *NsDiskOp

// NsDiskOpRemove enum exports
var NsDiskOpRemove *NsDiskOp

func init() {
	pNsDiskOpAdd = cNsDiskOpAdd
	NsDiskOpAdd = &pNsDiskOpAdd

	pNsDiskOpRemove = cNsDiskOpRemove
	NsDiskOpRemove = &pNsDiskOpRemove

}
