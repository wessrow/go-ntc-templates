package models

type CiscoS300ShowLldpNeighbors struct {
	Neighbor_name	[]string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_interface	[]string	`json:"NEIGHBOR_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
}