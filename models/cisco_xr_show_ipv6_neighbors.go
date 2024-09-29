package models

type CiscoXrShowIpv6Neighbors struct {
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	State	string	`json:"STATE"`
	Interface	string	`json:"INTERFACE"`
	Location	string	`json:"LOCATION"`
}

var CiscoXrShowIpv6Neighbors_Template = `Value IPV6_ADDRESS (.+?)
Value AGE (\S+)
Value MAC_ADDRESS (\S+)
Value STATE (\S+)
Value INTERFACE (\S+)
Value LOCATION (\S+)

Start
  ^\S+\s+\S+\s+\S+\s+\S+:\S+:\S+\.\S+\s+\S+\s*
  ^IPv6\s+Address\s+Age\s+Link-layer\s+Add(r?)\s+State\s+Interface\s+Location\s*$$ -> Begin
  ^\s*$$
  ^. -> Error

Begin
  ^\[?${IPV6_ADDRESS}\]?\s+${AGE}\s+${MAC_ADDRESS}\s+${STATE}\s+${INTERFACE}\s+${LOCATION}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`