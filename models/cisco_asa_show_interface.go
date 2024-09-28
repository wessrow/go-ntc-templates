package models

type CiscoAsaShowInterface struct {
	Interface	string	`json:"INTERFACE"`
	Interface_zone	string	`json:"INTERFACE_ZONE"`
	Link_status	string	`json:"LINK_STATUS"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Bandwidth	string	`json:"BANDWIDTH"`
	Delay	string	`json:"DELAY"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Description	string	`json:"DESCRIPTION"`
	Bonded_port	string	`json:"BONDED_PORT"`
	Bonded_port_members	string	`json:"BONDED_PORT_MEMBERS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Mtu	string	`json:"MTU"`
	Vlan_id	string	`json:"VLAN_ID"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Onemin_in_pps	string	`json:"ONEMIN_IN_PPS"`
	Onemin_in_rate	string	`json:"ONEMIN_IN_RATE"`
	Onemin_out_pps	string	`json:"ONEMIN_OUT_PPS"`
	Onemin_out_rate	string	`json:"ONEMIN_OUT_RATE"`
	Onemin_drop_rate	string	`json:"ONEMIN_DROP_RATE"`
	Fivemin_in_pps	string	`json:"FIVEMIN_IN_PPS"`
	Fivemin_in_rate	string	`json:"FIVEMIN_IN_RATE"`
	Fivemin_out_pps	string	`json:"FIVEMIN_OUT_PPS"`
	Fivemin_out_rate	string	`json:"FIVEMIN_OUT_RATE"`
	Fivemin_drop_rate	string	`json:"FIVEMIN_DROP_RATE"`
	Tunnel_source_interface	string	`json:"TUNNEL_SOURCE_INTERFACE"`
	Tunnel_source_ip	string	`json:"TUNNEL_SOURCE_IP"`
	Tunnel_destination_ip	string	`json:"TUNNEL_DESTINATION_IP"`
	Tunnel_mode	string	`json:"TUNNEL_MODE"`
	Tunnel_ipsec_profile	string	`json:"TUNNEL_IPSEC_PROFILE"`
}