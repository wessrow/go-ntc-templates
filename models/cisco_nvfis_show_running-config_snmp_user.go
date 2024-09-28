package models

type CiscoNvfisShowRunningConfigSnmpUser struct {
	Username	string	`json:"USERNAME"`
	Version	string	`json:"VERSION"`
	User_group	string	`json:"USER_GROUP"`
	Auth_protocol	string	`json:"AUTH_PROTOCOL"`
	Priv_protocol	string	`json:"PRIV_PROTOCOL"`
	Auth_key	string	`json:"AUTH_KEY"`
	Priv_key	string	`json:"PRIV_KEY"`
}