package models

type CiscoIosShowNvePeers struct {
	Interface	string	`json:"INTERFACE"`
	Vni	string	`json:"VNI"`
	Type	string	`json:"TYPE"`
	Peer	string	`json:"PEER"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	State	string	`json:"STATE"`
}