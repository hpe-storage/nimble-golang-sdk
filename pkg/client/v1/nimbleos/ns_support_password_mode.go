// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsSupportPasswordMode Enum.

type NsSupportPasswordMode string

const (
	cNsSupportPasswordModeCiphertext NsSupportPasswordMode = "ciphertext"
)

var pNsSupportPasswordModeCiphertext NsSupportPasswordMode

// NsSupportPasswordModeCiphertext enum exports
var NsSupportPasswordModeCiphertext *NsSupportPasswordMode

func init() {
	pNsSupportPasswordModeCiphertext = cNsSupportPasswordModeCiphertext
	NsSupportPasswordModeCiphertext = &pNsSupportPasswordModeCiphertext

}
