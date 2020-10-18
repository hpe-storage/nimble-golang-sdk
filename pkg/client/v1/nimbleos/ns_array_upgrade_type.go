// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsArrayUpgradeType Enum.

type NsArrayUpgradeType string

const (
	cNsArrayUpgradeTypeOffline NsArrayUpgradeType = "offline"
	cNsArrayUpgradeTypeInvalid NsArrayUpgradeType = "invalid"
	cNsArrayUpgradeTypeOnline  NsArrayUpgradeType = "online"
)

var pNsArrayUpgradeTypeOffline NsArrayUpgradeType
var pNsArrayUpgradeTypeInvalid NsArrayUpgradeType
var pNsArrayUpgradeTypeOnline NsArrayUpgradeType

// NsArrayUpgradeTypeOffline enum exports
var NsArrayUpgradeTypeOffline *NsArrayUpgradeType

// NsArrayUpgradeTypeInvalid enum exports
var NsArrayUpgradeTypeInvalid *NsArrayUpgradeType

// NsArrayUpgradeTypeOnline enum exports
var NsArrayUpgradeTypeOnline *NsArrayUpgradeType

func init() {
	pNsArrayUpgradeTypeOffline = cNsArrayUpgradeTypeOffline
	NsArrayUpgradeTypeOffline = &pNsArrayUpgradeTypeOffline

	pNsArrayUpgradeTypeInvalid = cNsArrayUpgradeTypeInvalid
	NsArrayUpgradeTypeInvalid = &pNsArrayUpgradeTypeInvalid

	pNsArrayUpgradeTypeOnline = cNsArrayUpgradeTypeOnline
	NsArrayUpgradeTypeOnline = &pNsArrayUpgradeTypeOnline

}
