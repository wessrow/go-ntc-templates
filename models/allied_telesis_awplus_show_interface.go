package models

type AlliedTelesisAwplusShowInterface struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Vrf	string	`json:"VRF"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Description	string	`json:"DESCRIPTION"`
	Mtu	string	`json:"MTU"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
}