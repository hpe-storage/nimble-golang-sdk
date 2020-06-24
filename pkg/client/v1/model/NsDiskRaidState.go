/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsDiskRaidState.</p>
 */

type NsDiskRaIDState string 

const (
	NSDISKRAIDSTATE_OKAY NsDiskRaIDState = "okay"
	NSDISKRAIDSTATE_N_A NsDiskRaIDState = "N/A"
	NSDISKRAIDSTATE_RESYNCHRONIZING NsDiskRaIDState = "resynchronizing"
	NSDISKRAIDSTATE_FAULTY NsDiskRaIDState = "faulty"
	NSDISKRAIDSTATE_SPARE NsDiskRaIDState = "spare"

) 
