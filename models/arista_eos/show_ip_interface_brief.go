package arista_eos

type ShowIpInterfaceBrief struct {
	Mtu        string `json:"MTU"`
	Interface  string `json:"INTERFACE"`
	Ip_address string `json:"IP_ADDRESS"`
	Status     string `json:"STATUS"`
	Protocol   string `json:"PROTOCOL"`
}

var ShowIpInterfaceBrief_Template string = `Value INTERFACE (\S+)
Value IP_ADDRESS (\S+)
Value STATUS (\S+)
Value PROTOCOL (\S+)
Value MTU (\d+)

Start
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${STATUS}\s+${PROTOCOL}\s+${MTU} -> Record
`
