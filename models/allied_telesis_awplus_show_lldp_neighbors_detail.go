package models

type AlliedTelesisAwplusShowLldpNeighborsDetail struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_port_description	string	`json:"NEIGHBOR_PORT_DESCRIPTION"`
	Neighbor_port_type	string	`json:"NEIGHBOR_PORT_TYPE"`
	Neighbor_port_id	string	`json:"NEIGHBOR_PORT_ID"`
	Neighbor	string	`json:"NEIGHBOR"`
	Management_ip	string	`json:"MANAGEMENT_IP"`
}