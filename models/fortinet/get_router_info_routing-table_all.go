package fortinet

type GetRouterInfoRoutingTableAll struct {
	Type             string `json:"TYPE"`
	Destination      string `json:"DESTINATION"`
	Distance         string `json:"DISTANCE"`
	Metric           string `json:"METRIC"`
	Gateway          string `json:"GATEWAY"`
	Interface        string `json:"INTERFACE"`
	Last_time_update string `json:"LAST_TIME_UPDATE"`
}

var GetRouterInfoRoutingTableAll_Template string = `Value Filldown TYPE ((?:K|C|S|R|B|O|IA|N1|N2|E1|E2|i|L1|L2|ia|\*|\s)+?)
Value Filldown DESTINATION (\S+)
Value DISTANCE (\d+)
Value METRIC (\d+)
Value Required GATEWAY (\S+|is\s+directly\s+connected)
Value Required INTERFACE (\S+)
Value LAST_TIME_UPDATE (\S+)

Start
  ^\s*Codes:\s+K\s+-\s+kernel,\s+C\s+-\s+connected,\s+S\s+-\s+static,\s+R\s+-\s+RIP,\s+B\s+-\s+BGP\s*$$
  ^\s*O\s+-\s+OSPF,\s+IA\s+-\s+OSPF\s+inter\s+area\s*$$
  ^\s*N1\s+-\s+OSPF\s+NSSA\s+external\s+type\s+1,\s+N2\s+-\s+OSPF\s+NSSA\s+external\s+type\s+2\s*$$
  ^\s*E1\s+-\s+OSPF\s+external\s+type\s+1,\s+E2\s+-\s+OSPF\s+external\s+type\s+2\s*$$
  ^\s*i\s+-\s+IS-IS,\s+L1\s+-\s+IS-IS\s+level-1,\s+L2\s+-\s+IS-IS\s+level-2,\s+ia\s+-\s+IS-IS\s+inter\s+area\s*$$
  ^\s*\*\s+-\s+candidate\s+default\s*$$
  ^\s*Routing\s+table\s+for\s+VRF=\d+\s*$$
  ^\s*${TYPE}\s+${DESTINATION}(?:\s+\[${DISTANCE}/${METRIC}\]\s+via)?\s+${GATEWAY},\s*${INTERFACE}(?:\s*,\s*${LAST_TIME_UPDATE})?\s*$$ -> Record
  ^\s*(?:\[${DISTANCE}/${METRIC}\]\s+via\s+)?${GATEWAY},\s*${INTERFACE}(?:\s*,\s*${LAST_TIME_UPDATE})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
