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

var CiscoNxosShowIpOspfNeighbor_Template = `Value Filldown OSPF_PID (\S+)
Value Filldown VRF (\S+)
Value NEIGHBOR_ID (\d+\.\d+\.\d+\.\d+)
Value STATE (\S+.\/.\S+)
Value UPTIME (\S+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required INTERFACE (\S+)

Start
  ^\s+\w+\s+\w+\s+\w+\s+${OSPF_PID}\s+[Vv][Rr][Ff]\s+${VRF} -> Record
  ^\s+${NEIGHBOR_ID}\s+\d+\s+${STATE}\s+${UPTIME}\s+${IP_ADDRESS}\s+${INTERFACE} -> Record

`