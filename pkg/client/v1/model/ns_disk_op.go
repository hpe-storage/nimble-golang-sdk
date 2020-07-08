// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsDiskOp Enum.

type NsDiskOp string

const (
	cNsDiskOpAdd    NsDiskOp = "add"
	cNsDiskOpRemove NsDiskOp = "remove"
)

var pNsDiskOpAdd NsDiskOp
var pNsDiskOpRemove NsDiskOp

// Export
var NsDiskOpAdd *NsDiskOp
var NsDiskOpRemove *NsDiskOp

func init() {
	pNsDiskOpAdd = cNsDiskOpAdd
	NsDiskOpAdd = &pNsDiskOpAdd

	pNsDiskOpRemove = cNsDiskOpRemove
	NsDiskOpRemove = &pNsDiskOpRemove

}
