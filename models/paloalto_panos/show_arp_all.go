package paloalto_panos

type ShowArpAll struct {
	Interface   string `json:"INTERFACE"`
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Port        string `json:"PORT"`
	Status      string `json:"STATUS"`
	Ttl         string `json:"TTL"`
}

var ShowArpAll_Template string = `Value INTERFACE (\S+)
Value IP_ADDRESS (\S+)
Value MAC_ADDRESS ([0-9A-Fa-f:]+|incomplete)
Value PORT (\S+)
Value STATUS (\S+)
Value TTL (\d+)

Start
  ^interface\s+ip address\s+hw address\s+port\s+status\s+ttl -> Start_record

Start_record
  ^\s*---
  ^${INTERFACE}\s+${IP_ADDRESS}\s+\(*${MAC_ADDRESS}\)*\s+${PORT}\s+${STATUS}\s+${TTL} -> Record
  ^\s*$$
  ^. -> Error
`
