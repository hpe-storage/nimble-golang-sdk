// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsSensorState Enum.</p>
 */

type NsSensorState string 

const (
	NSSENSORSTATE_SENSOR_OK NsSensorState = "sensor_ok"
	NSSENSORSTATE_SENSOR_ALERT_COND NsSensorState = "sensor_alert_cond"
	NSSENSORSTATE_SENSOR_MISSING NsSensorState = "sensor_missing"
	NSSENSORSTATE_SENSOR_READING_UNAVAIL NsSensorState = "sensor_reading_unavail"
	NSSENSORSTATE_SENSOR_FAILED NsSensorState = "sensor_failed"

) 
