package models

type CiscoIosShowSwitchDetailStackPorts struct {
	Switch	string	`json:"SWITCH"`
	Status1	string	`json:"STATUS1"`
	Status2	string	`json:"STATUS2"`
	Neighbor1	string	`json:"NEIGHBOR1"`
	Neighbor2	string	`json:"NEIGHBOR2"`
}