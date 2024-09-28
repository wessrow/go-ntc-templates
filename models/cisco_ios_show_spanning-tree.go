package models

type CiscoIosShowSpanningTree struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Interface	string	`json:"INTERFACE"`
	Role	string	`json:"ROLE"`
	Status	string	`json:"STATUS"`
	Cost	string	`json:"COST"`
	Port_priority	string	`json:"PORT_PRIORITY"`
	Port_id	string	`json:"PORT_ID"`
	Type	string	`json:"TYPE"`
}