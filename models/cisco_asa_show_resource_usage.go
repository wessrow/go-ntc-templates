package models

type CiscoAsaShowResourceUsage struct {
	Resource	string	`json:"RESOURCE"`
	Current	string	`json:"CURRENT"`
	Peak	string	`json:"PEAK"`
	Limit	string	`json:"LIMIT"`
	Denied	string	`json:"DENIED"`
	Context	string	`json:"CONTEXT"`
}