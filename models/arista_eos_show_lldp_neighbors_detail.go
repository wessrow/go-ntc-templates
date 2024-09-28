package models

type AristaEosShowLldpNeighborsDetail struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_count	string	`json:"NEIGHBOR_COUNT"`
	Age	string	`json:"AGE"`
}