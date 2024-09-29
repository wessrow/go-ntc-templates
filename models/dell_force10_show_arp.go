package models

type DellForce10ShowArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Port	string	`json:"PORT"`
	Age	string	`json:"AGE"`
	Vlan_id	string	`json:"VLAN_ID"`
}

var DellForce10ShowArp_Template = `Value IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3})
Value MAC_ADDRESS ((?:[0-9a-fA-F]{2}\:){5}[0-9a-fA-F]{2})
Value PORT (\w{2} \d{1,2}\/\d{1,2}|\-)
Value AGE (\d+|\-)
Value VLAN_ID (\d+|\-)

Start
  ^Protocol\s+Address -> ARP

ARP
  ^Internet\s+${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${PORT}\s+(Vl )?${VLAN_ID} -> Record

`