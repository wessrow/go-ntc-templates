package models

type CiscoAsaShowOspfNeighbor struct {
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Priority	string	`json:"PRIORITY"`
	State	string	`json:"STATE"`
	Dead_time	string	`json:"DEAD_TIME"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
}

var CiscoAsaShowOspfNeighbor_Template = `Value NEIGHBOR_ID (\d+\.\d+\.\d+\.\d+)
Value PRIORITY (\S+)
Value STATE (\S+\/\s+\-|\S+)
Value DEAD_TIME (\d+:\d+:\d+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value INTERFACE (\S+)

Start
  ^Neighbor\s+ID\s+Pri\s+State\s+Dead\s+Time\s+Address\s+Interface\s*$$
  ^${NEIGHBOR_ID}\s+${PRIORITY}\s+${STATE}\s+${DEAD_TIME}\s+${IP_ADDRESS}\s+${INTERFACE} -> Record
  ^\s*$$
  ^. -> Error

`