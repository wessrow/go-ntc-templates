package models

type CiscoNxosShowLldpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Capabilities	string	`json:"CAPABILITIES"`
}