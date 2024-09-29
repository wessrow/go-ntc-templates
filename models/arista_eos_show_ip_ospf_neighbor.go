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

var AristaEosShowIpOspfNeighbor_Template = `Value Required NEIGHBOR_ID (\d+.\d+.\d+.\d+)
Value INSTANCE (\d+)
Value VRF (\S+)
Value PRIORITY (\d+)
Value STATE (\S+)
Value DEAD_TIME (\d+:\d+:\d+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value INTERFACE ([\w\./-]+)

Start
  ^${NEIGHBOR_ID}\s+${VRF}\s+${PRIORITY}\s+${STATE}\s+${DEAD_TIME}\s+${IP_ADDRESS}\s+${INTERFACE} -> Record
  ^\s*${NEIGHBOR_ID}\s+${INSTANCE}\s+${VRF}\s+${PRIORITY}\s+${STATE}\s+${DEAD_TIME}\s+${IP_ADDRESS}\s+${INTERFACE}\s*$$ -> Record

`