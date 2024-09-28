package models

type CiscoAsaShowObjectGroupNetwork struct {
	Name	string	`json:"NAME"`
	Description	string	`json:"DESCRIPTION"`
	Type	string	`json:"TYPE"`
	Host	string	`json:"HOST"`
	Net_object	string	`json:"NET_OBJECT"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Grp_object	string	`json:"GRP_OBJECT"`
}