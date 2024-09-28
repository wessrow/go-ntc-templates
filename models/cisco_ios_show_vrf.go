package models

type CiscoIosShowVrf struct {
	Name	string	`json:"NAME"`
	Default_rd	string	`json:"DEFAULT_RD"`
	Protocols	string	`json:"PROTOCOLS"`
	Interfaces	[]string	`json:"INTERFACES"`
}