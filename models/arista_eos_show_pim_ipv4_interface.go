package models

type AristaEosShowPimIpv4Interface struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Mode	string	`json:"MODE"`
	Neighbors	string	`json:"NEIGHBORS"`
	Hello_interval	string	`json:"HELLO_INTERVAL"`
	Dr_priority	string	`json:"DR_PRIORITY"`
	Dr_address	string	`json:"DR_ADDRESS"`
	Packets_queued	string	`json:"PACKETS_QUEUED"`
	Packets_dropped	string	`json:"PACKETS_DROPPED"`
}