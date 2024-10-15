package alcatel_sros

type Ping struct {
	Loss_pct        string `json:"LOSS_PCT"`
	Bounce_qty      string `json:"BOUNCE_QTY"`
	Rtt_max         string `json:"RTT_MAX"`
	Rtt_min         string `json:"RTT_MIN"`
	Pkt_size        string `json:"PKT_SIZE"`
	Rtt_avg         string `json:"RTT_AVG"`
	Std_dev         string `json:"STD_DEV"`
	Destination     string `json:"DESTINATION"`
	Sent_qty        string `json:"SENT_QTY"`
	Success_qty     string `json:"SUCCESS_QTY"`
	Response_stream string `json:"RESPONSE_STREAM"`
	Duplicate_qty   string `json:"DUPLICATE_QTY"`
}

var Ping_Template string = `Value Required DESTINATION (\S+)
Value Required SENT_QTY (\d+)
Value Required SUCCESS_QTY (\d+)
Value LOSS_PCT (\S+)
Value RESPONSE_STREAM ([\.\!/U/Q]+)
Value DUPLICATE_QTY (\d+)
Value BOUNCE_QTY (\d+)
Value RTT_AVG (\d+.\d+)
Value RTT_MAX (\d+.\d+)
Value RTT_MIN (\d+.\d+)
Value STD_DEV (\d+.\d+)
Value PKT_SIZE (\d+)


Start
  ^PING\s+${DESTINATION}\s+${PKT_SIZE}\s+data\s+bytes*$$
  ^(?:${RESPONSE_STREAM})
  ^\.*$$
  ^\s*$$
  ^-+
  ^${SENT_QTY}\s+packet(?:s)?\s+transmitted,(?:\s+${BOUNCE_QTY}\s+packet(?:s)?\s+bounced,)?\s+${SUCCESS_QTY}\s+packet(?:s)?\s+received,\s+(?:${DUPLICATE_QTY}\s+duplicate(?:s)?)?(?:${LOSS_PCT}%\s+packet\s+loss)?
  ^(?:round-trip\s+min\s+=\s+${RTT_MIN}ms,\s+avg\s+=\s+${RTT_AVG}ms,\s+max\s+=\s+${RTT_MAX}ms,\s+stddev\s+=\s+${STD_DEV}ms)?
  # Error out if raw data does not match any above rules.
  ^.* -> Error "Could not parse line:"
`
