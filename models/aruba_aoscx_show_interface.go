package models

type ArubaAoscxShowInterface struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Link_admin	string	`json:"LINK_ADMIN"`
	Link_state_info	string	`json:"LINK_STATE_INFO"`
	Link_transitions	string	`json:"LINK_TRANSITIONS"`
	Interface_desc	string	`json:"INTERFACE_DESC"`
	Hw_type	string	`json:"HW_TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Mtu	string	`json:"MTU"`
	If_type	string	`json:"IF_TYPE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Duplex	string	`json:"DUPLEX"`
	Qos_trust	string	`json:"QOS_TRUST"`
	Speed	string	`json:"SPEED"`
	Auto_neg	string	`json:"AUTO_NEG"`
	Flow_control	string	`json:"FLOW_CONTROL"`
	Error_control	string	`json:"ERROR_CONTROL"`
	Vlan_mode	string	`json:"VLAN_MODE"`
	Vlan_access	string	`json:"VLAN_ACCESS"`
	Vlan_native	string	`json:"VLAN_NATIVE"`
	Vlan_trunk	[]string	`json:"VLAN_TRUNK"`
	Aggregated_interfaces	[]string	`json:"AGGREGATED_INTERFACES"`
	Rx_total_packets	string	`json:"RX_TOTAL_PACKETS"`
	Rx_total_bytes	string	`json:"RX_TOTAL_BYTES"`
	Rx_unicast_packets	string	`json:"RX_UNICAST_PACKETS"`
	Rx_mcast_packets	string	`json:"RX_MCAST_PACKETS"`
	Rx_broadcast_packets	string	`json:"RX_BROADCAST_PACKETS"`
	Rx_errors	string	`json:"RX_ERRORS"`
	Rx_dropped	string	`json:"RX_DROPPED"`
	Rx_crc_fcs	string	`json:"RX_CRC_FCS"`
	Tx_total_packets	string	`json:"TX_TOTAL_PACKETS"`
	Tx_total_bytes	string	`json:"TX_TOTAL_BYTES"`
	Tx_unicast_packets	string	`json:"TX_UNICAST_PACKETS"`
	Tx_mcast_packets	string	`json:"TX_MCAST_PACKETS"`
	Tx_broadcast_packets	string	`json:"TX_BROADCAST_PACKETS"`
	Tx_errors	string	`json:"TX_ERRORS"`
	Tx_dropped	string	`json:"TX_DROPPED"`
	Tx_crc_fcs	string	`json:"TX_CRC_FCS"`
	Tx_collision	string	`json:"TX_COLLISION"`
}