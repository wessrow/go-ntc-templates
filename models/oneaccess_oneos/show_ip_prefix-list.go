package oneaccess_oneos 

type ShowIpPrefixList struct {
	Protocol	string	`json:"PROTOCOL"`
	Name	string	`json:"NAME"`
	Description	string	`json:"DESCRIPTION"`
	Sequence	string	`json:"SEQUENCE"`
	Action	string	`json:"ACTION"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Le	string	`json:"LE"`
	Ge	string	`json:"GE"`
}

var ShowIpPrefixList_Template = `Value Required,Filldown PROTOCOL (\S+)
Value Required,Filldown NAME (\S+)
Value Filldown DESCRIPTION (.*)
Value Required SEQUENCE (\d+)
Value ACTION (\S+)
Value NETWORK ([0-9a-f:\.]+|any)
Value NETMASK (\d+)
Value LE (\d+)
Value GE (\d+)

Start
  ^${PROTOCOL}:\sip\s+prefix-list\s+${NAME}
  ^\s+Description:\s${DESCRIPTION}
  ^\s+seq\s+${SEQUENCE}\s+${ACTION}\s+${NETWORK}(/${NETMASK}|)(?:\s+ge\s+${GE}|)(?:\s+le\s+${LE}|)\s*$$ -> Record
  ^\s*$$
  ^.+ -> Error
`