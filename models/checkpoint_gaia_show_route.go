package models

type CheckpointGaiaShowRoute struct {
	Protocol	string	`json:"PROTOCOL"`
	Network	string	`json:"NETWORK"`
	Mask	string	`json:"MASK"`
	Nexthopip	string	`json:"NEXTHOPIP"`
	Interface	string	`json:"INTERFACE"`
	Comment	string	`json:"COMMENT"`
}