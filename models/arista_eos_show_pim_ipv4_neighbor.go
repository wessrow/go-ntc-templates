package models

type AristaEosShowPimIpv4Neighbor struct {
	Vrf	string	`json:"VRF"`
	Neighbor	string	`json:"NEIGHBOR"`
	Interface	string	`json:"INTERFACE"`
	Uptime	string	`json:"UPTIME"`
	Expires	string	`json:"EXPIRES"`
	Mode	string	`json:"MODE"`
	Transport	string	`json:"TRANSPORT"`
}