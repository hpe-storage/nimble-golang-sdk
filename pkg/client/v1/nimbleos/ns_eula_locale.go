// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsEulaLocale Enum.

type NsEulaLocale string

const (
	cNsEulaLocaleEn NsEulaLocale = "en"
)

var pNsEulaLocaleEn NsEulaLocale

// NsEulaLocaleEn enum exports
var NsEulaLocaleEn *NsEulaLocale

func init() {
	pNsEulaLocaleEn = cNsEulaLocaleEn
	NsEulaLocaleEn = &pNsEulaLocaleEn

}
