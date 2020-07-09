// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsEulaFormat Enum.

type NsEulaFormat string

const (
	cNsEulaFormatHtml NsEulaFormat = "html"
	cNsEulaFormatText NsEulaFormat = "text"
)

var pNsEulaFormatHtml NsEulaFormat
var pNsEulaFormatText NsEulaFormat

// NsEulaFormatHtml enum exports
var NsEulaFormatHtml *NsEulaFormat

// NsEulaFormatText enum exports
var NsEulaFormatText *NsEulaFormat

func init() {
	pNsEulaFormatHtml = cNsEulaFormatHtml
	NsEulaFormatHtml = &pNsEulaFormatHtml

	pNsEulaFormatText = cNsEulaFormatText
	NsEulaFormatText = &pNsEulaFormatText

}
