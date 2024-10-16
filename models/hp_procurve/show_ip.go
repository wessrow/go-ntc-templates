package hp_procurve

type ShowIp struct {
	Ip_address  string `json:"IP_ADDRESS"`
	Subnet_mask string `json:"SUBNET_MASK"`
	Proxy       string `json:"PROXY"`
	Local       string `json:"LOCAL"`
	Vlan_name   string `json:"VLAN_NAME"`
	Config      string `json:"CONFIG"`
}

var ShowIp_Template string = `Value VLAN_NAME (.+?)
Value CONFIG (\S+)
Value IP_ADDRESS (\S+)
Value SUBNET_MASK (\S+)
Value PROXY (\S+)
Value LOCAL (\S+)

Start
  ^\s*Internet\s+\(IP\)
  ^\s*IP\s+Routing
  ^\s*Default\s+Gateway
  ^\s*Default\s+TT
  ^\s*Arp\s+Age
  ^\s*Domain\s+Suffix
  ^\s*DNS\s+server
  ^\s+\|\s+Proxy
  ^\s*VLAN\s+ -> Interface
  ^\s*$$
  ^. -> Error

Interface
  ^\s*----
  ^\s*${VLAN_NAME}\s+\|\s+${CONFIG}\s+${IP_ADDRESS}\s+${SUBNET_MASK}\s+${PROXY}\s+${LOCAL}\s*$$ -> Record
  ^\s*${VLAN_NAME}\s+\|\s+${CONFIG}\s+${IP_ADDRESS}\s+${SUBNET_MASK} -> Record
  ^\s*${VLAN_NAME}\s+\|\s+${CONFIG}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
