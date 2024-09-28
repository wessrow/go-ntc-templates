package models

type CiscoIosShowMplsInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Mpls_ip	string	`json:"MPLS_IP"`
	Protocol	string	`json:"PROTOCOL"`
	Tunnel	string	`json:"TUNNEL"`
	Bgp	string	`json:"BGP"`
	Static	string	`json:"STATIC"`
	Operational	string	`json:"OPERATIONAL"`
}