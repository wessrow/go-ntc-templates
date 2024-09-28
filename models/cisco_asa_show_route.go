package models

type CiscoAsaShowRoute struct {
	Protocol	string	`json:"PROTOCOL"`
	Type	string	`json:"TYPE"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthopip	string	`json:"NEXTHOPIP"`
	Nexthopif	string	`json:"NEXTHOPIF"`
	Uptime	string	`json:"UPTIME"`
}