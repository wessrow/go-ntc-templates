package cisco_ios 

type ShowIpPrefixList struct {
	Name	string	`json:"NAME"`
	Seq	string	`json:"SEQ"`
	Action	string	`json:"ACTION"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Le	string	`json:"LE"`
	Ge	string	`json:"GE"`
}

var ShowIpPrefixList_Template = `Value Required,Filldown NAME (\S+)
Value Required SEQ (\d+)
Value ACTION (\S+)
Value NETWORK ([0-9a-f:\.]+)
Value NETMASK (\d+)
Value LE (\d+)
Value GE (\d+)

Start
  ^ip\s+prefix-list\s+${NAME}:
  ^\s+seq\s+${SEQ}\s+${ACTION}\s+${NETWORK}/${NETMASK}(?:\s+ge\s+${GE}|)(?:\s+le\s+${LE}|)\s*$$ -> Record
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^.+ -> Error
`