package models

type CiscoIosShowIpEigrpTopology struct {
	Process_id	string	`json:"PROCESS_ID"`
	Router_id	string	`json:"ROUTER_ID"`
	Code	string	`json:"CODE"`
	Network	string	`json:"NETWORK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Successors	string	`json:"SUCCESSORS"`
	Fd	string	`json:"FD"`
	Tag	string	`json:"TAG"`
	Adv_router	[]string	`json:"ADV_ROUTER"`
	Adv_fd	[]string	`json:"ADV_FD"`
	Adv_rd	[]string	`json:"ADV_RD"`
	Out_interface	[]string	`json:"OUT_INTERFACE"`
	Source	string	`json:"SOURCE"`
}