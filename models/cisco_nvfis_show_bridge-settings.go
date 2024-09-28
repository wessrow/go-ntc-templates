package models

type CiscoNvfisShowBridgeSettings struct {
	Name	string	`json:"NAME"`
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Ipv6_prefixlen	string	`json:"IPV6_PREFIXLEN"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Mtu	string	`json:"MTU"`
	Secondary_ip_address	string	`json:"SECONDARY_IP_ADDRESS"`
	Secondary_ip_netmask	string	`json:"SECONDARY_IP_NETMASK"`
	Vlan_tag	string	`json:"VLAN_TAG"`
	Dhcp	string	`json:"DHCP"`
	Ipv6_dhcp	string	`json:"IPV6_DHCP"`
	Ipv6_slaac	string	`json:"IPV6_SLAAC"`
	Dpdk	string	`json:"DPDK"`
}