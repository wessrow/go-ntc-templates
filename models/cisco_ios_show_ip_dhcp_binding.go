package models

type CiscoIosShowIpDhcpBinding struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Hardware_address	string	`json:"HARDWARE_ADDRESS"`
	Expiration	string	`json:"EXPIRATION"`
	Type	string	`json:"TYPE"`
}