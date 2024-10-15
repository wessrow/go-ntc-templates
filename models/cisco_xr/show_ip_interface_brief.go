package cisco_xr

type ShowIpInterfaceBrief struct {
	Vrf        string `json:"VRF"`
	Interface  string `json:"INTERFACE"`
	Ip_address string `json:"IP_ADDRESS"`
	Status     string `json:"STATUS"`
	Proto      string `json:"PROTO"`
}

var ShowIpInterfaceBrief_Template string = `Value INTERFACE (.+?)
Value IP_ADDRESS (\S+)
Value STATUS (Up|Down|Shutdown)
Value PROTO (Up|Down)
Value VRF (\S+)

Start
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${STATUS}\s+${PROTO}\s+${VRF} -> Record
`
