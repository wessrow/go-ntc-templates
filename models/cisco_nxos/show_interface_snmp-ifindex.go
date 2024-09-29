package cisco_nxos 

type ShowInterfaceSnmpIfindex struct {
	Port	string	`json:"PORT"`
	Ifmib	string	`json:"IFMIB"`
	Ifindex_hex	string	`json:"IFINDEX_HEX"`
}

var ShowInterfaceSnmpIfindex_Template = `Value PORT (\S+)
Value IFMIB (\S+)
Value IFINDEX_HEX (\S+)

Start
  ^Port\s+IFMIB\s+Ifindex\s+\(hex\)\s*$$
  ^${PORT}\s+${IFMIB}\s+\(${IFINDEX_HEX}\s*\)\s*$$ -> Record
  ^-*$$
  ^\s*$$
  ^. -> Error

EOF`