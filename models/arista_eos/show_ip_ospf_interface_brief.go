package arista_eos

type ShowIpOspfInterfaceBrief struct {
	State         string `json:"STATE"`
	Neighbors     string `json:"NEIGHBORS"`
	Instance      string `json:"INSTANCE"`
	Vrf           string `json:"VRF"`
	Ip_address    string `json:"IP_ADDRESS"`
	Prefix_length string `json:"PREFIX_LENGTH"`
	Cost          string `json:"COST"`
	Interface     string `json:"INTERFACE"`
	Area          string `json:"AREA"`
}

var ShowIpOspfInterfaceBrief_Template string = `Value INTERFACE (\S+)
Value INSTANCE (\d+)
Value VRF (\S+)
Value AREA (\d+\.\d+\.\d+\.\d+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value COST (\d+)
Value STATE (\S+)
Value NEIGHBORS (\d+)

Start
  ^\s*Interface\s+Instance VRF\s+Area\s+IP Address\s+Cost\s+State\s+Nbrs\s*$$
  ^\s*${INTERFACE}\s+${INSTANCE}\s+${VRF}\s+${AREA}\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${COST}\s+${STATE}\s+${NEIGHBORS}\s*$$ -> Record
`
