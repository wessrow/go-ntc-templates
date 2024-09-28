package models

type CiscoIosShowRunningConfigPartitionAccessList struct {
	Acl_name	string	`json:"ACL_NAME"`
	Acl_type	string	`json:"ACL_TYPE"`
	Line_num	string	`json:"LINE_NUM"`
	Comment	[]string	`json:"COMMENT"`
	Action	string	`json:"ACTION"`
	Protocol	string	`json:"PROTOCOL"`
	Src_host	string	`json:"SRC_HOST"`
	Src_any	string	`json:"SRC_ANY"`
	Src_network	string	`json:"SRC_NETWORK"`
	Src_wildcard	string	`json:"SRC_WILDCARD"`
	Src_port_match	string	`json:"SRC_PORT_MATCH"`
	Src_port	string	`json:"SRC_PORT"`
	Src_port_range_start	string	`json:"SRC_PORT_RANGE_START"`
	Src_port_range_end	string	`json:"SRC_PORT_RANGE_END"`
	Dst_host	string	`json:"DST_HOST"`
	Dst_any	string	`json:"DST_ANY"`
	Dst_network	string	`json:"DST_NETWORK"`
	Dst_wildcard	string	`json:"DST_WILDCARD"`
	Dst_port_match	string	`json:"DST_PORT_MATCH"`
	Dst_port	string	`json:"DST_PORT"`
	Dst_port_range_start	string	`json:"DST_PORT_RANGE_START"`
	Dst_port_range_end	string	`json:"DST_PORT_RANGE_END"`
	Flags_match	string	`json:"FLAGS_MATCH"`
	Tcp_flag	string	`json:"TCP_FLAG"`
	Log	string	`json:"LOG"`
	Icmp_type	string	`json:"ICMP_TYPE"`
	Time	string	`json:"TIME"`
}