package models

type CheckpointGaiaShowIpv6Route struct {
	Protocol	string	`json:"PROTOCOL"`
	Network	string	`json:"NETWORK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Nexthopip	string	`json:"NEXTHOPIP"`
	Interface	string	`json:"INTERFACE"`
	Comment	string	`json:"COMMENT"`
}