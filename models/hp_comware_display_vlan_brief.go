package models

type HpComwareDisplayVlanBrief struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
}

var HpComwareDisplayVlanBrief_Template = `Value VLAN_ID (\d+)
Value VLAN_NAME (\w+)

Start
  ^VLAN ID\s+ -> VLAN

VLAN
  ^${VLAN_ID}\s+${VLAN_NAME} -> Record

`