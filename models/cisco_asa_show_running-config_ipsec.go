package models

type CiscoAsaShowRunningConfigIpsec struct {
	Ike_version	string	`json:"IKE_VERSION"`
	Name	string	`json:"NAME"`
	Encryption	[]string	`json:"ENCRYPTION"`
	Auth	[]string	`json:"AUTH"`
}