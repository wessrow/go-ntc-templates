package cisco_nxos 

type ShowIpEigrpNeighbors struct {
	As	string	`json:"AS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Hold	string	`json:"HOLD"`
	Uptime	string	`json:"UPTIME"`
	Srtt	string	`json:"SRTT"`
	Rto	string	`json:"RTO"`
	Q_cnt	string	`json:"Q_CNT"`
	Seq_num	string	`json:"SEQ_NUM"`
}

var ShowIpEigrpNeighbors_Template = `Value Filldown AS (\d+)
Value IP_ADDRESS ([0-9A-Fa-f:\.]+)
Value INTERFACE (\S+)
Value Filldown VRF (\S+)
Value HOLD (\d+)
Value UPTIME (\d+:\d+:\d+)
Value SRTT (\d+)
Value RTO (\d+)
Value Q_CNT (\d+)
Value SEQ_NUM (\d+)

Start
  # Match the EIGRP process and AS number line
  ^IP-EIGRP neighbors for process ${AS} VRF ${VRF}\s*$$
  # Match the header lines
  ^\s*H\s+Address\s+Interface\s+Hold\s+Uptime\s+SRTT\s+RTO\s+Q\s+Seq\s*$$
  ^\s*\(sec\)\s+\(ms\)\s+Cnt\s+Num\s*$$
  # Match the data lines
  ^\d+\s+${IP_ADDRESS}\s+${INTERFACE}\s+${HOLD}\s+${UPTIME}\s+${SRTT}\s+${RTO}\s+${Q_CNT}\s+${SEQ_NUM}\s*$$ -> Record
  ^. -> Error

EOF
`