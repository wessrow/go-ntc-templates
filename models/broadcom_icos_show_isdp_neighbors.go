package models

type BroadcomIcosShowIsdpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Holdtime	string	`json:"HOLDTIME"`
	Capabilities	string	`json:"CAPABILITIES"`
	Platform	string	`json:"PLATFORM"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
}