package models

type ZyxelOsCfgLanGet struct {
	Name	string	`json:"NAME"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Ipv4_dhcp	string	`json:"IPV4_DHCP"`
	Ipv6_enable	string	`json:"IPV6_ENABLE"`
}

var ZyxelOsCfgLanGet_Template = `Value NAME (.+?)
Value IP_ADDRESS ((\d+\.){3}\d+)
Value IPV4_DHCP (Enable|Disable)
Value IPV6_ENABLE (true|false)

Start
  ^Group\sName\s+LAN\sIP\sAddress\s+IPv4\sDHCP\sServer\s+IPv6\sEnable\s*$$ -> LANTable
  ^\s*$$
  ^. -> Error

LANTable
  ^${NAME}\s+${IP_ADDRESS}\s+${IPV4_DHCP}\s+${IPV6_ENABLE}\s*$$ -> Record
  ^${NAME}\s+${IPV4_DHCP}\s+${IPV6_ENABLE}\s*$$ -> Record
  ^Command\sSuccessful.\s*$$
  ^\s*$$
  ^. -> Error

`