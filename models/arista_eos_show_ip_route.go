package models

type AristaEosShowIpRoute struct {
	Vrf	string	`json:"VRF"`
	Protocol	string	`json:"PROTOCOL"`
	Network	string	`json:"NETWORK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Direct	string	`json:"DIRECT"`
	Next_hop	[]string	`json:"NEXT_HOP"`
	Interface	[]string	`json:"INTERFACE"`
}