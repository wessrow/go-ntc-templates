package models

type AristaEosShowIpOspfSummary struct {
	Instance	string	`json:"INSTANCE"`
	Router_id	string	`json:"ROUTER_ID"`
	Vrf	string	`json:"VRF"`
	Area	string	`json:"AREA"`
	Type	string	`json:"TYPE"`
	Interfaces	string	`json:"INTERFACES"`
	Neighbors	string	`json:"NEIGHBORS"`
	Neighbors_full	string	`json:"NEIGHBORS_FULL"`
	Router_lsas	string	`json:"ROUTER_LSAS"`
	Network_lsas	string	`json:"NETWORK_LSAS"`
	Summary_lsas	string	`json:"SUMMARY_LSAS"`
	Asbr_lsas	string	`json:"ASBR_LSAS"`
	Nssa_lsas	string	`json:"NSSA_LSAS"`
}