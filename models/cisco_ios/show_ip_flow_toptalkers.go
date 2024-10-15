package cisco_ios

type ShowIpFlowToptalkers struct {
	Pkt        string `json:"PKT"`
	Src_intf   string `json:"SRC_INTF"`
	Src_ipaddr string `json:"SRC_IPADDR"`
	Src_port   string `json:"SRC_PORT"`
	Dst_intf   string `json:"DST_INTF"`
	Dst_ipaddr string `json:"DST_IPADDR"`
	Proto      string `json:"PROTO"`
	Dst_port   string `json:"DST_PORT"`
}

var ShowIpFlowToptalkers_Template string = `Value SRC_INTF (\S+)
Value SRC_IPADDR ([0-9A-Fa-f:\.]+)
Value SRC_PORT ([A-Z0-9]+)
Value DST_INTF (\S+)
Value DST_IPADDR ([0-9A-Fa-f:\.]+)
Value PROTO ([A-Z0-9]+)
Value DST_PORT ([A-Z0-9]+)
Value PKT ([0-9]+)

Start
  ^${SRC_INTF}\s+${SRC_IPADDR}\s+${DST_INTF}\s+${DST_IPADDR}\s+${PROTO}\s+${SRC_PORT}\s+${DST_PORT}\s+${PKT} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
