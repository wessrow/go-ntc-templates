package models

type ArubaOsShowIpInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Admin	string	`json:"ADMIN"`
	Protocol	string	`json:"PROTOCOL"`
}