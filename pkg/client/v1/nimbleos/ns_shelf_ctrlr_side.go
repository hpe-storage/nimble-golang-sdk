// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsShelfCtrlrSide Enum.

type NsShelfCtrlrSide string

const (
 cNsShelfCtrlrSideA NsShelfCtrlrSide = "A"
 cNsShelfCtrlrSideB NsShelfCtrlrSide = "B"
 cNsShelfCtrlrSideUnknown NsShelfCtrlrSide = "unknown"
)

var pNsShelfCtrlrSideA NsShelfCtrlrSide
var pNsShelfCtrlrSideB NsShelfCtrlrSide
var pNsShelfCtrlrSideUnknown NsShelfCtrlrSide

// NsShelfCtrlrSideA enum exports
var NsShelfCtrlrSideA *NsShelfCtrlrSide

// NsShelfCtrlrSideB enum exports
var NsShelfCtrlrSideB *NsShelfCtrlrSide

// NsShelfCtrlrSideUnknown enum exports
var NsShelfCtrlrSideUnknown *NsShelfCtrlrSide

func init() {
 pNsShelfCtrlrSideA = cNsShelfCtrlrSideA
 NsShelfCtrlrSideA = &pNsShelfCtrlrSideA

 pNsShelfCtrlrSideB = cNsShelfCtrlrSideB
 NsShelfCtrlrSideB = &pNsShelfCtrlrSideB

 pNsShelfCtrlrSideUnknown = cNsShelfCtrlrSideUnknown
 NsShelfCtrlrSideUnknown = &pNsShelfCtrlrSideUnknown

}

