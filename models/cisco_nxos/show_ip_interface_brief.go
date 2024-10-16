package cisco_nxos

type ShowIpInterfaceBrief struct {
	Interface  string `json:"INTERFACE"`
	Ip_address string `json:"IP_ADDRESS"`
	Status     string `json:"STATUS"`
	Link       string `json:"LINK"`
	Proto      string `json:"PROTO"`
	Vrf        string `json:"VRF"`
}

var ShowIpInterfaceBrief_Template string = `Value Filldown VRF (\S+)
Value Required INTERFACE (\S+)
Value Required IP_ADDRESS ([a-zA-Z0-9./]+|forward-enabled)
Value STATUS (\S+-\S+)
Value LINK (\S+-\S+)
Value PROTO (\S+-\S+)

Start
  ^IP\s+Interface\s+Status\s+for\s+VRF\s+"${VRF}"\(\d+\)
  ^Interface\s+IP\s+Address\s+Interface\s+Status
  ^\S+\s+unnumbered\s+\S+
  ^\s+\(\w+\)
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${PROTO}/${LINK}/${STATUS} -> Record
  ^\s*$$
  ^.*$$ -> Error
`
