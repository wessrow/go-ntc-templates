package mikrotik_routeros 

type Ping struct {
	Seq	[]string	`json:"SEQ"`
	Host	[]string	`json:"HOST"`
	Size	[]string	`json:"SIZE"`
	Ttl	[]string	`json:"TTL"`
	Time	[]string	`json:"TIME"`
	Status	[]string	`json:"STATUS"`
	Sent	string	`json:"SENT"`
	Received	string	`json:"RECEIVED"`
	Packet_loss	string	`json:"PACKET_LOSS"`
	Min_rtt	string	`json:"MIN_RTT"`
	Avg_rtt	string	`json:"AVG_RTT"`
	Max_rtt	string	`json:"MAX_RTT"`
}

var Ping_Template = `Value List SEQ (\d+)
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