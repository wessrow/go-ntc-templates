package models

type CiscoAsaShowBgpSummary struct {
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Bgp_neigh	string	`json:"BGP_NEIGH"`
	Neigh_as	string	`json:"NEIGH_AS"`
	Up_down	string	`json:"UP_DOWN"`
	State_pfxrcd	string	`json:"STATE_PFXRCD"`
}