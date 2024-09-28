package models

type CiscoIosShowRepTopology struct {
	Switch	string	`json:"SWITCH"`
	Interface	string	`json:"INTERFACE"`
	Edge	string	`json:"EDGE"`
	Role	string	`json:"ROLE"`
}