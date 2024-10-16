package cisco_ios

type ShowIpDhcpBinding struct {
	Hardware_address string `json:"HARDWARE_ADDRESS"`
	Expiration       string `json:"EXPIRATION"`
	Type             string `json:"TYPE"`
	Ip_address       string `json:"IP_ADDRESS"`
}

var ShowIpDhcpBinding_Template string = `Value IP_ADDRESS (\S+)
Value HARDWARE_ADDRESS (\S+)
Value EXPIRATION (\S+)
Value TYPE (\S+)

Start
  ^\s*Bindings from all pools not associated with VRF:\s*$$
  ^\s*IP\s+address\s+Client-ID/(?:Hardware\s+address/(?:User\s+name)?)?\s+Lease expiration\s+Type\s*$$ -> DhcpTable
  ^\s*$$
  ^. -> Error

DhcpTable
  ^\s*Hardware\s+address/\s*$$
  ^\s*User\s+name\s*$$
  ^\s*${IP_ADDRESS}\s+${HARDWARE_ADDRESS}\s+${EXPIRATION}\s+${TYPE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
