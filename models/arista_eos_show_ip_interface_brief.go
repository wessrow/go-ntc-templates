package models

type AristaEosShowIpInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Protocol	string	`json:"PROTOCOL"`
	Mtu	string	`json:"MTU"`
}