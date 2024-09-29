package models

type CheckpointGaiaShowDns struct {
	Domain	string	`json:"DOMAIN"`
	Dns_servers	[]string	`json:"DNS_SERVERS"`
}

var CheckpointGaiaShowDns_Template = `Value DOMAIN (\S+)
Value List DNS_SERVERS (\S+)

Start
  ^DNS\ssetup
  ^Name\s+Value
  ^Domain\s+${DOMAIN}
  ^DNS\sserver\s+${DNS_SERVERS} -> Continue
  ^DNS\sserver
  ^\s*$$
  ^. -> Error

`