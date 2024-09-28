package models

type CiscoNxosShowIpv6InterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Linkipaddr	string	`json:"LINKIPADDR"`
	Status	string	`json:"STATUS"`
	Proto	string	`json:"PROTO"`
}