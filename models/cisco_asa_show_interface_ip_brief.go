package models

type CiscoAsaShowInterfaceIpBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Proto	string	`json:"PROTO"`
}