package zyxel_os

type CfgLanGetNameName struct {
	Name             string `json:"NAME"`
	Dhcp             string `json:"DHCP"`
	Dhcp_start_ip    string `json:"DHCP_START_IP"`
	Ipv6_active      string `json:"IPV6_ACTIVE"`
	Ip_address       string `json:"IP_ADDRESS"`
	Netmask          string `json:"NETMASK"`
	Dhcp_end_ip      string `json:"DHCP_END_IP"`
	Dhcp_autoreserve string `json:"DHCP_AUTORESERVE"`
	Dhcp_lease       string `json:"DHCP_LEASE"`
}

var CfgLanGetNameName_Template string = `Value NAME (\S.*)
Value IP_ADDRESS (\S*)
Value NETMASK (\S*)
Value DHCP (Enable|Disable)
Value DHCP_START_IP (\S*)
Value DHCP_END_IP (\S*)
Value DHCP_AUTORESERVE (true|false)
Value DHCP_LEASE (\S.*)
Value IPV6_ACTIVE (true|false)

Start
  ^Group\sName\s+${NAME}\s*$$
  ^IP\sAddress\s+${IP_ADDRESS}\s*$$
  ^Subnet\sMask\s+${NETMASK}\s*$$
  ^DHCP\s+${DHCP}\s*$$
  ^\s+Beginning\sIP\sAddress\s+${DHCP_START_IP}\s*$$
  ^\s+Ending\sIP\sAddress\s+${DHCP_END_IP}\s*$$
  ^\s+AutoReserveLanIp\s+${DHCP_AUTORESERVE}\s*$$
  ^\s+DHCP\sServer\sLease\sTime\s+${DHCP_LEASE}\s*$$
  ^\s+DNS\sValues\s.*$$
  ^IPv6\sActive\s+${IPV6_ACTIVE}\s*$$
  ^Command\sSuccessful.\s*$$
  ^\s*$$
  ^. -> Error
`
