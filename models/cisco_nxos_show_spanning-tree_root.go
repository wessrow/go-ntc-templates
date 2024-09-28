package models

type CiscoNxosShowSpanningTreeRoot struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Priority	string	`json:"PRIORITY"`
	Root_id	string	`json:"ROOT_ID"`
	Root_cost	string	`json:"ROOT_COST"`
	Hello_time	string	`json:"HELLO_TIME"`
	Max_age	string	`json:"MAX_AGE"`
	Fwd_dly	string	`json:"FWD_DLY"`
	Root_port	string	`json:"ROOT_PORT"`
}