package models

type CiscoIosShowObjectGroup struct {
	Type	string	`json:"TYPE"`
	Name	string	`json:"NAME"`
	Description	string	`json:"DESCRIPTION"`
	Nested_groups	string	`json:"NESTED_GROUPS"`
	Any	string	`json:"ANY"`
	Host	string	`json:"HOST"`
	Host_range_start	string	`json:"HOST_RANGE_START"`
	Host_range_end	string	`json:"HOST_RANGE_END"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Protocol	string	`json:"PROTOCOL"`
	Port_match	string	`json:"PORT_MATCH"`
	Port	string	`json:"PORT"`
	Port_range_start	string	`json:"PORT_RANGE_START"`
	Port_range_end	string	`json:"PORT_RANGE_END"`
	Icmp_type	string	`json:"ICMP_TYPE"`
}