package models

type AristaEosShowMlag struct {
	Domain	string	`json:"DOMAIN"`
	Interface	string	`json:"INTERFACE"`
	Peer_addr	string	`json:"PEER_ADDR"`
	Peer_link	string	`json:"PEER_LINK"`
	State	string	`json:"STATE"`
}