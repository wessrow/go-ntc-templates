package models

type CiscoNxosShowIpRoute struct {
	Vrf	string	`json:"VRF"`
	Protocol	string	`json:"PROTOCOL"`
	Type	string	`json:"TYPE"`
	Network	string	`json:"NETWORK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Nexthop_if	string	`json:"NEXTHOP_IF"`
	Uptime	string	`json:"UPTIME"`
	Nexthop_vrf	string	`json:"NEXTHOP_VRF"`
	Tag	string	`json:"TAG"`
	Segid	string	`json:"SEGID"`
	Tunnelid	string	`json:"TUNNELID"`
	Encap	string	`json:"ENCAP"`
}