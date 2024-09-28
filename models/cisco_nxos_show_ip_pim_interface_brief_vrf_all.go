package models

type CiscoNxosShowIpPimInterfaceBriefVrfAll struct {
	Vrf	string	`json:"VRF"`
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Pim_dr_address	string	`json:"PIM_DR_ADDRESS"`
	Neighbor_count	string	`json:"NEIGHBOR_COUNT"`
	Border_interface	string	`json:"BORDER_INTERFACE"`
}