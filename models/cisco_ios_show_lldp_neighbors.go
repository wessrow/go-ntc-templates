package models

type CiscoIosShowLldpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Capabilities	string	`json:"CAPABILITIES"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
}