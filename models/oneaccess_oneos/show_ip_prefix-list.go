package oneaccess_oneos

type ShowIpPrefixList struct {
	Le          string `json:"LE"`
	Ge          string `json:"GE"`
	Sequence    string `json:"SEQUENCE"`
	Network     string `json:"NETWORK"`
	Description string `json:"DESCRIPTION"`
	Action      string `json:"ACTION"`
	Netmask     string `json:"NETMASK"`
	Protocol    string `json:"PROTOCOL"`
	Name        string `json:"NAME"`
}

var ShowIpPrefixList_Template string = `Value Required,Filldown PROTOCOL (\S+)
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
