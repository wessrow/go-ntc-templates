package cisco_xr

type ShowOspfVrfAllNeighbor struct {
	Vrf             string `json:"VRF"`
	State           string `json:"STATE"`
	Dead_time       string `json:"DEAD_TIME"`
	Interface       string `json:"INTERFACE"`
	Process         string `json:"PROCESS"`
	Neighbor_id     string `json:"NEIGHBOR_ID"`
	Priority        string `json:"PRIORITY"`
	Ip_address      string `json:"IP_ADDRESS"`
	Neighbor_uptime string `json:"NEIGHBOR_UPTIME"`
}

var ShowOspfVrfAllNeighbor_Template string = `Value Filldown PROCESS (\d+)
Value Filldown VRF (\S+)
Value Required NEIGHBOR_ID (\d+.\d+.\d+.\d+)
Value PRIORITY (\d+)
Value STATE (\S+\/\s+\-|\S+)
Value DEAD_TIME (\d+:\d+:\d+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value INTERFACE (\S+)
Value NEIGHBOR_UPTIME (\S+)

Start
  ^\s*Neighbors for OSPF ${PROCESS}, VRF ${VRF}\s*$$
  ^\s*${NEIGHBOR_ID}\s+${PRIORITY}\s+${STATE}\s+${DEAD_TIME}\s+${IP_ADDRESS}\s+${INTERFACE}\s*$$
  ^\s*Neighbor is up for ${NEIGHBOR_UPTIME}\s*$$ -> Record
`
