package models

type CiscoNxosShowInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Status	string	`json:"STATUS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Peer_ip_address	string	`json:"PEER_IP_ADDRESS"`
	Speed	string	`json:"SPEED"`
	Mtu	string	`json:"MTU"`
	Vlan_id	string	`json:"VLAN_ID"`
	Type	string	`json:"TYPE"`
	Mode	string	`json:"MODE"`
	Reason	string	`json:"REASON"`
	Portch	string	`json:"PORTCH"`
	Description	string	`json:"DESCRIPTION"`
	Protocol	string	`json:"PROTOCOL"`
	Vcid	string	`json:"VCID"`
}