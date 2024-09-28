package models

type CiscoNxosShowNvePeers struct {
	Interface	string	`json:"INTERFACE"`
	Peer	string	`json:"PEER"`
	State	string	`json:"STATE"`
	Type	string	`json:"TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
}