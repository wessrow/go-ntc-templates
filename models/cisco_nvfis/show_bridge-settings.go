package cisco_nvfis

type ShowBridgeSettings struct {
	Ipv6_slaac           string `json:"IPV6_SLAAC"`
	Interface            string `json:"INTERFACE"`
	Ip_address           string `json:"IP_ADDRESS"`
	Ipv6_address         string `json:"IPV6_ADDRESS"`
	Vlan_tag             string `json:"VLAN_TAG"`
	Netmask              string `json:"NETMASK"`
	Ipv6_prefixlen       string `json:"IPV6_PREFIXLEN"`
	Mac_address          string `json:"MAC_ADDRESS"`
	Dpdk                 string `json:"DPDK"`
	Name                 string `json:"NAME"`
	Mtu                  string `json:"MTU"`
	Secondary_ip_address string `json:"SECONDARY_IP_ADDRESS"`
	Secondary_ip_netmask string `json:"SECONDARY_IP_NETMASK"`
	Dhcp                 string `json:"DHCP"`
	Ipv6_dhcp            string `json:"IPV6_DHCP"`
}

var ShowBridgeSettings_Template string = `Value NAME (\S+)
Value INTERFACE (\S+)
Value IP_ADDRESS (\S+)
Value NETMASK (\S+)
Value IPV6_ADDRESS (\S+)
Value IPV6_PREFIXLEN (\S+)
Value MAC_ADDRESS (\S+)
Value MTU (\d+)
Value SECONDARY_IP_ADDRESS (\S+)
Value SECONDARY_IP_NETMASK (\S+)
Value VLAN_TAG (\S+)
Value DHCP (\S+)
Value IPV6_DHCP (\S+)
Value IPV6_SLAAC (\S+)
Value DPDK (\S+)


Start
  ^bridge-settings\s${NAME}
  ^\sip-info\sinterface\s+${INTERFACE}
  ^\sip-info\sipv4_address\s+${IP_ADDRESS}
  ^\sip-info\snetmask\s+${NETMASK}
  ^\sip-info\sglobal\sipv6\saddress\s+${IPV6_ADDRESS}
  ^\sip-info\sglobal\sipv6\sprefixlen\s+${IPV6_PREFIXLEN}
  ^\sip-info\smac_address\s+${MAC_ADDRESS}
  ^\sip-info\smtu\s+${MTU}
  ^\sip-info\ssecondary-ip\s+${SECONDARY_IP_ADDRESS}
  ^\sip-info\ssecondary-ip-netmask\s+${SECONDARY_IP_NETMASK}
  ^\svlan\stag\s${VLAN_TAG}
  ^\sdhcp\s${DHCP}
  ^\sdhcp-ipv6\s${IPV6_DHCP}
  ^\sslaac-ipv6\s${IPV6_SLAAC}
  ^\sdpdk-enabled\s${DPDK} -> Record

EOF`
