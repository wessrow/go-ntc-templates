package mikrotik_routeros 

type IpAddressExportVerbose struct {
	Ip	string	`json:"IP"`
	Mask	string	`json:"MASK"`
	Comment	string	`json:"COMMENT"`
	Disabled	string	`json:"DISABLED"`
	Interface	string	`json:"INTERFACE"`
	Network	string	`json:"NETWORK"`
}

var IpAddressExportVerbose_Template = `Value IP (\S+)
Value MASK (\d+)
Value COMMENT (.*?)
Value DISABLED (yes|no)
Value INTERFACE (\S+)
Value NETWORK (\S+)

Start
  ^\s*#.*$$
  ^/ip\s+address\s*$$
  ^\s*add\s+address=${IP}/${MASK}(?:\s+comment="${COMMENT}")?\s+disabled=${DISABLED}\s+interface=${INTERFACE}\s+network=${NETWORK}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`