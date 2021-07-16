// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPoolMigrate Enum.

type NsPoolMigrate string

const (
	cNsPoolMigrateIn   NsPoolMigrate = "in"
	cNsPoolMigrateNone NsPoolMigrate = "none"
	cNsPoolMigrateOut  NsPoolMigrate = "out"
)

var pNsPoolMigrateIn NsPoolMigrate
var pNsPoolMigrateNone NsPoolMigrate
var pNsPoolMigrateOut NsPoolMigrate

// NsPoolMigrateIn enum exports
var NsPoolMigrateIn *NsPoolMigrate

// NsPoolMigrateNone enum exports
var NsPoolMigrateNone *NsPoolMigrate

// NsPoolMigrateOut enum exports
var NsPoolMigrateOut *NsPoolMigrate

func init() {
	pNsPoolMigrateIn = cNsPoolMigrateIn
	NsPoolMigrateIn = &pNsPoolMigrateIn

	pNsPoolMigrateNone = cNsPoolMigrateNone
	NsPoolMigrateNone = &pNsPoolMigrateNone

	pNsPoolMigrateOut = cNsPoolMigrateOut
	NsPoolMigrateOut = &pNsPoolMigrateOut

}
