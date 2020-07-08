// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsShelfCtrlrSide Enum.

type NsShelfCtrlrSide string

const (
	cNsShelfCtrlrSideA       NsShelfCtrlrSide = "A"
	cNsShelfCtrlrSideB       NsShelfCtrlrSide = "B"
	cNsShelfCtrlrSideUnknown NsShelfCtrlrSide = "unknown"
)

var pNsShelfCtrlrSideA NsShelfCtrlrSide
var pNsShelfCtrlrSideB NsShelfCtrlrSide
var pNsShelfCtrlrSideUnknown NsShelfCtrlrSide

// Export
var NsShelfCtrlrSideA *NsShelfCtrlrSide
var NsShelfCtrlrSideB *NsShelfCtrlrSide
var NsShelfCtrlrSideUnknown *NsShelfCtrlrSide

func init() {
	pNsShelfCtrlrSideA = cNsShelfCtrlrSideA
	NsShelfCtrlrSideA = &pNsShelfCtrlrSideA

	pNsShelfCtrlrSideB = cNsShelfCtrlrSideB
	NsShelfCtrlrSideB = &pNsShelfCtrlrSideB

	pNsShelfCtrlrSideUnknown = cNsShelfCtrlrSideUnknown
	NsShelfCtrlrSideUnknown = &pNsShelfCtrlrSideUnknown

}
