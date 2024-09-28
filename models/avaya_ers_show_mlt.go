package models

type AvayaErsShowMlt struct {
	Id	string	`json:"ID"`
	Name	string	`json:"NAME"`
	Active_members	string	`json:"ACTIVE_MEMBERS"`
	Inactive_members	string	`json:"INACTIVE_MEMBERS"`
	Bpdu	string	`json:"BPDU"`
	Mode	string	`json:"MODE"`
	Status	string	`json:"STATUS"`
	Type	string	`json:"TYPE"`
	Lacp_key	string	`json:"LACP_KEY"`
}