package cisco_nxos

type ShowIpDhcpRelayAddress struct {
	Vrf           string `json:"VRF"`
	Interface     string `json:"INTERFACE"`
	Relay_address string `json:"RELAY_ADDRESS"`
}

var ShowIpDhcpRelayAddress_Template string = `Value INTERFACE (\S+)
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
