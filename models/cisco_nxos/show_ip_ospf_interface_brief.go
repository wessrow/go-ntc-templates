package cisco_nxos

type ShowIpOspfInterfaceBrief struct {
	Process         string `json:"PROCESS"`
	Area            string `json:"AREA"`
	Neighbors_fc    string `json:"NEIGHBORS_FC"`
	Status          string `json:"STATUS"`
	Vrf             string `json:"VRF"`
	Interface_count string `json:"INTERFACE_COUNT"`
	Interface       string `json:"INTERFACE"`
	Cost            string `json:"COST"`
	State           string `json:"STATE"`
}

var ShowIpOspfInterfaceBrief_Template string = `Value Filldown PROCESS (\S+)
Value Filldown VRF (\S+)
Value Filldown INTERFACE_COUNT (\d+)
Value Required INTERFACE (\S+)
Value AREA (\S+)
Value COST (\d+)
Value STATE (\S+)
Value NEIGHBORS_FC (\d+)
Value STATUS (up|down)

Start
  ^\s+OSPF\s+Process\s+ID\s+${PROCESS}\s+VRF\s+${VRF}
  ^\s+Total\s+number\s+of\s+interface:\s+${INTERFACE_COUNT}
  ^\s+Interface\s+ID\s+Area\s+Cost\s+State\s+Neighbors\s+Status
  ^\s+${INTERFACE}\s+\d+\s+${AREA}\s+${COST}\s+${STATE}\s+${NEIGHBORS_FC}\s+${STATUS} -> Record
  ^\s*$$
  ^. -> Error
`
