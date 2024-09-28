package models

type CiscoAsaShowRunningConfigObjectNetwork struct {
	Name	string	`json:"NAME"`
	Description	string	`json:"DESCRIPTION"`
	Type	string	`json:"TYPE"`
	Host	string	`json:"HOST"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Start_ip	string	`json:"START_IP"`
	End_ip	string	`json:"END_IP"`
	Fqdn	string	`json:"FQDN"`
}