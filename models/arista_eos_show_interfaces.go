package models

type AristaEosShowInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Bia	string	`json:"BIA"`
	Description	string	`json:"DESCRIPTION"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mtu	string	`json:"MTU"`
	Bandwidth	string	`json:"BANDWIDTH"`
	Interface_up_time	string	`json:"INTERFACE_UP_TIME"`
	Link_status_change	string	`json:"LINK_STATUS_CHANGE"`
}