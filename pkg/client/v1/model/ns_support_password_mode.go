// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsSupportPasswordMode Enum.

type NsSupportPasswordMode string

const (
	cNsSupportPasswordModeCiphertext NsSupportPasswordMode = "ciphertext"
)

var pNsSupportPasswordModeCiphertext NsSupportPasswordMode

// Export
var NsSupportPasswordModeCiphertext *NsSupportPasswordMode

func init() {
	pNsSupportPasswordModeCiphertext = cNsSupportPasswordModeCiphertext
	NsSupportPasswordModeCiphertext = &pNsSupportPasswordModeCiphertext

}
