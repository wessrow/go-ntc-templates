package models

type ArubaOsShowIpInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Admin	string	`json:"ADMIN"`
	Protocol	string	`json:"PROTOCOL"`
}

var ArubaOsShowIpInterfaceBrief_Template = `Value INTERFACE (\S+\s\S+)
Value IP_ADDRESS (\S+)
Value NETMASK (\S+)
Value ADMIN (\S+)
Value PROTOCOL (\S+)

Start
  ^${INTERFACE}\s+${IP_ADDRESS}\s+/\s+${NETMASK}\s+${ADMIN}\s+${PROTOCOL} -> Record

`