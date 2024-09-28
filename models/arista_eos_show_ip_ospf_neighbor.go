package models

type AristaEosShowIpOspfNeighbor struct {
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Instance	string	`json:"INSTANCE"`
	Vrf	string	`json:"VRF"`
	Priority	string	`json:"PRIORITY"`
	State	string	`json:"STATE"`
	Dead_time	string	`json:"DEAD_TIME"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
}