package cisco_xr

type ShowDhcpIpv4ProxyBinding struct {
	Ip_address      string `json:"IP_ADDRESS"`
	State           string `json:"STATE"`
	Lease_remaining string `json:"LEASE_REMAINING"`
	Mac_address     string `json:"MAC_ADDRESS"`
	Interface       string `json:"INTERFACE"`
	Vrf             string `json:"VRF"`
	Sublabel        string `json:"SUBLABEL"`
}

var ShowDhcpIpv4ProxyBinding_Template string = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value STATE (\S+)
Value LEASE_REMAINING (\S+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (\S+)
Value VRF (\S+)
Value SUBLABEL (\S+)

Start
  ^${MAC_ADDRESS}\s+${IP_ADDRESS}\s+${STATE}\s+${LEASE_REMAINING}\s+${INTERFACE}\s+${VRF}\s+${SUBLABEL} -> Record
`
