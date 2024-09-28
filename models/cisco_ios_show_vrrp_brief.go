package models

type CiscoIosShowVrrpBrief struct {
	Interface	string	`json:"INTERFACE"`
	Group	string	`json:"GROUP"`
	Addr_family	string	`json:"ADDR_FAMILY"`
	Priority	string	`json:"PRIORITY"`
	Time	string	`json:"TIME"`
	Own	string	`json:"OWN"`
	Preempt	string	`json:"PREEMPT"`
	State	string	`json:"STATE"`
	Master_ip_address	string	`json:"MASTER_IP_ADDRESS"`
	Virtual_ip_address	string	`json:"VIRTUAL_IP_ADDRESS"`
}