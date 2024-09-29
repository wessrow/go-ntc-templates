package cisco_asa 

type ShowRunningConfigObjectNetwork struct {
	Name	string	`json:"NAME"`
	Description	string	`json:"DESCRIPTION"`
	Type	string	`json:"TYPE"`
	Host	string	`json:"HOST"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Start_ip	string	`json:"START_IP"`
	End_ip	string	`json:"END_IP"`
	Fqdn	string	`json:"FQDN"`
}

var ShowRunningConfigObjectNetwork_Template = `Value Required NAME (\S+)
Value DESCRIPTION (.+)
Value TYPE (host|subnet|range|fqdn)
Value HOST (\S+)
Value NETWORK (\S+)
Value NETMASK (\S+)
Value PREFIX_LENGTH (\d+)
Value START_IP (\S+)
Value END_IP (\S+)
Value FQDN (\S+)


Start
  ^object\s+network -> Continue.Record
  ^object\s+network\s+${NAME}\s*
  ^\s+description\s+${DESCRIPTION}\s*
  ^\s+${TYPE} -> Continue
  ^\s+subnet\s+${NETWORK}\s+${NETMASK}\s*
  ^\s+subnet\s+${NETWORK}\/${PREFIX_LENGTH}\s*
  ^\s+range\s+${START_IP}\s+${END_IP}\s*
  ^\s+host\s+${HOST}\s*
  ^\s+fqdn\s+${FQDN}\s*
  ^. -> Error
`