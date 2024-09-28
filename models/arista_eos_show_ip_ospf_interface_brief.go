package models

type AristaEosShowIpOspfInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Instance	string	`json:"INSTANCE"`
	Vrf	string	`json:"VRF"`
	Area	string	`json:"AREA"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Cost	string	`json:"COST"`
	State	string	`json:"STATE"`
	Neighbors	string	`json:"NEIGHBORS"`
}