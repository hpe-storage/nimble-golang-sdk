// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsPlatOrientation Enum.

type NsPlatOrientation string

const (
	cNsPlatOrientationRightToLeft NsPlatOrientation = "right_to_left"
	cNsPlatOrientationLeftToRight NsPlatOrientation = "left_to_right"
)

var pNsPlatOrientationRightToLeft NsPlatOrientation
var pNsPlatOrientationLeftToRight NsPlatOrientation

// Export
var NsPlatOrientationRightToLeft *NsPlatOrientation
var NsPlatOrientationLeftToRight *NsPlatOrientation

func init() {
	pNsPlatOrientationRightToLeft = cNsPlatOrientationRightToLeft
	NsPlatOrientationRightToLeft = &pNsPlatOrientationRightToLeft

	pNsPlatOrientationLeftToRight = cNsPlatOrientationLeftToRight
	NsPlatOrientationLeftToRight = &pNsPlatOrientationLeftToRight

}
