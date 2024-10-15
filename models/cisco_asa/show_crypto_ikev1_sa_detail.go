package cisco_asa

type ShowCryptoIkev1SaDetail struct {
	Peer           string `json:"PEER"`
	Dir            string `json:"DIR"`
	Rekey          string `json:"REKEY"`
	Encryption     string `json:"ENCRYPTION"`
	Authentication string `json:"AUTHENTICATION"`
	Lifetime       string `json:"LIFETIME"`
	Sequence       string `json:"SEQUENCE"`
	Type           string `json:"TYPE"`
	State          string `json:"STATE"`
	Hash           string `json:"HASH"`
}

var ShowCryptoIkev1SaDetail_Template string = `Value SEQUENCE (\d+)
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
