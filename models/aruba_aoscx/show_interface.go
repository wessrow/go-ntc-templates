package aruba_aoscx 

type ShowInterface struct {
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

var ShowInterface_Template = `Value Required INTERFACE (\S+)
Value LINK_STATUS (\w+)
Value LINK_ADMIN (\w+)
Value LINK_STATE_INFO (\.*)
Value LINK_TRANSITIONS (\d+)
Value INTERFACE_DESC (.*)
Value HW_TYPE (\w+)
Value MAC_ADDRESS (\S+)
Value MTU (\d+)
Value IF_TYPE (\S+)
Value IP_ADDRESS (\S+)
Value DUPLEX (\w+)
Value QOS_TRUST (\S+)
Value SPEED (\S+\s\S+) 
Value AUTO_NEG (\S+)
Value FLOW_CONTROL (\S+)
Value ERROR_CONTROL (\S+)
Value VLAN_MODE (\S+)
Value VLAN_ACCESS (\S+)
Value VLAN_NATIVE (\S+)
Value List VLAN_TRUNK ([^,]+)
Value List AGGREGATED_INTERFACES (\S+)
Value RX_TOTAL_PACKETS (\d+)
Value RX_TOTAL_BYTES (\d+)
Value RX_UNICAST_PACKETS (\d+)
Value RX_MCAST_PACKETS (\d+)
Value RX_BROADCAST_PACKETS (\d+)
Value RX_ERRORS (\d+)
Value RX_DROPPED (\d+)
Value RX_CRC_FCS (\d+)
Value TX_TOTAL_PACKETS (\d+)
Value TX_TOTAL_BYTES (\d+)
Value TX_UNICAST_PACKETS (\d+)
Value TX_MCAST_PACKETS (\d+)
Value TX_BROADCAST_PACKETS (\d+)
Value TX_ERRORS (\d+)
Value TX_DROPPED (\d+)
Value TX_CRC_FCS (\S+)
Value TX_COLLISION (\d+)

Start
  ^\s*(Interface|Aggregate)\s+\S+\s+is -> Continue.Record
  ^\s*(Interface|Aggregate)\s+${INTERFACE}\s+is\s+${LINK_STATUS}
  ^\s*Admin\s*state\s*is\s*${LINK_ADMIN}
  ^\s*Link\s*transitions\s*:\s*${LINK_TRANSITIONS}
  ^\s*Description\s*:\s*${INTERFACE_DESC}
  ^\s*State\s*information\s*:\s*${LINK_STATE_INFO}
  ^\s*Hardware\s*:\s*${HW_TYPE},?\s*MAC\s*Address:\s+${MAC_ADDRESS}
  ^\s*MAC\s+Address\s*:\s*${MAC_ADDRESS}
  ^\s*IPv4\s+address\s+${IP_ADDRESS}
  ^\s*MTU\s*${MTU}
  ^\s*Type\s*${IF_TYPE}
  ^\s*${DUPLEX}-duplex
  ^\s*qos\s*trust\s*${QOS_TRUST}
  ^\s*Speed\s*:*\s*${SPEED}
  ^\s*Auto-negotiation\s*is\s*${AUTO_NEG}
  ^\s*Flow-control\s*:\s*${FLOW_CONTROL}
  ^\s*Error-control\s*:\s*${ERROR_CONTROL}
  ^\s*VLAN Mode\s*:\s*${VLAN_MODE}
  ^\s*Access VLAN\s*:\s*${VLAN_ACCESS}
  ^\s*Native VLAN\s*:\s*${VLAN_NATIVE}
  ^\s*Allowed VLAN List\s*:\s+${VLAN_TRUNK},* -> Continue
  ^\s*Allowed VLAN List\s*:\s+(?:[^,]+,){1}${VLAN_TRUNK},* -> Continue
  ^\s*Allowed VLAN List\s*:\s+(?:[^,]+,){2}${VLAN_TRUNK},* -> Continue
  ^\s*Allowed VLAN List\s*:\s+(?:[^,]+,){3}${VLAN_TRUNK},* -> Continue
  ^\s*Allowed VLAN List\s*:\s+(?:[^,]+,){4}${VLAN_TRUNK},* -> Continue
  ^\s*Allowed VLAN List\s*:\s+(?:[^,]+,){5}${VLAN_TRUNK},* -> Continue
  ^\s*Allowed VLAN List\s*:
  ^\s*Link state
  ^\s*Aggregated-interfaces\s*:\s+${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:\s+(?:\S+\s+){1}${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:\s+(?:\S+\s+){2}${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:\s+(?:\S+\s+){3}${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:\s+(?:\S+\s+){4}${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:\s+(?:\S+\s+){5}${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:\s+(?:\S+\s+){6}${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:\s+(?:\S+\s+){7}${AGGREGATED_INTERFACES}\s* -> Continue
  ^\s*Aggregated-interfaces\s*:
  ^\s*Aggregation-key
  ^\s*Aggregate\s+mode
  ^\s*Persona
  ^\s*rate-limit\s+broadcast
  ^\s*L3\s+Counters
  ^\s*MDI mode
  ^\s*Rate collection interval
  ^\s*Energy-Efficient
  ^\s*Rx -> Rx
  ^\s*Rate\s+RX\s+TX\s+Total -> Rate
  ^\s*Statistic\s+RX\s+TX\s+Total -> Statistic
  ^. -> Error

Rx
  ^\s*${RX_TOTAL_PACKETS} total packets\s*${RX_TOTAL_BYTES} total bytes
  ^\s*${RX_UNICAST_PACKETS}   unicast packets
  ^\s*${RX_MCAST_PACKETS}   multicast packets
  ^\s*${RX_BROADCAST_PACKETS}   broadcast packets
  ^\s*${RX_ERRORS} errors\s*${RX_DROPPED} dropped
  ^\s*${RX_CRC_FCS} CRC/FCS
  ^\s*Tx -> Tx
  ^. -> Error

Tx
  ^\s*${TX_TOTAL_PACKETS} total packets\s*${TX_TOTAL_BYTES} total bytes
  ^\s*${TX_UNICAST_PACKETS}   unicast packets
  ^\s*${TX_MCAST_PACKETS}   multicast packets
  ^\s*${TX_BROADCAST_PACKETS}   broadcast packets
  ^\s*${TX_ERRORS} errors\s*${TX_DROPPED} dropped
  ^\s*${TX_CRC_FCS} CRC/FCS
  ^\s*${TX_COLLISION}\s*collision
  ^\s*$$ -> Start
  ^. -> Error

Rate
  ^\s*----
  ^\s*Mbits\s+/\s+sec
  ^\s*KPkts\s+/\s+sec
  ^\s+Unicast
  ^\s+Multicast
  ^\s+Broadcast
  ^\s+Utilization
  ^\s*$$ -> Statistic
  ^. -> Error

Statistic
  ^\s*Statistic\s+RX\s+TX\s+Total
  ^\s*----
  ^\s*Packets\s+${RX_TOTAL_PACKETS}\s+${TX_TOTAL_PACKETS}\s+\d+
  ^\s+Unicast\s+${RX_UNICAST_PACKETS}\s+${TX_UNICAST_PACKETS}\s+\d+
  ^\s+Multicast\s+${RX_MCAST_PACKETS}\s+${TX_MCAST_PACKETS}\s+\d+
  ^\s+Broadcast\s+${RX_BROADCAST_PACKETS}\s+${TX_BROADCAST_PACKETS}\s+\d+
  ^\s*Bytes\s+${RX_TOTAL_BYTES}\s+${TX_TOTAL_BYTES}\s+\d+
  ^\s*Jumbos
  ^\s*Pause\s+Frames
  ^\s*Dropped\s+${RX_DROPPED}\s+${TX_DROPPED}\s+\d+
  ^\s*Errors\s+${RX_ERRORS}\s+${TX_ERRORS}\s+\d+
  ^\s+CRC/FCS\s+${RX_CRC_FCS}\s+${TX_CRC_FCS}\s+\d+
  ^\s+Collision\s+\S+\s+${TX_COLLISION}\s+\d+
  ^\s*Runts
  ^\s*Giants
  ^\s*L3\s+Packets
  ^\s*$$ -> Start
  ^. -> Error
`