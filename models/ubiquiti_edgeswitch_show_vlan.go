package models

type UbiquitiEdgeswitchShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Type	string	`json:"TYPE"`
}

var UbiquitiEdgeswitchShowVlan_Template = `Value Required VLAN_ID (\d+)
Value VLAN_NAME (.*?)
Value TYPE (\S+)

Start
  ^${VLAN_ID}\s+${VLAN_NAME}\s+${TYPE}\s*$$ -> Record
  ^VLAN\s+ID\s+VLAN\s+Name\s+VLAN\s+Type\s*$$
  ^-+\s+-+\s+-+\s*$$
  ^\s*$$
  ^. -> Error

`