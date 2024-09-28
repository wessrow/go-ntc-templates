package models

type CiscoApicFabricShowVlanExtended struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	[]string	`json:"VLAN_NAME"`
	Vlan_encap	[]string	`json:"VLAN_ENCAP"`
	Vlan_ports	[]string	`json:"VLAN_PORTS"`
	Node_number	string	`json:"NODE_NUMBER"`
	Node_name	string	`json:"NODE_NAME"`
}