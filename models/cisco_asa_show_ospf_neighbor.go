package models

type CiscoAsaShowOspfNeighbor struct {
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Priority	string	`json:"PRIORITY"`
	State	string	`json:"STATE"`
	Dead_time	string	`json:"DEAD_TIME"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
}