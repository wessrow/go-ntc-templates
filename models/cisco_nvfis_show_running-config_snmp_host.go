package models

type CiscoNvfisShowRunningConfigSnmpHost struct {
	Username	string	`json:"USERNAME"`
	Hostname	string	`json:"HOSTNAME"`
	Version	string	`json:"VERSION"`
	Port	string	`json:"PORT"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Securitylevel	string	`json:"SECURITYLEVEL"`
}