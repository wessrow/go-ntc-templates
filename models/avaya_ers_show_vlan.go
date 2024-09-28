package models

type AvayaErsShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Vlan_type	string	`json:"VLAN_TYPE"`
	Vlan_protocol	string	`json:"VLAN_PROTOCOL"`
	Vlan_pid	string	`json:"VLAN_PID"`
	Vlan_active	string	`json:"VLAN_ACTIVE"`
	Vlan_ivl_svl	string	`json:"VLAN_IVL_SVL"`
	Vlan_mgmt	string	`json:"VLAN_MGMT"`
	Vlan_port_members	[]string	`json:"VLAN_PORT_MEMBERS"`
}