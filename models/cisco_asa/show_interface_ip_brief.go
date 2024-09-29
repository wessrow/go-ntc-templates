package cisco_asa 

type ShowInterfaceIpBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Proto	string	`json:"PROTO"`
}

var ShowInterfaceIpBrief_Template = `Value INTERFACE (\S+)
Value IP_ADDRESS (\S+)
Value STATUS (up|down|administratively down)
Value PROTO (up|down)

Start
  ^Interface\s+IP-Address\s+OK\?\s+Method\s+Status\s+Protocol\s*$$
  ^${INTERFACE}\s+${IP_ADDRESS}\s+\w+\s+\w+\s+${STATUS}\s+${PROTO} -> Record
  ^\s*$$
  ^. -> Error
`