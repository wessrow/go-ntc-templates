package zyxel_os

type CfgIpaliasGet struct {
	Ip_alias_netmask  string `json:"IP_ALIAS_NETMASK"`
	Publan_active     string `json:"PUBLAN_ACTIVE"`
	Publan_ip_address string `json:"PUBLAN_IP_ADDRESS"`
	Publan_netmask    string `json:"PUBLAN_NETMASK"`
	Name              string `json:"NAME"`
	Ip_alias_active   string `json:"IP_ALIAS_ACTIVE"`
	Ip_alias_address  string `json:"IP_ALIAS_ADDRESS"`
}

var CfgIpaliasGet_Template string = `Value NAME (.+?)
Value IP_ALIAS_ACTIVE (0|1)
Value IP_ALIAS_ADDRESS (\S+)
Value IP_ALIAS_NETMASK (\S+)
Value PUBLAN_ACTIVE (0|1)
Value PUBLAN_IP_ADDRESS (\S+)
Value PUBLAN_NETMASK (\S+)

Start
  ^Group\sName\s+Active\s+IPv4\sArrdess\s+Subnet\smask\s+Active\s+IPv4\sArrdess\s+Subnet\smask\s+Offer\sPublic\sIP\sby\sDHCP\s+Enable\sARP\sProxy\s*$$ -> IPALIASTable

IPALIASTable
  ^${NAME}\s+${IP_ALIAS_ACTIVE}\s+${IP_ALIAS_ADDRESS}\s+${IP_ALIAS_NETMASK}\s+${PUBLAN_ACTIVE}\s+${PUBLAN_IP_ADDRESS}\s+${PUBLAN_NETMASK}\s+(0|1)\s+(0|1)\s*$$ -> Record
  ^Command\sSuccessful.\s*$$
  ^\s*$$
  ^. -> Error
`
