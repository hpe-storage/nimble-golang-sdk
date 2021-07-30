// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFqdnOrIpAndPortObject - Object wrapper of Fqdn/IP Address and Port numbers.

// Export NsFqdnOrIpAndPortObjectFields provides field names to use in filter parameters, for example.
var NsFqdnOrIpAndPortObjectFields *NsFqdnOrIpAndPortObjectStringFields

func init() {
	fieldSyslogServer := "syslog_server"
	fieldSyslogPort := "syslog_port"

	NsFqdnOrIpAndPortObjectFields = &NsFqdnOrIpAndPortObjectStringFields{
		SyslogServer: &fieldSyslogServer,
		SyslogPort:   &fieldSyslogPort,
	}
}

type NsFqdnOrIpAndPortObject struct {
	// SyslogServer - A Fqdn/IP Address.
	SyslogServer *string `json:"syslog_server,omitempty"`
	// SyslogPort - Syslog port number.
	SyslogPort *int64 `json:"syslog_port,omitempty"`
}

// Struct for NsFqdnOrIpAndPortObjectFields
type NsFqdnOrIpAndPortObjectStringFields struct {
	SyslogServer *string
	SyslogPort   *string
}
