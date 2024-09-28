package models

type CiscoIosShowNveVni struct {
	Interface	string	`json:"INTERFACE"`
	Vni	string	`json:"VNI"`
	Mcast_group	string	`json:"MCAST_GROUP"`
	State	string	`json:"STATE"`
	Mode	string	`json:"MODE"`
	Bd	string	`json:"BD"`
	Vrf	string	`json:"VRF"`
}