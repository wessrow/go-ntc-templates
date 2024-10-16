package ubiquiti_edgerouter

type ShowDhcpv6ServerLeases struct {
	Ip_address string `json:"IP_ADDRESS"`
	Expiration string `json:"EXPIRATION"`
	State      string `json:"STATE"`
}

var ShowDhcpv6ServerLeases_Template string = `Value Required IP_ADDRESS (\S+)
Value EXPIRATION (\d{4}(\/\d{2}){2}\s\d{2}(:\d{2}){2})
Value STATE (\S+)

Start
  ^IPv6\sAddress\s+Expiration\s+State\s*$$
  ^-+\s+-+\s+-+\s*$$ -> DHCPv6Table

DHCPv6Table
  ^${IP_ADDRESS}\s+${EXPIRATION}\s+${STATE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
