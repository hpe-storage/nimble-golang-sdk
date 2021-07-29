// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPlatOrientation Enum.

type NsPlatOrientation string

const (
	cNsPlatOrientationRightToLeft NsPlatOrientation = "right_to_left"
	cNsPlatOrientationLeftToRight NsPlatOrientation = "left_to_right"
)

var pNsPlatOrientationRightToLeft NsPlatOrientation
var pNsPlatOrientationLeftToRight NsPlatOrientation

// NsPlatOrientationRightToLeft enum exports
var NsPlatOrientationRightToLeft *NsPlatOrientation

// NsPlatOrientationLeftToRight enum exports
var NsPlatOrientationLeftToRight *NsPlatOrientation

func init() {
	pNsPlatOrientationRightToLeft = cNsPlatOrientationRightToLeft
	NsPlatOrientationRightToLeft = &pNsPlatOrientationRightToLeft

	pNsPlatOrientationLeftToRight = cNsPlatOrientationLeftToRight
	NsPlatOrientationLeftToRight = &pNsPlatOrientationLeftToRight

}
