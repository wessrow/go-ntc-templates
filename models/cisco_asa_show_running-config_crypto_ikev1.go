package models

type CiscoAsaShowRunningConfigCryptoIkev1 struct {
	Ike_version	string	`json:"IKE_VERSION"`
	Policy_id	string	`json:"POLICY_ID"`
	Auth_method	string	`json:"AUTH_METHOD"`
	Encryption	string	`json:"ENCRYPTION"`
	Auth_algorithm	string	`json:"AUTH_ALGORITHM"`
	Dh_group	string	`json:"DH_GROUP"`
	Lifetime	string	`json:"LIFETIME"`
}