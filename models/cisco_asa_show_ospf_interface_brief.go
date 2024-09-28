package models

type CiscoAsaShowOspfInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Process	string	`json:"PROCESS"`
	Area	string	`json:"AREA"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Cost	string	`json:"COST"`
	State	string	`json:"STATE"`
	Neighbors_fc	string	`json:"NEIGHBORS_FC"`
}