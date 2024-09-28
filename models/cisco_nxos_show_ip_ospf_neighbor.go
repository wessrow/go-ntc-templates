package models

type CiscoNxosShowIpOspfNeighbor struct {
	Ospf_pid	string	`json:"OSPF_PID"`
	Vrf	string	`json:"VRF"`
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	State	string	`json:"STATE"`
	Uptime	string	`json:"UPTIME"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
}