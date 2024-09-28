package models

type CiscoIosShowStandby struct {
	Interface	string	`json:"INTERFACE"`
	Group	string	`json:"GROUP"`
	Version	string	`json:"VERSION"`
	State	string	`json:"STATE"`
	State_changes	string	`json:"STATE_CHANGES"`
	State_last_change	string	`json:"STATE_LAST_CHANGE"`
	Virtual_ip	string	`json:"VIRTUAL_IP"`
	Secondary_ips	[]string	`json:"SECONDARY_IPS"`
	Active_virtual_mac	string	`json:"ACTIVE_VIRTUAL_MAC"`
	Local_virtual_mac	string	`json:"LOCAL_VIRTUAL_MAC"`
	Hello_time	string	`json:"HELLO_TIME"`
	Hold_time	string	`json:"HOLD_TIME"`
	Authentication	string	`json:"AUTHENTICATION"`
	Preemption	string	`json:"PREEMPTION"`
	Active_router	string	`json:"ACTIVE_ROUTER"`
	Active_router_priority	string	`json:"ACTIVE_ROUTER_PRIORITY"`
	Active_router_mac	string	`json:"ACTIVE_ROUTER_MAC"`
	Standby_router	string	`json:"STANDBY_ROUTER"`
	Standby_router_priority	string	`json:"STANDBY_ROUTER_PRIORITY"`
	Priority	string	`json:"PRIORITY"`
	Group_name	string	`json:"GROUP_NAME"`
	Flags	string	`json:"FLAGS"`
	Track_item	string	`json:"TRACK_ITEM"`
	Track_type	string	`json:"TRACK_TYPE"`
	Track_state	string	`json:"TRACK_STATE"`
	Track_decrement_time	string	`json:"TRACK_DECREMENT_TIME"`
}