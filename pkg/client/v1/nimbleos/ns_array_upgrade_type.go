// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsArrayUpgradeType Enum.

type NsArrayUpgradeType string

const (
	cNsArrayUpgradeTypeOffline NsArrayUpgradeType = "offline"
	cNsArrayUpgradeTypeInvalid NsArrayUpgradeType = "invalid"
)

var pNsArrayUpgradeTypeOffline NsArrayUpgradeType
var pNsArrayUpgradeTypeInvalid NsArrayUpgradeType

// NsArrayUpgradeTypeOffline enum exports
var NsArrayUpgradeTypeOffline *NsArrayUpgradeType

// NsArrayUpgradeTypeInvalid enum exports
var NsArrayUpgradeTypeInvalid *NsArrayUpgradeType

func init() {
	pNsArrayUpgradeTypeOffline = cNsArrayUpgradeTypeOffline
	NsArrayUpgradeTypeOffline = &pNsArrayUpgradeTypeOffline

	pNsArrayUpgradeTypeInvalid = cNsArrayUpgradeTypeInvalid
	NsArrayUpgradeTypeInvalid = &pNsArrayUpgradeTypeInvalid

}
