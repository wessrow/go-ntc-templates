package cisco_asa

type ShowRunningConfigCryptoIkev1 struct {
	Lifetime       string `json:"LIFETIME"`
	Ike_version    string `json:"IKE_VERSION"`
	Policy_id      string `json:"POLICY_ID"`
	Auth_method    string `json:"AUTH_METHOD"`
	Encryption     string `json:"ENCRYPTION"`
	Auth_algorithm string `json:"AUTH_ALGORITHM"`
	Dh_group       string `json:"DH_GROUP"`
}

var ShowRunningConfigCryptoIkev1_Template string = `Value IKE_VERSION (\d+)
Value POLICY_ID (\d+)
Value AUTH_METHOD (\S+)
Value ENCRYPTION (\S+)
Value AUTH_ALGORITHM (\S+)
Value DH_GROUP (group\s*\d+)
Value LIFETIME (\d+)


Start
  ^crypto\s+ikev1 -> Continue.Record
  ^crypto\s+ikev${IKE_VERSION}\s+policy\s+${POLICY_ID}\s*
  ^\s+authentication\s+${AUTH_METHOD}\s*
  ^\s+encryption\s+${ENCRYPTION}\s*
  ^\s+hash\s+${AUTH_ALGORITHM}\s*
  ^\s+${DH_GROUP}\s*
  ^\s+lifetime\s+${LIFETIME}\s*
`
