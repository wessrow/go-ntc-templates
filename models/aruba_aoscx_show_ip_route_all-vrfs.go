package models

type ArubaAoscxShowIpRouteAllVrfs struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Vrf	string	`json:"VRF"`
	Interface	[]string	`json:"INTERFACE"`
	Metric	[]string	`json:"METRIC"`
	Status	[]string	`json:"STATUS"`
}