// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsControllerSide Enum.

type NsControllerSide string

const (
	cNsControllerSideA NsControllerSide = "A"
	cNsControllerSideB NsControllerSide = "B"
)

var pNsControllerSideA NsControllerSide
var pNsControllerSideB NsControllerSide

// Export
var NsControllerSideA *NsControllerSide
var NsControllerSideB *NsControllerSide

func init() {
	pNsControllerSideA = cNsControllerSideA
	NsControllerSideA = &pNsControllerSideA

	pNsControllerSideB = cNsControllerSideB
	NsControllerSideB = &pNsControllerSideB

}
