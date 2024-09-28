package models

type CiscoNxosShowIpOspfDatabase struct {
	Router_id	string	`json:"ROUTER_ID"`
	Process_id	string	`json:"PROCESS_ID"`
	Area	string	`json:"AREA"`
	Link_id	string	`json:"LINK_ID"`
	Adv_router	string	`json:"ADV_ROUTER"`
	Age	string	`json:"AGE"`
	Link_count	string	`json:"LINK_COUNT"`
	Tag	string	`json:"TAG"`
}