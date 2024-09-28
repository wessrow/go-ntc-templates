package models

type CiscoNxosShowIpOspfInterfaceBrief struct {
	Process	string	`json:"PROCESS"`
	Vrf	string	`json:"VRF"`
	Interface_count	string	`json:"INTERFACE_COUNT"`
	Interface	string	`json:"INTERFACE"`
	Area	string	`json:"AREA"`
	Cost	string	`json:"COST"`
	State	string	`json:"STATE"`
	Neighbors_fc	string	`json:"NEIGHBORS_FC"`
	Status	string	`json:"STATUS"`
}