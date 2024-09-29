package models

type HpComwareDisplayMacAddress struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Vlan_id	string	`json:"VLAN_ID"`
	State	string	`json:"STATE"`
	Interface	string	`json:"INTERFACE"`
	Aging	string	`json:"AGING"`
}

var HpComwareDisplayMacAddress_Template = `Value MAC_ADDRESS (^[\d\w-]+)
Value VLAN_ID ([\d]+)
Value STATE (\S+)
Value INTERFACE (\S+)
Value AGING (\S+)

Start
  ^${MAC_ADDRESS}\s+${VLAN_ID}\s+${STATE}\s+${INTERFACE}\s+${AGING} -> Record

`