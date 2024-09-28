package models

type CiscoNxosShowForwardingIpv4Route struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Nexthop	[]string	`json:"NEXTHOP"`
	Interface	[]string	`json:"INTERFACE"`
}