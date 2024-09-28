package models

type CiscoIosShowVlans struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Interfaces	[]string	`json:"INTERFACES"`
	Ip_addresses	[]string	`json:"IP_ADDRESSES"`
	Ipv6_addresses	[]string	`json:"IPV6_ADDRESSES"`
	Inner_vlans	[]string	`json:"INNER_VLANS"`
}