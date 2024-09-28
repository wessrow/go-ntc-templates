package models

type CiscoIosShowApCdpNeighbors struct {
	Ap_name	string	`json:"AP_NAME"`
	Ap_ip	string	`json:"AP_IP"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Neighbor_ip	string	`json:"NEIGHBOR_IP"`
	Neighbor_port	string	`json:"NEIGHBOR_PORT"`
}