// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsDiskRaidState Enum.</p>
 */

type NsDiskRaidState string 

const (
	NSDISKRAIDSTATE_OKAY NsDiskRaidState = "okay"
	NSDISKRAIDSTATE_N_A NsDiskRaidState = "N/A"
	NSDISKRAIDSTATE_RESYNCHRONIZING NsDiskRaidState = "resynchronizing"
	NSDISKRAIDSTATE_FAULTY NsDiskRaidState = "faulty"
	NSDISKRAIDSTATE_SPARE NsDiskRaidState = "spare"

) 
