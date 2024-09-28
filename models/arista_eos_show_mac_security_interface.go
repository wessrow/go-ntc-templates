package models

type AristaEosShowMacSecurityInterface struct {
	Interface	string	`json:"INTERFACE"`
	Sci	string	`json:"SCI"`
	Controlled_port	string	`json:"CONTROLLED_PORT"`
	Key_in_use	string	`json:"KEY_IN_USE"`
}