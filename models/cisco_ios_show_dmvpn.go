package models

type CiscoIosShowDmvpn struct {
	Peer_nbma	string	`json:"PEER_NBMA"`
	Peer_tunnel	string	`json:"PEER_TUNNEL"`
	State	string	`json:"STATE"`
	Uptime	string	`json:"UPTIME"`
	Attribute	string	`json:"ATTRIBUTE"`
}