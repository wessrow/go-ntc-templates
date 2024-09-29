package models

type AristaEosShowIpInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Protocol	string	`json:"PROTOCOL"`
	Mtu	string	`json:"MTU"`
}

var AristaEosShowIpInterfaceBrief_Template = `Value INTERFACE (\S+)
Value IP_ADDRESS (\S+)
Value STATUS (\S+)
Value PROTOCOL (\S+)
Value MTU (\d+)

Start
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${STATUS}\s+${PROTOCOL}\s+${MTU} -> Record

`