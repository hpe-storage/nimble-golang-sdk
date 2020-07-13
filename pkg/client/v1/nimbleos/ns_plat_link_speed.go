// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPlatLinkSpeed Enum.

type NsPlatLinkSpeed string

const (
	cNsPlatLinkSpeedLinkSpeed10000m  NsPlatLinkSpeed = "link_speed_10000M"
	cNsPlatLinkSpeedLinkSpeed10m     NsPlatLinkSpeed = "link_speed_10M"
	cNsPlatLinkSpeedLinkSpeedUnknown NsPlatLinkSpeed = "link_speed_unknown"
	cNsPlatLinkSpeedLinkSpeed1000m   NsPlatLinkSpeed = "link_speed_1000M"
	cNsPlatLinkSpeedLinkSpeed100m    NsPlatLinkSpeed = "link_speed_100M"
)

var pNsPlatLinkSpeedLinkSpeed10000m NsPlatLinkSpeed
var pNsPlatLinkSpeedLinkSpeed10m NsPlatLinkSpeed
var pNsPlatLinkSpeedLinkSpeedUnknown NsPlatLinkSpeed
var pNsPlatLinkSpeedLinkSpeed1000m NsPlatLinkSpeed
var pNsPlatLinkSpeedLinkSpeed100m NsPlatLinkSpeed

// NsPlatLinkSpeedLinkSpeed10000m enum exports
var NsPlatLinkSpeedLinkSpeed10000m *NsPlatLinkSpeed

// NsPlatLinkSpeedLinkSpeed10m enum exports
var NsPlatLinkSpeedLinkSpeed10m *NsPlatLinkSpeed

// NsPlatLinkSpeedLinkSpeedUnknown enum exports
var NsPlatLinkSpeedLinkSpeedUnknown *NsPlatLinkSpeed

// NsPlatLinkSpeedLinkSpeed1000m enum exports
var NsPlatLinkSpeedLinkSpeed1000m *NsPlatLinkSpeed

// NsPlatLinkSpeedLinkSpeed100m enum exports
var NsPlatLinkSpeedLinkSpeed100m *NsPlatLinkSpeed

func init() {
	pNsPlatLinkSpeedLinkSpeed10000m = cNsPlatLinkSpeedLinkSpeed10000m
	NsPlatLinkSpeedLinkSpeed10000m = &pNsPlatLinkSpeedLinkSpeed10000m

	pNsPlatLinkSpeedLinkSpeed10m = cNsPlatLinkSpeedLinkSpeed10m
	NsPlatLinkSpeedLinkSpeed10m = &pNsPlatLinkSpeedLinkSpeed10m

	pNsPlatLinkSpeedLinkSpeedUnknown = cNsPlatLinkSpeedLinkSpeedUnknown
	NsPlatLinkSpeedLinkSpeedUnknown = &pNsPlatLinkSpeedLinkSpeedUnknown

	pNsPlatLinkSpeedLinkSpeed1000m = cNsPlatLinkSpeedLinkSpeed1000m
	NsPlatLinkSpeedLinkSpeed1000m = &pNsPlatLinkSpeedLinkSpeed1000m

	pNsPlatLinkSpeedLinkSpeed100m = cNsPlatLinkSpeedLinkSpeed100m
	NsPlatLinkSpeedLinkSpeed100m = &pNsPlatLinkSpeedLinkSpeed100m

}
