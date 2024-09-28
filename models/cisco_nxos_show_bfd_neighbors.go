package models

type CiscoNxosShowBfdNeighbors struct {
	Our_address	string	`json:"OUR_ADDRESS"`
	Neighbor_address	string	`json:"NEIGHBOR_ADDRESS"`
	Ld_rd	string	`json:"LD_RD"`
	Rh_rs	string	`json:"RH_RS"`
	Holddown	string	`json:"HOLDDOWN"`
	State	string	`json:"STATE"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Type	string	`json:"TYPE"`
}