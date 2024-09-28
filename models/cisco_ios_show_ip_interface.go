package models

type CiscoIosShowIpInterface struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Ip_address	[]string	`json:"IP_ADDRESS"`
	Prefix_length	[]string	`json:"PREFIX_LENGTH"`
	Vrf	string	`json:"VRF"`
	Mtu	string	`json:"MTU"`
	Ip_helper	[]string	`json:"IP_HELPER"`
	Outgoing_acl	string	`json:"OUTGOING_ACL"`
	Inbound_acl	string	`json:"INBOUND_ACL"`
}