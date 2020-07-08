// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsArrayUpgradeType Enum.

type NsArrayUpgradeType string

const (
	cNsArrayUpgradeTypeOffline NsArrayUpgradeType = "offline"
	cNsArrayUpgradeTypeInvalid NsArrayUpgradeType = "invalid"
)

var pNsArrayUpgradeTypeOffline NsArrayUpgradeType
var pNsArrayUpgradeTypeInvalid NsArrayUpgradeType

// Export
var NsArrayUpgradeTypeOffline *NsArrayUpgradeType
var NsArrayUpgradeTypeInvalid *NsArrayUpgradeType

func init() {
	pNsArrayUpgradeTypeOffline = cNsArrayUpgradeTypeOffline
	NsArrayUpgradeTypeOffline = &pNsArrayUpgradeTypeOffline

	pNsArrayUpgradeTypeInvalid = cNsArrayUpgradeTypeInvalid
	NsArrayUpgradeTypeInvalid = &pNsArrayUpgradeTypeInvalid

}
