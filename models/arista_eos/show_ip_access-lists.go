package arista_eos

type ShowIpAccessLists struct {
	Destination   string `json:"DESTINATION"`
	Name          string `json:"NAME"`
	Source        string `json:"SOURCE"`
	Port_range    string `json:"PORT_RANGE"`
	Port_modifier string `json:"PORT_MODIFIER"`
	Modifier      string `json:"MODIFIER"`
	Sn            string `json:"SN"`
	Action        string `json:"ACTION"`
	Protocol      string `json:"PROTOCOL"`
}

var ShowIpAccessLists_Template string = `Value Filldown NAME (\S+)
Value Required SN (\d+)
Value ACTION (\w+)
Value PROTOCOL ([^any][\w+\d+]+)
Value SOURCE (host\s\d+\.\d+\.\d+\.\d+|any|\d+\.\d+\.\d+\.\d+/\d+)
Value PORT_MODIFIER (eq|gt|lt|neq|range)
Value PORT_RANGE ([(\S+\s\S+)]+)
Value DESTINATION (host\s\d+\.\d+\.\d+\.\d+|any|\d+\.\d+\.\d+\.\d+/\d+) 
Value MODIFIER (.*)

Start
  ^.*List\s+${NAME}(\s+\[readonly\])* 
  ^\s+${SN}\s+${ACTION}\s+${PROTOCOL}\s+${SOURCE}(\s+${PORT_MODIFIER})*(\s${PORT_RANGE})*\s+${DESTINATION}(\s+${MODIFIER})* -> Record

`
