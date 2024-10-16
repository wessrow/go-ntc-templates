package cisco_ios

type ShowIpDhcpSnoopingBinding struct {
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Vlan        string `json:"VLAN"`
	Interface   string `json:"INTERFACE"`
	Type        string `json:"TYPE"`
	Lease       string `json:"LEASE"`
}

var ShowIpDhcpSnoopingBinding_Template string = `Value Required IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3})
Value Required MAC_ADDRESS ([0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2})
Value Required VLAN ([0-9]+)
Value Required INTERFACE (\S+)
Value Required TYPE (\S+)
Value Required LEASE (\S+)

Start
  ^\s*M[Aa][Cc]\s*Address\s+I[Pp]\s*Address\s+Lease\(sec\)\s+Type\s+VLAN\s+Interface\s*$$
  ^\s*-+\s+-+
  ^\s*${MAC_ADDRESS}\s+${IP_ADDRESS}\s+${LEASE}\s+${TYPE}\s+${VLAN}\s+${INTERFACE} -> Record
  ^\s*Total\s+number\s+of\s+bindings
  ^\s*$$
  ^. -> Error
`
