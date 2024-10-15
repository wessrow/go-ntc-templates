package cisco_xr

type ShowOspfNeighbor struct {
	Dead_time       string `json:"DEAD_TIME"`
	Ip_address      string `json:"IP_ADDRESS"`
	Interface       string `json:"INTERFACE"`
	Neighbor_uptime string `json:"NEIGHBOR_UPTIME"`
	Neighbor_id     string `json:"NEIGHBOR_ID"`
	Priority        string `json:"PRIORITY"`
	State           string `json:"STATE"`
}

var ShowOspfNeighbor_Template string = `Value NEIGHBOR_ID (\d+.\d+.\d+.\d+)
Value PRIORITY (\d+)
Value STATE (\S+\/\s+\-|\S+)
Value DEAD_TIME (\d+:\d+:\d+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value INTERFACE (\S+)
Value NEIGHBOR_UPTIME (\S+)

Start
  ^${NEIGHBOR_ID}\s+${PRIORITY}\s+${STATE}\s+${DEAD_TIME}\s+${IP_ADDRESS}\s+${INTERFACE}
  ^\s+Neighbor is up for ${NEIGHBOR_UPTIME} -> Record
`
