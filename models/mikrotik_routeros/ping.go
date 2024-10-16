package mikrotik_routeros

type Ping struct {
	Packet_loss string   `json:"PACKET_LOSS"`
	Min_rtt     string   `json:"MIN_RTT"`
	Avg_rtt     string   `json:"AVG_RTT"`
	Seq         []string `json:"SEQ"`
	Host        []string `json:"HOST"`
	Size        []string `json:"SIZE"`
	Received    string   `json:"RECEIVED"`
	Max_rtt     string   `json:"MAX_RTT"`
	Ttl         []string `json:"TTL"`
	Time        []string `json:"TIME"`
	Status      []string `json:"STATUS"`
	Sent        string   `json:"SENT"`
}

var Ping_Template string = `Value List SEQ (\d+)
Value List HOST (\S+)
Value List SIZE (\d+)
Value List TTL (\d+)
Value List TIME (\d+)
Value List STATUS (.*)
Value SENT (\d+)
Value RECEIVED (\d+)
Value PACKET_LOSS (-?\d+)
Value MIN_RTT (\d+)
Value AVG_RTT (\d+)
Value MAX_RTT (\d+)

Start
  ^\s*SEQ\s+HOST\s+SIZE\s+TTL\s+TIME\s+STATUS\s*$$ -> PingsTable
  ^\s*$$
  ^. -> Error

PingsTable
  ^\s*${SEQ}\s+${HOST}(?:\s+${SIZE})?(?:\s+${TTL})?(?:\s+${TIME}ms)?(?:\s+${STATUS})?\s*$$
  ^\s*sent=${SENT}\s+received=${RECEIVED}\s+packet-loss=${PACKET_LOSS}%(?:\s+min-rtt=${MIN_RTT}ms\s+avg-rtt=${AVG_RTT}ms\s+max-rtt=${MAX_RTT}ms)?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
