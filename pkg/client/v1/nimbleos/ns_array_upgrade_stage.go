// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsArrayUpgradeStage Enum.

type NsArrayUpgradeStage string

const (
	cNsArrayUpgradeStagePrepare         NsArrayUpgradeStage = "prepare"
	cNsArrayUpgradeStageMixedGeneration NsArrayUpgradeStage = "mixed_generation"
	cNsArrayUpgradeStageFinalUpgrade    NsArrayUpgradeStage = "final_upgrade"
	cNsArrayUpgradeStageDisableMirror   NsArrayUpgradeStage = "disable_mirror"
	cNsArrayUpgradeStageDriveMigration  NsArrayUpgradeStage = "drive_migration"
	cNsArrayUpgradeStageEnableMirror    NsArrayUpgradeStage = "enable_mirror"
	cNsArrayUpgradeStageFinish          NsArrayUpgradeStage = "finish"
	cNsArrayUpgradeStageNone            NsArrayUpgradeStage = "none"
	cNsArrayUpgradeStageNewGeneration   NsArrayUpgradeStage = "new_generation"
)

var pNsArrayUpgradeStagePrepare NsArrayUpgradeStage
var pNsArrayUpgradeStageMixedGeneration NsArrayUpgradeStage
var pNsArrayUpgradeStageFinalUpgrade NsArrayUpgradeStage
var pNsArrayUpgradeStageDisableMirror NsArrayUpgradeStage
var pNsArrayUpgradeStageDriveMigration NsArrayUpgradeStage
var pNsArrayUpgradeStageEnableMirror NsArrayUpgradeStage
var pNsArrayUpgradeStageFinish NsArrayUpgradeStage
var pNsArrayUpgradeStageNone NsArrayUpgradeStage
var pNsArrayUpgradeStageNewGeneration NsArrayUpgradeStage

// NsArrayUpgradeStagePrepare enum exports
var NsArrayUpgradeStagePrepare *NsArrayUpgradeStage

// NsArrayUpgradeStageMixedGeneration enum exports
var NsArrayUpgradeStageMixedGeneration *NsArrayUpgradeStage

// NsArrayUpgradeStageFinalUpgrade enum exports
var NsArrayUpgradeStageFinalUpgrade *NsArrayUpgradeStage

// NsArrayUpgradeStageDisableMirror enum exports
var NsArrayUpgradeStageDisableMirror *NsArrayUpgradeStage

// NsArrayUpgradeStageDriveMigration enum exports
var NsArrayUpgradeStageDriveMigration *NsArrayUpgradeStage

// NsArrayUpgradeStageEnableMirror enum exports
var NsArrayUpgradeStageEnableMirror *NsArrayUpgradeStage

// NsArrayUpgradeStageFinish enum exports
var NsArrayUpgradeStageFinish *NsArrayUpgradeStage

// NsArrayUpgradeStageNone enum exports
var NsArrayUpgradeStageNone *NsArrayUpgradeStage

// NsArrayUpgradeStageNewGeneration enum exports
var NsArrayUpgradeStageNewGeneration *NsArrayUpgradeStage

func init() {
	pNsArrayUpgradeStagePrepare = cNsArrayUpgradeStagePrepare
	NsArrayUpgradeStagePrepare = &pNsArrayUpgradeStagePrepare

	pNsArrayUpgradeStageMixedGeneration = cNsArrayUpgradeStageMixedGeneration
	NsArrayUpgradeStageMixedGeneration = &pNsArrayUpgradeStageMixedGeneration

	pNsArrayUpgradeStageFinalUpgrade = cNsArrayUpgradeStageFinalUpgrade
	NsArrayUpgradeStageFinalUpgrade = &pNsArrayUpgradeStageFinalUpgrade

	pNsArrayUpgradeStageDisableMirror = cNsArrayUpgradeStageDisableMirror
	NsArrayUpgradeStageDisableMirror = &pNsArrayUpgradeStageDisableMirror

	pNsArrayUpgradeStageDriveMigration = cNsArrayUpgradeStageDriveMigration
	NsArrayUpgradeStageDriveMigration = &pNsArrayUpgradeStageDriveMigration

	pNsArrayUpgradeStageEnableMirror = cNsArrayUpgradeStageEnableMirror
	NsArrayUpgradeStageEnableMirror = &pNsArrayUpgradeStageEnableMirror

	pNsArrayUpgradeStageFinish = cNsArrayUpgradeStageFinish
	NsArrayUpgradeStageFinish = &pNsArrayUpgradeStageFinish

	pNsArrayUpgradeStageNone = cNsArrayUpgradeStageNone
	NsArrayUpgradeStageNone = &pNsArrayUpgradeStageNone

	pNsArrayUpgradeStageNewGeneration = cNsArrayUpgradeStageNewGeneration
	NsArrayUpgradeStageNewGeneration = &pNsArrayUpgradeStageNewGeneration

}
