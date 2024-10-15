package cisco_asa

type Ping struct {
	Destination     string `json:"DESTINATION"`
	Timeout         string `json:"TIMEOUT"`
	Response_stream string `json:"RESPONSE_STREAM"`
	Success_qty     string `json:"SUCCESS_QTY"`
	Rtt_max         string `json:"RTT_MAX"`
	Sent_type       string `json:"SENT_TYPE"`
	Success_pct     string `json:"SUCCESS_PCT"`
	Rtt_min         string `json:"RTT_MIN"`
	Rtt_avg         string `json:"RTT_AVG"`
	Sent_qty        string `json:"SENT_QTY"`
}

var Ping_Template string = `Value Required SENT_QTY (\d+)
Value Required SENT_TYPE (.*)
Value Required DESTINATION (\S+)
Value Required TIMEOUT (\d+)
Value Required RESPONSE_STREAM ([\.\!/Q/U]+)
Value Required SUCCESS_PCT (\d+)
Value Required SUCCESS_QTY (\d+)
Value RTT_MIN (\d+)
Value RTT_AVG (\d+)
Value RTT_MAX (\d+)


Start
  ^Type\s+escape\s+sequence\s+to\s+abort.
  ^Sending\s+${SENT_QTY},\s+${SENT_TYPE}\s+to\s+${DESTINATION},\s+timeout\s+is\s+${TIMEOUT}\s+seconds:
  ^${RESPONSE_STREAM}
  ^Success\s+rate\s+is\s+${SUCCESS_PCT}\s+percent\s+\(${SUCCESS_QTY}\/\d+\)(?:,\s+round-trip\s+min/avg/max\s+\=\s+)?(?:${RTT_MIN}/${RTT_AVG}/${RTT_MAX})?(?:\sms)?
  ^\s*$$
  # Error out if raw data does not match any above rules.
  ^.* -> Error "Could not parse line:"
`
