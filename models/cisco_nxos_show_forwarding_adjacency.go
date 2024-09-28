package models

type CiscoNxosShowForwardingAdjacency struct {
	Slot	string	`json:"SLOT"`
	Nexthop	string	`json:"NEXTHOP"`
	Rewrite	string	`json:"REWRITE"`
	Interface	string	`json:"INTERFACE"`
	Origin_as	string	`json:"ORIGIN_AS"`
	Peer_as	string	`json:"PEER_AS"`
	Neighbor	string	`json:"NEIGHBOR"`
}