package models

type CiscoNxosShowIpInterfaceBrief struct {
	Vrf	string	`json:"VRF"`
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Link	string	`json:"LINK"`
	Proto	string	`json:"PROTO"`
}