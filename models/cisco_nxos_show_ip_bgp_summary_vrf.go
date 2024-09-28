package models

type CiscoNxosShowIpBgpSummaryVrf struct {
	Vrf	string	`json:"VRF"`
	Address_family	string	`json:"ADDRESS_FAMILY"`
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Bgp_neigh	string	`json:"BGP_NEIGH"`
	Bgp_ver	string	`json:"BGP_VER"`
	Neigh_as	string	`json:"NEIGH_AS"`
	Msg_rcvd	string	`json:"MSG_RCVD"`
	Msg_sent	string	`json:"MSG_SENT"`
	Tblver	string	`json:"TBLVER"`
	In_queue	string	`json:"IN_QUEUE"`
	Out_queue	string	`json:"OUT_QUEUE"`
	Up_down	string	`json:"UP_DOWN"`
	State_pfxrcd	string	`json:"STATE_PFXRCD"`
}