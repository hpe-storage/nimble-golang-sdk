// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsEulaPhase Enum.

type NsEulaPhase string

const (
	cNsEulaPhaseSoftware NsEulaPhase = "software"
	cNsEulaPhaseSetup    NsEulaPhase = "setup"
)

var pNsEulaPhaseSoftware NsEulaPhase
var pNsEulaPhaseSetup NsEulaPhase

// NsEulaPhaseSoftware enum exports
var NsEulaPhaseSoftware *NsEulaPhase

// NsEulaPhaseSetup enum exports
var NsEulaPhaseSetup *NsEulaPhase

func init() {
	pNsEulaPhaseSoftware = cNsEulaPhaseSoftware
	NsEulaPhaseSoftware = &pNsEulaPhaseSoftware

	pNsEulaPhaseSetup = cNsEulaPhaseSetup
	NsEulaPhaseSetup = &pNsEulaPhaseSetup

}
