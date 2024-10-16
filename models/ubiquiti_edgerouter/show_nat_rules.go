package ubiquiti_edgerouter

type ShowNatRules struct {
	Protocol     string `json:"PROTOCOL"`
	Src_port     string `json:"SRC_PORT"`
	Dest_port    string `json:"DEST_PORT"`
	Rule         string `json:"RULE"`
	Type         string `json:"TYPE"`
	Interface    string `json:"INTERFACE"`
	Src_address  string `json:"SRC_ADDRESS"`
	Dest_address string `json:"DEST_ADDRESS"`
}

var ShowNatRules_Template string = `Value RULE (\d+)
Value TYPE (\S+)
Value INTERFACE (\S+)
Value SRC_ADDRESS ((\d+\.\d+\.\d+\.\d+((/\d+)?))|---)
Value DEST_ADDRESS ((\d+\.\d+\.\d+\.\d+((/\d+)?))|---)
Value PROTOCOL (tcp|udp|icmp|all|tcp_udp|---)
Value SRC_PORT (\S+)
Value DEST_PORT (\S+)

Start
  ^Type Codes:  SRC - source, DST - destination, MASQ - masquerade
  ^\s+X at the front of rule implies rule is excluded\s*$$
  ^(rule\s+type\s+intf\s+translation)\s*$$
  ^([-\s]+)$$
  ^${RULE}\s+${TYPE}\s+${INTERFACE}\s+daddr\s${SRC_ADDRESS}\sto\s${DEST_ADDRESS}\s*$$
  ^${RULE}\s+${TYPE}\s+${INTERFACE}\s+saddr\s${DEST_ADDRESS}\sto\s${SRC_ADDRESS}\s*$$
  ^\s+proto\-${PROTOCOL}\s+(dport|sport)\s(${SRC_PORT}\sto\s${DEST_PORT}|ANY)\s*$$ -> Record
  ^. -> Error
`
