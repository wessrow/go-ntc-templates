package models

type CiscoIosShowIpv6Route struct {
	Protocol	string	`json:"PROTOCOL"`
	Network	string	`json:"NETWORK"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Nexthop_if	string	`json:"NEXTHOP_IF"`
}