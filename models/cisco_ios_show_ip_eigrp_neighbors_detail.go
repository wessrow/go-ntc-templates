package models

type CiscoIosShowIpEigrpNeighborsDetail struct {
	As	string	`json:"AS"`
	H	string	`json:"H"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Hold	string	`json:"HOLD"`
	Uptime	string	`json:"UPTIME"`
	Srtt	string	`json:"SRTT"`
	Rto	string	`json:"RTO"`
	Q_cnt	string	`json:"Q_CNT"`
	Seq_num	string	`json:"SEQ_NUM"`
	Version	string	`json:"VERSION"`
	Retrans	string	`json:"RETRANS"`
	Retries	string	`json:"RETRIES"`
	Prefixes	string	`json:"PREFIXES"`
	Topology_ids_from_peer	string	`json:"TOPOLOGY_IDS_FROM_PEER"`
	Max_neighbors	string	`json:"MAX_NEIGHBORS"`
	Current_neighbors	string	`json:"CURRENT_NEIGHBORS"`
}