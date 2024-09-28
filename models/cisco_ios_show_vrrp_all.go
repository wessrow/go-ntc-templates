package models

type CiscoIosShowVrrpAll struct {
	Interface	string	`json:"INTERFACE"`
	Group	string	`json:"GROUP"`
	Addr_family	string	`json:"ADDR_FAMILY"`
	State	string	`json:"STATE"`
	State_duration	string	`json:"STATE_DURATION"`
	Virtual_ip_address	string	`json:"VIRTUAL_IP_ADDRESS"`
	Virtual_mac_address	string	`json:"VIRTUAL_MAC_ADDRESS"`
	Adv_interval	string	`json:"ADV_INTERVAL"`
	Preempt	string	`json:"PREEMPT"`
	Priority	string	`json:"PRIORITY"`
	Priority_configured	string	`json:"PRIORITY_CONFIGURED"`
	Track_obj	string	`json:"TRACK_OBJ"`
	Track_status	string	`json:"TRACK_STATUS"`
	Track_action	string	`json:"TRACK_ACTION"`
	Master_ip_address	string	`json:"MASTER_IP_ADDRESS"`
	Master_priority	string	`json:"MASTER_PRIORITY"`
	Master_adv_interval	string	`json:"MASTER_ADV_INTERVAL"`
	Master_down_interval	string	`json:"MASTER_DOWN_INTERVAL"`
}