package models

type CiscoNxosShowIpPimNeighborVrfAll struct {
	Vrf	string	`json:"VRF"`
	Neighbor	string	`json:"NEIGHBOR"`
	Interface	string	`json:"INTERFACE"`
	Uptime	string	`json:"UPTIME"`
	Expires	string	`json:"EXPIRES"`
	Dr_priority	string	`json:"DR_PRIORITY"`
	Bidir_capable	string	`json:"BIDIR_CAPABLE"`
	Bfd_state	string	`json:"BFD_STATE"`
	Ecmp_redirect_capable	string	`json:"ECMP_REDIRECT_CAPABLE"`
}