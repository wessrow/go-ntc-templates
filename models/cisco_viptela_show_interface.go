package models

type CiscoViptelaShowInterface struct {
	Vpn	string	`json:"VPN"`
	Interface	string	`json:"INTERFACE"`
	Af_type	string	`json:"AF_TYPE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Admin_status	string	`json:"ADMIN_STATUS"`
	Oper_status	string	`json:"OPER_STATUS"`
	Tracker_status	string	`json:"TRACKER_STATUS"`
	Encap_type	string	`json:"ENCAP_TYPE"`
	Port_type	string	`json:"PORT_TYPE"`
	Mtu	string	`json:"MTU"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Speed	string	`json:"SPEED"`
	Duplex	string	`json:"DUPLEX"`
	Mss_adjust	string	`json:"MSS_ADJUST"`
	Uptime	string	`json:"UPTIME"`
	Rx_packets	string	`json:"RX_PACKETS"`
	Tx_packets	string	`json:"TX_PACKETS"`
}