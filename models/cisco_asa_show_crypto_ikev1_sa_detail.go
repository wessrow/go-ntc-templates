package models

type CiscoAsaShowCryptoIkev1SaDetail struct {
	Sequence	string	`json:"SEQUENCE"`
	Peer	string	`json:"PEER"`
	Type	string	`json:"TYPE"`
	Dir	string	`json:"DIR"`
	Rekey	string	`json:"REKEY"`
	State	string	`json:"STATE"`
	Encryption	string	`json:"ENCRYPTION"`
	Hash	string	`json:"HASH"`
	Authentication	string	`json:"AUTHENTICATION"`
	Lifetime	string	`json:"LIFETIME"`
}