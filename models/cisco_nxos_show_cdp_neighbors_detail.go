package models

type CiscoNxosShowCdpNeighborsDetail struct {
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Platform	string	`json:"PLATFORM"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Interface_ip	string	`json:"INTERFACE_IP"`
	Capabilities	string	`json:"CAPABILITIES"`
}