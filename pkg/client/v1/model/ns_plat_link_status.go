// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsPlatLinkStatus Enum.

type NsPlatLinkStatus string

const (
	cNsPlatLinkStatusLinkStatusDown    NsPlatLinkStatus = "link_status_down"
	cNsPlatLinkStatusLinkStatusUp      NsPlatLinkStatus = "link_status_up"
	cNsPlatLinkStatusLinkStatusUnknown NsPlatLinkStatus = "link_status_unknown"
)

var pNsPlatLinkStatusLinkStatusDown NsPlatLinkStatus
var pNsPlatLinkStatusLinkStatusUp NsPlatLinkStatus
var pNsPlatLinkStatusLinkStatusUnknown NsPlatLinkStatus

// Export
var NsPlatLinkStatusLinkStatusDown *NsPlatLinkStatus
var NsPlatLinkStatusLinkStatusUp *NsPlatLinkStatus
var NsPlatLinkStatusLinkStatusUnknown *NsPlatLinkStatus

func init() {
	pNsPlatLinkStatusLinkStatusDown = cNsPlatLinkStatusLinkStatusDown
	NsPlatLinkStatusLinkStatusDown = &pNsPlatLinkStatusLinkStatusDown

	pNsPlatLinkStatusLinkStatusUp = cNsPlatLinkStatusLinkStatusUp
	NsPlatLinkStatusLinkStatusUp = &pNsPlatLinkStatusLinkStatusUp

	pNsPlatLinkStatusLinkStatusUnknown = cNsPlatLinkStatusLinkStatusUnknown
	NsPlatLinkStatusLinkStatusUnknown = &pNsPlatLinkStatusLinkStatusUnknown

}
