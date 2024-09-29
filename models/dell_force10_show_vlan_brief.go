package models

type DellForce10ShowVlanBrief struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Stg	string	`json:"STG"`
	Mac_aging	string	`json:"MAC_AGING"`
	Ip_address	string	`json:"IP_ADDRESS"`
}

var DellForce10ShowVlanBrief_Template = `Value VLAN_ID (\d+)
Value VLAN_NAME (\S+(\s\S+)*)
Value STG (\d+)
Value MAC_AGING (\d+)
Value IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3}(/\d{1,2})?|\w+)

Start
  ^\s*VLAN\s+Name -> VLAN

VLAN
  ^\s*${VLAN_ID}\s+${VLAN_NAME}?\s+${STG}\s+${MAC_AGING}\s+${IP_ADDRESS} -> Record

`