package models

type CiscoNxosShowLldpNeighborsDetail struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Capabilities	string	`json:"CAPABILITIES"`
	Vlan_id	string	`json:"VLAN_ID"`
}