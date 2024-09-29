package vyatta_vyos 

type ShowInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	[]string	`json:"IP_ADDRESS"`
	Status	string	`json:"STATUS"`
	Description	string	`json:"DESCRIPTION"`
}

var ShowInterfaces_Template = `Value INTERFACE (\S+)
Value List IP_ADDRESS (\S+)
Value STATUS (\S+)
Value DESCRIPTION (.+?)

Start
  ^Codes:
  ^Interface\s+IP\s+Address\s+S/L\s+Description
  ^-+\s+-+\s+-+\s+-+$$ -> Interface
  ^\s*$$
  ^. -> Error

Interface
  ^\S+\s+ -> Continue.Record
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${STATUS}\s+(${DESCRIPTION}|)\s*$$
  ^\s+${IP_ADDRESS}\s*$$
  ^\s*$$
  ^. -> Error
`