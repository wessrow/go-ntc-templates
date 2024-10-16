package fortinet

type ExecutePing struct {
	Received             string   `json:"RECEIVED"`
	Min_rtt              string   `json:"MIN_RTT"`
	Max_rtt              string   `json:"MAX_RTT"`
	Sent_packet_size     []string `json:"SENT_PACKET_SIZE"`
	Sent_packet_icmp_seq []string `json:"SENT_PACKET_ICMP_SEQ"`
	Sent_packet_ttl      []string `json:"SENT_PACKET_TTL"`
	Address              string   `json:"ADDRESS"`
	Transmitted          string   `json:"TRANSMITTED"`
	Sent_packet_address  []string `json:"SENT_PACKET_ADDRESS"`
	Sent_packet_time     []string `json:"SENT_PACKET_TIME"`
	Packet_loss          string   `json:"PACKET_LOSS"`
	Avg_rtt              string   `json:"AVG_RTT"`
}

var ExecutePing_Template string = `Value List SENT_PACKET_SIZE (\d+)
Value List SENT_PACKET_ADDRESS (\S+)
Value List SENT_PACKET_ICMP_SEQ (\d+)
Value List SENT_PACKET_TTL (\d+)
Value List SENT_PACKET_TIME (\d+(?:\.\d+)?)
Value ADDRESS (\S+)
Value TRANSMITTED (\d+)
Value RECEIVED (\d+)
Value PACKET_LOSS (\d+)
Value MIN_RTT (\d+(?:\.\d+)?)
Value AVG_RTT (\d+(?:\.\d+)?)
Value MAX_RTT (\d+(?:\.\d+)?)

Start
  ^\s*PING.*\s*$$
  ^\s*${SENT_PACKET_SIZE}\s+bytes\s+from\s+${SENT_PACKET_ADDRESS}:\s+icmp_seq=${SENT_PACKET_ICMP_SEQ}\s+ttl=${SENT_PACKET_TTL}\s+time=${SENT_PACKET_TIME}\s+.*$$
  ^\s*sendto\s+failed\s*$$
  ^\s*---\s+${ADDRESS}\s+ping\s+statistics\s+---\s*$$
  ^\s*${TRANSMITTED}\s+packets\s+transmitted\,\s+${RECEIVED}\s+packets\s+received\,\s+${PACKET_LOSS}\%\s+packet\s+loss\s*$$
  ^\s*round-trip\s+min/avg/max\s+=\s+${MIN_RTT}/${AVG_RTT}/${MAX_RTT}\s+ms\s*$$
  ^\s*$$
  ^. -> Error
`
