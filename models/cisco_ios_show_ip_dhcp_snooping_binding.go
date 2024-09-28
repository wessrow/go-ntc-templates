package models

type CiscoIosShowIpDhcpSnoopingBinding struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Vlan	string	`json:"VLAN"`
	Interface	string	`json:"INTERFACE"`
	Type	string	`json:"TYPE"`
	Lease	string	`json:"LEASE"`
}