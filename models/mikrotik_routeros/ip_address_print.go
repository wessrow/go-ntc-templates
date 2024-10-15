package mikrotik_routeros

type IpAddressPrint struct {
	Num       string `json:"NUM"`
	Flags     string `json:"FLAGS"`
	Comment   string `json:"COMMENT"`
	Ip        string `json:"IP"`
	Subnet    string `json:"SUBNET"`
	Network   string `json:"NETWORK"`
	Interface string `json:"INTERFACE"`
}

var IpAddressPrint_Template string = `Value NUM (\d+)
Value FLAGS ((?:X|I|D)+)
Value COMMENT (.*)
Value IP (\S+)
Value SUBNET (\d+)
Value NETWORK (\S+)
Value INTERFACE (.*?)

Start
  ^\s*Flags:\s*(?:X\s*-\s*disabled,)?\s*(?:I\s*-\s*invalid,)?\s*(?:D\s*-\s*(?:DYNAMIC|dynamic))?\s*$$
  ^\s*Columns:\s*ADDRESS,\s*NETWORK,\s*INTERFACE\s*$$
  ^\s*#\s+ADDRESS\s+NETWORK\s+INTERFACE\s*$$ -> IPsTable
  ^\s*$$
  ^. -> Error

IPsTable
  ^\s*${NUM}\s+(?:${FLAGS}\s+)?${IP}/${SUBNET}\s+${NETWORK}\s+${INTERFACE}\s*$$ -> Record
  ^\s*;{3}\s*${COMMENT}\s*$$
  ^\s*${NUM}\s+(?:${FLAGS}\s+)?;;;\s+${COMMENT}\s*$$
  ^\s*${IP}/${SUBNET}\s+${NETWORK}\s+${INTERFACE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
