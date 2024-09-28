package models

type CiscoIosShowIsisNeighbors struct {
	System_id	string	`json:"SYSTEM_ID"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	State	string	`json:"STATE"`
	Hold_time	string	`json:"HOLD_TIME"`
	Circuit_id	string	`json:"CIRCUIT_ID"`
}