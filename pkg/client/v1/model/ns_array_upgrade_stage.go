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

// Export
var NsArrayUpgradeStagePrepare *NsArrayUpgradeStage
var NsArrayUpgradeStageFinish *NsArrayUpgradeStage
var NsArrayUpgradeStageNone *NsArrayUpgradeStage

func init() {
	pNsArrayUpgradeStagePrepare = cNsArrayUpgradeStagePrepare
	NsArrayUpgradeStagePrepare = &pNsArrayUpgradeStagePrepare

	pNsArrayUpgradeStageFinish = cNsArrayUpgradeStageFinish
	NsArrayUpgradeStageFinish = &pNsArrayUpgradeStageFinish

	pNsArrayUpgradeStageNone = cNsArrayUpgradeStageNone
	NsArrayUpgradeStageNone = &pNsArrayUpgradeStageNone

}
