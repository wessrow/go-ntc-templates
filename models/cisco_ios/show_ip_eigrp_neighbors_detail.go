package cisco_ios

type ShowIpEigrpNeighborsDetail struct {
	Hold                   string `json:"HOLD"`
	Uptime                 string `json:"UPTIME"`
	Seq_num                string `json:"SEQ_NUM"`
	Srtt                   string `json:"SRTT"`
	Q_cnt                  string `json:"Q_CNT"`
	Topology_ids_from_peer string `json:"TOPOLOGY_IDS_FROM_PEER"`
	Retries                string `json:"RETRIES"`
	As                     string `json:"AS"`
	H                      string `json:"H"`
	Interface              string `json:"INTERFACE"`
	Version                string `json:"VERSION"`
	Max_neighbors          string `json:"MAX_NEIGHBORS"`
	Current_neighbors      string `json:"CURRENT_NEIGHBORS"`
	Ip_address             string `json:"IP_ADDRESS"`
	Rto                    string `json:"RTO"`
	Retrans                string `json:"RETRANS"`
	Prefixes               string `json:"PREFIXES"`
}

var ShowIpEigrpNeighborsDetail_Template string = `Value Filldown AS (\d+)
Value Required H (\d+)
Value Required IP_ADDRESS ([0-9A-Fa-f:\.]+)
Value INTERFACE (\S+)
Value HOLD (\d+)
Value UPTIME (\S+)
Value SRTT (\d+)
Value RTO (\d+)
Value Q_CNT (\d+)
Value SEQ_NUM (\d+)
Value VERSION (\S+)
Value RETRANS (\d+)
Value RETRIES (\d+)
Value PREFIXES (\d+)
Value TOPOLOGY_IDS_FROM_PEER (\d+)
Value Fillup MAX_NEIGHBORS (\d+)
Value Fillup CURRENT_NEIGHBORS (\d+)

Start
  ^\s*(IP-|)EIGRP(-IPv[46](:\(\d+\))?|)\s+[Nn]eighbors\s+for\s+(process\s+|AS\()${AS}(\)|)\s*$$
  ^\s*(IP-|)EIGRP(-IPv[46](:\(\d+\))?|)\s+[Nn]eighbors\s+for\s+(process\s+|AS\()${AS}(\)|)\s+VRF default\s*$$
  ^\s*H\s+Address\s+Interface\s+Hold\s+Uptime\s+SRTT\s+RTO\s+Q\s+Seq\s*$$
  ^\s+\(sec\)\s+\(ms\)\s+Cnt\s+Num\s*$$ -> Neighbors

Neighbors
  ^\s*${H}\s+${IP_ADDRESS}\s+${INTERFACE}\s+${HOLD}\s+${UPTIME}\s+${SRTT}\s+${RTO}\s+${Q_CNT}\s+${SEQ_NUM}\s*$$
  ^\s*Version\s+${VERSION},\s+Retrans:\s+${RETRANS},\s+Retries:\s+${RETRIES},\s+Prefixes:\s+${PREFIXES}\s*$$
  ^\s*Topology-ids\s+from\s+peer\s+\-\s+${TOPOLOGY_IDS_FROM_PEER}\s*$$ -> Record
  ^\s*Max\s+Nbrs\:\s+${MAX_NEIGHBORS}\,\s+Current\s+Nbrs\:\s+${CURRENT_NEIGHBORS}\s*$$
  ^\s*$$
  ^.* -> Error

EOF
`
