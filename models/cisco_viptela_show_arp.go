package models

type CiscoViptelaShowArp struct {
	Vpn	string	`json:"VPN"`
	Name	string	`json:"NAME"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	State	string	`json:"STATE"`
	Idle_timer	string	`json:"IDLE_TIMER"`
	Uptime	string	`json:"UPTIME"`
}