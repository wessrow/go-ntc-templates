package checkpoint_gaia

type ShowDns struct {
	Domain      string   `json:"DOMAIN"`
	Dns_servers []string `json:"DNS_SERVERS"`
}

var ShowDns_Template string = `Value DOMAIN (\S+)
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
