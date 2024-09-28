package models

type CiscoNxosShowIpEigrpNeighbors struct {
	As	string	`json:"AS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Hold	string	`json:"HOLD"`
	Uptime	string	`json:"UPTIME"`
	Srtt	string	`json:"SRTT"`
	Rto	string	`json:"RTO"`
	Q_cnt	string	`json:"Q_CNT"`
	Seq_num	string	`json:"SEQ_NUM"`
}