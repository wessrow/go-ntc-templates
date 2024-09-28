package models

type BroadcomIcosShowLldpRemoteDeviceAll struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Remote_id	string	`json:"REMOTE_ID"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
}