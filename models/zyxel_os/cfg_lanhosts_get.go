package zyxel_os

type CfgLanhostsGet struct {
	Mac_address  string `json:"MAC_ADDRESS"`
	Source       string `json:"SOURCE"`
	Conn_type    string `json:"CONN_TYPE"`
	Hostname     string `json:"HOSTNAME"`
	Ip_address   string `json:"IP_ADDRESS"`
	Ipv6_address string `json:"IPV6_ADDRESS"`
}

var CfgLanhostsGet_Template string = `Value HOSTNAME (\S+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value IPV6_ADDRESS (\S+)
Value MAC_ADDRESS (\S+)
Value SOURCE (\S+)
Value CONN_TYPE (\S+)

Start
  ^Name\s+IP\sAddress\s+IPv6\sAddress\s+MAC\sAddress\s+Address\sSource\s+Connection\sType\s*$$ -> HOSTSTable
  ^\s*$$
  ^. -> Error

HOSTSTable
  ^${HOSTNAME}\s+${IP_ADDRESS}\s+${IPV6_ADDRESS}\s+${MAC_ADDRESS}\s+${SOURCE}\s+${CONN_TYPE}\s*$$ -> Record
  ^Command Successful.\s*$$
  ^\s*$$
  ^. -> Error
`
