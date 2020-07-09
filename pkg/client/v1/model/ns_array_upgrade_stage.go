// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsArrayUpgradeStage Enum.

type NsArrayUpgradeStage string

const (
	cNsArrayUpgradeStagePrepare NsArrayUpgradeStage = "prepare"
	cNsArrayUpgradeStageFinish  NsArrayUpgradeStage = "finish"
	cNsArrayUpgradeStageNone    NsArrayUpgradeStage = "none"
)

var pNsArrayUpgradeStagePrepare NsArrayUpgradeStage
var pNsArrayUpgradeStageFinish NsArrayUpgradeStage
var pNsArrayUpgradeStageNone NsArrayUpgradeStage

// NsArrayUpgradeStagePrepare enum exports
var NsArrayUpgradeStagePrepare *NsArrayUpgradeStage

// NsArrayUpgradeStageFinish enum exports
var NsArrayUpgradeStageFinish *NsArrayUpgradeStage

// NsArrayUpgradeStageNone enum exports
var NsArrayUpgradeStageNone *NsArrayUpgradeStage

func init() {
	pNsArrayUpgradeStagePrepare = cNsArrayUpgradeStagePrepare
	NsArrayUpgradeStagePrepare = &pNsArrayUpgradeStagePrepare

	pNsArrayUpgradeStageFinish = cNsArrayUpgradeStageFinish
	NsArrayUpgradeStageFinish = &pNsArrayUpgradeStageFinish

	pNsArrayUpgradeStageNone = cNsArrayUpgradeStageNone
	NsArrayUpgradeStageNone = &pNsArrayUpgradeStageNone

}
