package models

type AlliedTelesisAwplusShowIpRoute struct {
	Vrf	string	`json:"VRF"`
	Protocol	string	`json:"PROTOCOL"`
	Type	string	`json:"TYPE"`
	Network	string	`json:"NETWORK"`
	Mask	string	`json:"MASK"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Nexthop_vrf	string	`json:"NEXTHOP_VRF"`
	Nexthop_if	string	`json:"NEXTHOP_IF"`
	Uptime	string	`json:"UPTIME"`
}