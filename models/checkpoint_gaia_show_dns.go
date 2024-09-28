package models

type CheckpointGaiaShowDns struct {
	Domain	string	`json:"DOMAIN"`
	Dns_servers	[]string	`json:"DNS_SERVERS"`
}