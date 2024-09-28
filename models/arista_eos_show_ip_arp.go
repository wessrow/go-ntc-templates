package models

type AristaEosShowIpArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
}