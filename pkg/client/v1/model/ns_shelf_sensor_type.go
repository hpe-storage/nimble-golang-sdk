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

// Export
var NsShelfSensorTypeFan *NsShelfSensorType
var NsShelfSensorTypeNvram *NsShelfSensorType
var NsShelfSensorTypeTemperature *NsShelfSensorType
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
