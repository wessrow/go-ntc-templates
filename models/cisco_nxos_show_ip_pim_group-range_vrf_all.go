package models

type CiscoNxosShowIpPimGroupRangeVrfAll struct {
	Vrf	string	`json:"VRF"`
	Group_range	string	`json:"GROUP_RANGE"`
	Action	string	`json:"ACTION"`
	Mode	string	`json:"MODE"`
	Rp_address	string	`json:"RP_ADDRESS"`
	Shared_tree_range	string	`json:"SHARED_TREE_RANGE"`
	Origin	string	`json:"ORIGIN"`
}