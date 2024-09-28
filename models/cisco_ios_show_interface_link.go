package models

type CiscoIosShowInterfaceLink struct {
	Port	string	`json:"PORT"`
	Name	string	`json:"NAME"`
	Downtime	string	`json:"DOWNTIME"`
	Since	string	`json:"SINCE"`
	Uptime	string	`json:"UPTIME"`
}