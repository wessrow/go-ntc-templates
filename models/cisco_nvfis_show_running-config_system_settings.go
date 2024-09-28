package models

type CiscoNvfisShowRunningConfigSystemSettings struct {
	Hostname	string	`json:"HOSTNAME"`
	Dpdk	string	`json:"DPDK"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Default_gateway	string	`json:"DEFAULT_GATEWAY"`
}