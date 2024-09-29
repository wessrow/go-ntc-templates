package oneaccess_oneos 

type ShowIpInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Ok	string	`json:"OK"`
	Status	string	`json:"STATUS"`
	Protocol	string	`json:"PROTOCOL"`
	Description	string	`json:"DESCRIPTION"`
}

var ShowIpInterfaceBrief_Template = `Value INTERFACE (\S+(?:\s\S+)?)
Value IP_ADDRESS (\S+)
Value OK (YES|NO|Yes|No|yes|no)
Value STATUS (up|down|administratively down)
Value PROTOCOL (up|down)
Value DESCRIPTION (.*?)

Start
  ^Interface\s+IP[-\s]Address\s+OK\?\s+Status\s+Protocol\s+Description\s*$$
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${OK}\s+${STATUS}\s+${PROTOCOL}\s+${DESCRIPTION}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`