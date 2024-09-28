package models

type ArubaAoscxShowLldpNeighborsInfoDetail struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Capabilities_supported	string	`json:"CAPABILITIES_SUPPORTED"`
	Capabilities	string	`json:"CAPABILITIES"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Neighbor_port_id	string	`json:"NEIGHBOR_PORT_ID"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
}