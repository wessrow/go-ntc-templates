package cisco_nxos 

type ShowIpDhcpRelayAddress struct {
	Interface	string	`json:"INTERFACE"`
	Relay_address	string	`json:"RELAY_ADDRESS"`
	Vrf	string	`json:"VRF"`
}

var ShowIpDhcpRelayAddress_Template = `Value INTERFACE (\S+)
Value RELAY_ADDRESS (\d+.\d+.\d+.\d+)
Value VRF (\S+)

Start
  ^\s*Interface\s+Relay\s+Address\s+VRF\s+Name\s*$$
  ^\s*-+
  ^\s*${INTERFACE}\s+${RELAY_ADDRESS}\s*$$ -> Record
  ^\s*${INTERFACE}\s+${RELAY_ADDRESS}\s+${VRF}\s*$$ -> Record
  ^\s*$$
  ^.* -> Error
`