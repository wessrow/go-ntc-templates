package models

type CiscoIosShowIpBgpVpnv4AllNeighbors struct {
	Neighbor	string	`json:"NEIGHBOR"`
	Vrf	string	`json:"VRF"`
	Remote_as	string	`json:"REMOTE_AS"`
	Local_as	string	`json:"LOCAL_AS"`
	Peer_group	string	`json:"PEER_GROUP"`
	Remote_router_id	string	`json:"REMOTE_ROUTER_ID"`
	Bgp_state	string	`json:"BGP_STATE"`
	Localhost_ip	string	`json:"LOCALHOST_IP"`
	Localhost_port	string	`json:"LOCALHOST_PORT"`
	Remote_ip	string	`json:"REMOTE_IP"`
	Remote_port	string	`json:"REMOTE_PORT"`
	Inbound_routemap	string	`json:"INBOUND_ROUTEMAP"`
	Outbound_routemap	string	`json:"OUTBOUND_ROUTEMAP"`
}