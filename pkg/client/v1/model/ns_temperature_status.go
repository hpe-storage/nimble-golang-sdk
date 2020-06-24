// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsTemperatureStatus Enum.</p>
 */

type NsTemperatureStatus string 

const (
	NSTEMPERATURESTATUS_TEMPERATURE_UNKNOWN NsTemperatureStatus = "temperature_unknown"
	NSTEMPERATURESTATUS_TEMPERATURE_ALERTED NsTemperatureStatus = "temperature_alerted"
	NSTEMPERATURESTATUS_TEMPERATURE_OKAY NsTemperatureStatus = "temperature_okay"
	NSTEMPERATURESTATUS_TEMPERATURE_FAIL NsTemperatureStatus = "temperature_fail"

) 
