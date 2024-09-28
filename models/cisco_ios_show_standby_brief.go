package models

type CiscoIosShowStandbyBrief struct {
	Interface	string	`json:"INTERFACE"`
	Group	string	`json:"GROUP"`
	Priority	string	`json:"PRIORITY"`
	Preempt	string	`json:"PREEMPT"`
	State	string	`json:"STATE"`
	Active	string	`json:"ACTIVE"`
	Standby	string	`json:"STANDBY"`
	Virtual_ip_address	string	`json:"VIRTUAL_IP_ADDRESS"`
}