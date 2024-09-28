package models

type CiscoIosTraceroute struct {
	Hop_num	string	`json:"HOP_NUM"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Fqdn	string	`json:"FQDN"`
	Rtt_response	[]string	`json:"RTT_RESPONSE"`
	Details	string	`json:"DETAILS"`
}