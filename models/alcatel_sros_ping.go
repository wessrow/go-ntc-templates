package models

type AlcatelSrosPing struct {
	Destination	string	`json:"DESTINATION"`
	Sent_qty	string	`json:"SENT_QTY"`
	Success_qty	string	`json:"SUCCESS_QTY"`
	Loss_pct	string	`json:"LOSS_PCT"`
	Response_stream	string	`json:"RESPONSE_STREAM"`
	Duplicate_qty	string	`json:"DUPLICATE_QTY"`
	Bounce_qty	string	`json:"BOUNCE_QTY"`
	Rtt_avg	string	`json:"RTT_AVG"`
	Rtt_max	string	`json:"RTT_MAX"`
	Rtt_min	string	`json:"RTT_MIN"`
	Std_dev	string	`json:"STD_DEV"`
	Pkt_size	string	`json:"PKT_SIZE"`
}