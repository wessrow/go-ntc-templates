package models

type CiscoIosShowDhcpLease struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Interface	string	`json:"INTERFACE"`
	Server	string	`json:"SERVER"`
	State_id	string	`json:"STATE_ID"`
	State	string	`json:"STATE"`
	Transaction_id	string	`json:"TRANSACTION_ID"`
	Time_lease	string	`json:"TIME_LEASE"`
	Time_renewal	string	`json:"TIME_RENEWAL"`
	Time_rebound	string	`json:"TIME_REBOUND"`
	Time_next_fire	string	`json:"TIME_NEXT_FIRE"`
	Gateway	string	`json:"GATEWAY"`
	Retry_count	string	`json:"RETRY_COUNT"`
	Client_id	string	`json:"CLIENT_ID"`
	Hostname	string	`json:"HOSTNAME"`
}