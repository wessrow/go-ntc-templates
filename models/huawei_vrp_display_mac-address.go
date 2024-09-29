package models

type HuaweiVrpDisplayMacAddress struct {
	Destination_address	string	`json:"DESTINATION_ADDRESS"`
	Destination_port	string	`json:"DESTINATION_PORT"`
	Type	string	`json:"TYPE"`
	Vlan_id	string	`json:"VLAN_ID"`
}

var HuaweiVrpDisplayMacAddress_Template = `Value DESTINATION_ADDRESS ([0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4})
Value DESTINATION_PORT ([^,\s]+)
Value TYPE (\S+)
Value VLAN_ID (\d+)

Start
  ^-+$$
  ^\s+$$
  ^MAC\s+Address
  ^${DESTINATION_ADDRESS}\s+${VLAN_ID}/([^/]+)/([^/]+)\s+${DESTINATION_PORT}\s+${TYPE} -> Record
  ^Total\s+items
  ^. -> Error

`