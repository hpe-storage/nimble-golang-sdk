// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsShelfSensorType Enum.

type NsShelfSensorType string

const (
	cNsShelfSensorTypeFan         NsShelfSensorType = "fan"
	cNsShelfSensorTypeNvram       NsShelfSensorType = "nvram"
	cNsShelfSensorTypeTemperature NsShelfSensorType = "temperature"
	cNsShelfSensorTypePowerSupply NsShelfSensorType = "power supply"
)

var pNsShelfSensorTypeFan NsShelfSensorType
var pNsShelfSensorTypeNvram NsShelfSensorType
var pNsShelfSensorTypeTemperature NsShelfSensorType
var pNsShelfSensorTypePowerSupply NsShelfSensorType

// NsShelfSensorTypeFan enum exports
var NsShelfSensorTypeFan *NsShelfSensorType

// NsShelfSensorTypeNvram enum exports
var NsShelfSensorTypeNvram *NsShelfSensorType

// NsShelfSensorTypeTemperature enum exports
var NsShelfSensorTypeTemperature *NsShelfSensorType

// NsShelfSensorTypePowerSupply enum exports
var NsShelfSensorTypePowerSupply *NsShelfSensorType

func init() {
	pNsShelfSensorTypeFan = cNsShelfSensorTypeFan
	NsShelfSensorTypeFan = &pNsShelfSensorTypeFan

	pNsShelfSensorTypeNvram = cNsShelfSensorTypeNvram
	NsShelfSensorTypeNvram = &pNsShelfSensorTypeNvram

	pNsShelfSensorTypeTemperature = cNsShelfSensorTypeTemperature
	NsShelfSensorTypeTemperature = &pNsShelfSensorTypeTemperature

	pNsShelfSensorTypePowerSupply = cNsShelfSensorTypePowerSupply
	NsShelfSensorTypePowerSupply = &pNsShelfSensorTypePowerSupply

}
