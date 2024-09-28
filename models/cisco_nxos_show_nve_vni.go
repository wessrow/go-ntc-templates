package models

type CiscoNxosShowNveVni struct {
	Interface	string	`json:"INTERFACE"`
	Vni	string	`json:"VNI"`
	Mcast_group	string	`json:"MCAST_GROUP"`
	State	string	`json:"STATE"`
	Mode	string	`json:"MODE"`
	Type	string	`json:"TYPE"`
	Bd_vrf	string	`json:"BD_VRF"`
}