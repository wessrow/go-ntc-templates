package models

type CiscoIosShowIpFlowToptalkers struct {
	Src_intf	string	`json:"SRC_INTF"`
	Src_ipaddr	string	`json:"SRC_IPADDR"`
	Src_port	string	`json:"SRC_PORT"`
	Dst_intf	string	`json:"DST_INTF"`
	Dst_ipaddr	string	`json:"DST_IPADDR"`
	Proto	string	`json:"PROTO"`
	Dst_port	string	`json:"DST_PORT"`
	Pkt	string	`json:"PKT"`
}