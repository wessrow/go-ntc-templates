package cisco_ios 

type Ping struct {
	Sent_qty	string	`json:"SENT_QTY"`
	Sent_type	string	`json:"SENT_TYPE"`
	Destination	string	`json:"DESTINATION"`
	Timeout	string	`json:"TIMEOUT"`
	Source	string	`json:"SOURCE"`
	Response_stream	[]string	`json:"RESPONSE_STREAM"`
	Success_pct	string	`json:"SUCCESS_PCT"`
	Success_qty	string	`json:"SUCCESS_QTY"`
	Rtt_min	string	`json:"RTT_MIN"`
	Rtt_avg	string	`json:"RTT_AVG"`
	Rtt_max	string	`json:"RTT_MAX"`
	Responders	[]string	`json:"RESPONDERS"`
}

var Ping_Template = `Value Required SENT_QTY (\d+)
Value Required SENT_TYPE (.*)
Value Required DESTINATION (\S+)
Value Required TIMEOUT (\d+)
Value SOURCE (\S+)
Value List RESPONSE_STREAM ([\.\!/U/Q]+)
Value SUCCESS_PCT (\d+)
Value SUCCESS_QTY (\d+)
Value RTT_MIN (\d+)
Value RTT_AVG (\d+)
Value RTT_MAX (\d+)
Value List RESPONDERS (.*)

Start
  ^Type\s+escape\s+sequence\s+to\s+abort.
  ^Sending\s+${SENT_QTY},\s+${SENT_TYPE}\s+to\s+${DESTINATION},\s+timeout\s+is\s+${TIMEOUT}\s+seconds:
  ^Reply\s+to\s+request\s+\d+\s+from\s+${RESPONDERS},\s+\d+\s+ms
  ^Packet\s+sent\s+with\s+a\s+source\s+address\s+of\s+${SOURCE}
  ^Packet\s+sent\s+with\s+the\s+DF\s+bit\s+set
  ^${RESPONSE_STREAM}
  ^Success\s+rate\s+is\s+${SUCCESS_PCT}\s+percent\s+\(${SUCCESS_QTY}\/\d+\)(?:,\s+round-trip\s+min/avg/max\s+\=\s+)?(?:${RTT_MIN}/${RTT_AVG}/${RTT_MAX})?(?:\sms)?
  ^\s*$$
  # Error out if raw data does not match any above rules.
  ^.* -> Error "Could not parse line:"`