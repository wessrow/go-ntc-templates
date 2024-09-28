package models

type BrocadeFastironShowLldpNeighborsDetail struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_port_id	string	`json:"NEIGHBOR_PORT_ID"`
	Holdtime	string	`json:"HOLDTIME"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Neighbor_lacp_index	string	`json:"NEIGHBOR_LACP_INDEX"`
	Neighbor_max_frame_size	string	`json:"NEIGHBOR_MAX_FRAME_SIZE"`
	Neighbor_mau	string	`json:"NEIGHBOR_MAU"`
	Capabilities	string	`json:"CAPABILITIES"`
}