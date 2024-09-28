package models

type AristaEosShowIpv6BgpSummary struct {
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Vrf	string	`json:"VRF"`
	Bgp_neigh	string	`json:"BGP_NEIGH"`
	Neigh_as	string	`json:"NEIGH_AS"`
	Msg_rcvd	string	`json:"MSG_RCVD"`
	Msg_sent	string	`json:"MSG_SENT"`
	In_queue	string	`json:"IN_QUEUE"`
	Out_queue	string	`json:"OUT_QUEUE"`
	Up_down	string	`json:"UP_DOWN"`
	State	string	`json:"STATE"`
	State_pfxrcd	string	`json:"STATE_PFXRCD"`
	State_pfxacc	string	`json:"STATE_PFXACC"`
}