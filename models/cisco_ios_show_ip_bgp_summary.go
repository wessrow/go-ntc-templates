package models

type CiscoIosShowIpBgpSummary struct {
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Address_family	string	`json:"ADDRESS_FAMILY"`
	Bgp_neighbor	string	`json:"BGP_NEIGHBOR"`
	Bgp_version	string	`json:"BGP_VERSION"`
	Neighbor_as	string	`json:"NEIGHBOR_AS"`
	Messages_received	string	`json:"MESSAGES_RECEIVED"`
	Messages_sent	string	`json:"MESSAGES_SENT"`
	Table_version	string	`json:"TABLE_VERSION"`
	Input_queue	string	`json:"INPUT_QUEUE"`
	Output_queue	string	`json:"OUTPUT_QUEUE"`
	Up_down	string	`json:"UP_DOWN"`
	State_or_prefixes_received	string	`json:"STATE_OR_PREFIXES_RECEIVED"`
}