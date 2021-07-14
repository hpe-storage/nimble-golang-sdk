// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsControllerId Enum.

type NsControllerId string

const (
 cNsControllerIdA NsControllerId = "A"
 cNsControllerIdB NsControllerId = "B"
 cNsControllerIdIndependent NsControllerId = "independent"
)

var pNsControllerIdA NsControllerId
var pNsControllerIdB NsControllerId
var pNsControllerIdIndependent NsControllerId

// NsControllerIdA enum exports
var NsControllerIdA *NsControllerId

// NsControllerIdB enum exports
var NsControllerIdB *NsControllerId

// NsControllerIdIndependent enum exports
var NsControllerIdIndependent *NsControllerId

func init() {
 pNsControllerIdA = cNsControllerIdA
 NsControllerIdA = &pNsControllerIdA

 pNsControllerIdB = cNsControllerIdB
 NsControllerIdB = &pNsControllerIdB

 pNsControllerIdIndependent = cNsControllerIdIndependent
 NsControllerIdIndependent = &pNsControllerIdIndependent

}

