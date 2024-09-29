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

var CiscoAsaShowCryptoIkev1SaDetail_Template = `Value SEQUENCE (\d+)
Value PEER (\d+\.\d+\.\d+\.\d+)
Value TYPE (\w+)
Value DIR (\w+)
Value REKEY (Yes|No)
Value STATE (\S+)
Value ENCRYPTION (\w+)
Value HASH (\w+)
Value AUTHENTICATION (\w+)
Value LIFETIME (\d+)

Start
  ^IKE\s+Peer\s+Type\s+Dir\s+Rky\s+State\s+Encrypt\s+Hash\s+Auth\s+Lifetime\s*
  ^${SEQUENCE}\s+${PEER}\s+${TYPE}\s+${DIR}\s+${REKEY}\s+${STATE}\s+${ENCRYPTION}\s+${HASH}\s+${AUTHENTICATION}\s+${LIFETIME}\s* -> Record

`