package models

type CiscoIosShowIpVrfInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Vrf	string	`json:"VRF"`
	Proto_state	string	`json:"PROTO_STATE"`
}