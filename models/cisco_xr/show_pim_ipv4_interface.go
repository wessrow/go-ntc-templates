package cisco_xr

type ShowPimIpv4Interface struct {
	Neighbors      string `json:"NEIGHBORS"`
	Hello_interval string `json:"HELLO_INTERVAL"`
	Dr_priority    string `json:"DR_PRIORITY"`
	Dr             string `json:"DR"`
	Vrf            string `json:"VRF"`
	Ip_address     string `json:"IP_ADDRESS"`
	Interface      string `json:"INTERFACE"`
	State          string `json:"STATE"`
}

var ShowPimIpv4Interface_Template string = `Value Filldown VRF (\S+)
Value Required IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value INTERFACE ([\w\./-]+)
Value STATE (\w+)
Value NEIGHBORS (\d+)
Value HELLO_INTERVAL (\d+)
Value DR_PRIORITY (\d+)
Value DR (\d+\.\d+\.\d+\.\d+)

Start
  ^\s*PIM interfaces in VRF ${VRF}\s*$$
  ^\s*Address\s+Interface\s+PIM\s+Nbr\s+Hello\s+DR\s+DR\s*$$
  ^\s*Count\s+Intvl\s+Prior\s*$$
  ^\s*${IP_ADDRESS}(\*)?\s+${INTERFACE}\s+${STATE}\s+${NEIGHBORS}\s+${HELLO_INTERVAL}\s+${DR_PRIORITY}\s+${DR}\s*$$ -> Record
`
