package models

type AristaEosShowIsisNeighbors struct {
	Instance	string	`json:"INSTANCE"`
	Vrf	string	`json:"VRF"`
	System_id	string	`json:"SYSTEM_ID"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
	Snpa	string	`json:"SNPA"`
	State	string	`json:"STATE"`
	Hold_time	string	`json:"HOLD_TIME"`
	Circuit_id	string	`json:"CIRCUIT_ID"`
}