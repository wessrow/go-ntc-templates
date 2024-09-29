package cisco_ios 

type ShowIpOspfNeighbor struct {
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Priority	string	`json:"PRIORITY"`
	State	string	`json:"STATE"`
	Dead_time	string	`json:"DEAD_TIME"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
}

var ShowIpOspfNeighbor_Template = `Value NEIGHBOR_ID (\d+.\d+.\d+.\d+)
Value PRIORITY (\d+)
Value STATE (\S+\/\s+\-|\S+)
Value DEAD_TIME (\d+:\d+:\d+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value INTERFACE (\S+)

Start
  ^${NEIGHBOR_ID}\s+${PRIORITY}\s+${STATE}\s+${DEAD_TIME}\s+${IP_ADDRESS}\s+${INTERFACE} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`