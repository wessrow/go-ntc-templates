package models

type CiscoIosShowIpv6AccessLists struct {
	Acl_name	string	`json:"ACL_NAME"`
	Action	string	`json:"ACTION"`
	Protocol	string	`json:"PROTOCOL"`
	Ipv6_source	string	`json:"IPV6_SOURCE"`
	Ipv6_destination	string	`json:"IPV6_DESTINATION"`
	Src_port_match	string	`json:"SRC_PORT_MATCH"`
	Sequence	string	`json:"SEQUENCE"`
	Port_number	string	`json:"PORT_NUMBER"`
	Log	string	`json:"LOG"`
	Authentification	string	`json:"AUTHENTIFICATION"`
	Routing	string	`json:"ROUTING"`
	Src_port_range_start	string	`json:"SRC_PORT_RANGE_START"`
	Src_port_range_end	string	`json:"SRC_PORT_RANGE_END"`
}