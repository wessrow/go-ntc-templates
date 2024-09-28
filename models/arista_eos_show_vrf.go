package models

type AristaEosShowVrf struct {
	Vrf	string	`json:"VRF"`
	Rd	string	`json:"RD"`
	Interfaces	[]string	`json:"INTERFACES"`
}