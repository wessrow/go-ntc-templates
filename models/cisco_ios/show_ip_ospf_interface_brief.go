package cisco_ios

type ShowIpOspfInterfaceBrief struct {
	Process       string `json:"PROCESS"`
	Area          string `json:"AREA"`
	Ip_address    string `json:"IP_ADDRESS"`
	Prefix_length string `json:"PREFIX_LENGTH"`
	Cost          string `json:"COST"`
	State         string `json:"STATE"`
	Neighbors_fc  string `json:"NEIGHBORS_FC"`
	Interface     string `json:"INTERFACE"`
}

var ShowIpOspfInterfaceBrief_Template string = `Value INTERFACE (\S+)
Value PROCESS (\d+)
Value AREA (\d+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value COST (\d+)
Value STATE (\S+)
Value NEIGHBORS_FC (\d+/\d+)

Start
  ^${INTERFACE}\s+${PROCESS}\s+${AREA}\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${COST}\s+${STATE}\s+${NEIGHBORS_FC} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
