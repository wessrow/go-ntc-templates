package models

type CiscoIosShowIpv6InterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ipv6_address	[]string	`json:"IPV6_ADDRESS"`
	Admin	string	`json:"ADMIN"`
	Protocol	string	`json:"PROTOCOL"`
}