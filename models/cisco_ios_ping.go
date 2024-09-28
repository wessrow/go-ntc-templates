package models

type CiscoIosPing struct {
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