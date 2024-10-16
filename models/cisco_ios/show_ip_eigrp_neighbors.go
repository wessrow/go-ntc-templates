package cisco_ios

type ShowIpEigrpNeighbors struct {
	Ip_address string `json:"IP_ADDRESS"`
	Uptime     string `json:"UPTIME"`
	Q_cnt      string `json:"Q_CNT"`
	Seq_num    string `json:"SEQ_NUM"`
	As         string `json:"AS"`
	Interface  string `json:"INTERFACE"`
	Hold       string `json:"HOLD"`
	Srtt       string `json:"SRTT"`
	Rto        string `json:"RTO"`
}

var ShowIpEigrpNeighbors_Template string = `Value Filldown AS (\d+)
Value Required IP_ADDRESS ([0-9A-Fa-f:\.]+)
Value INTERFACE (\S+)
Value HOLD (\d+)
Value UPTIME (\S+)
Value SRTT (\d+)
Value RTO (\d+)
Value Q_CNT (\d+)
Value SEQ_NUM (\d+)

Start
  ^.*\s+${IP_ADDRESS}\s+${INTERFACE}\s+${HOLD}\s+${UPTIME}\s+${SRTT}\s+${RTO}\s+${Q_CNT}\s+${SEQ_NUM}\s*$$ -> Record
  ^\s*(IP-|)EIGRP(-IPv[46](:\(\d+\))?|)\s+[Nn]eighbors\s+for\s+(process\s+|AS\()${AS}(\)|)\s*$$
  ^\s*(IP-|)EIGRP(-IPv[46](:\(\d+\))?|)\s+[Nn]eighbors\s+for\s+(process\s+|AS\()${AS}(\)|)\s+VRF default\s*$$
  ^\s*$$
  ^\s*H\s+Address\s+Interface\s+Hold\s+Uptime\s+SRTT\s+RTO\s+Q\s+Seq\s*$$
  ^\s+\(sec\)\s+\(ms\)\s+Cnt\s+Num\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^\s*Load\s+for
  ^Time\s+source\s+is
  ^. -> Error

EOF
`
