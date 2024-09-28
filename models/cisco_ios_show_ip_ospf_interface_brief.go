package models

type CiscoIosShowIpOspfInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Process	string	`json:"PROCESS"`
	Area	string	`json:"AREA"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Cost	string	`json:"COST"`
	State	string	`json:"STATE"`
	Neighbors_fc	string	`json:"NEIGHBORS_FC"`
}