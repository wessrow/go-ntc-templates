package paloalto_panos 

type ShowInterfaceLogical struct {
	Interface	string	`json:"INTERFACE"`
	Id	string	`json:"ID"`
	Vsys	string	`json:"VSYS"`
	Zone	string	`json:"ZONE"`
	Forwarding	string	`json:"FORWARDING"`
	Vlan_id	string	`json:"VLAN_ID"`
	Ip_address	string	`json:"IP_ADDRESS"`
}

var ShowInterfaceLogical_Template = `Value INTERFACE (\S+)
Value ID (\d+)
Value VSYS (\d+)
Value ZONE (\S+)
Value FORWARDING (\S+)
Value VLAN_ID (\d+)
Value IP_ADDRESS (\S+)

Start
  ^${INTERFACE}\s+${ID}\s+${VSYS}\s+${ZONE}\s+${FORWARDING}+\s+${VLAN_ID}\s+${IP_ADDRESS} -> Record
  ^${INTERFACE}\s+${ID}\s+${VSYS}\s+${FORWARDING}\s+${VLAN_ID}\s+${IP_ADDRESS} -> Record
`