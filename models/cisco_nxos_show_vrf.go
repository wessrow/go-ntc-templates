package models

type CiscoNxosShowVrf struct {
	Name	string	`json:"NAME"`
	Id	string	`json:"ID"`
	State	string	`json:"STATE"`
	Reason	string	`json:"REASON"`
}