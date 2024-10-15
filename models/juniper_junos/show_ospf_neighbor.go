package juniper_junos

type ShowOspfNeighbor struct {
	Priority    string `json:"PRIORITY"`
	Dead_time   string `json:"DEAD_TIME"`
	Ip_address  string `json:"IP_ADDRESS"`
	Interface   string `json:"INTERFACE"`
	State       string `json:"STATE"`
	Neighbor_id string `json:"NEIGHBOR_ID"`
}

var ShowOspfNeighbor_Template string = `Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value INTERFACE (\S+)
Value STATE (\S+)
Value NEIGHBOR_ID (\d+.\d+.\d+.\d+)
Value PRIORITY (\d+)
Value DEAD_TIME (\d+)


Start
  ^${IP_ADDRESS}\s+${INTERFACE}\s+${STATE}\s+${NEIGHBOR_ID}\s+${PRIORITY}\s+${DEAD_TIME} -> Record
  ^\s*$$
  ^{master:\d+}
`
