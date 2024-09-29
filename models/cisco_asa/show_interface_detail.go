package cisco_asa 

type ShowInterfaceDetail struct {
	Interface	string	`json:"INTERFACE"`
	Interface_name	string	`json:"INTERFACE_NAME"`
	Link_status	string	`json:"LINK_STATUS"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Bandwidth	string	`json:"BANDWIDTH"`
	Delay	string	`json:"DELAY"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Flowcontrol_in	string	`json:"FLOWCONTROL_IN"`
	Flowcontrol_out	string	`json:"FLOWCONTROL_OUT"`
	Bonded_port	string	`json:"BONDED_PORT"`
	Description	string	`json:"DESCRIPTION"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Mtu	string	`json:"MTU"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	In_packets	string	`json:"IN_PACKETS"`
	In_bytes	string	`json:"IN_BYTES"`
	No_buffer	string	`json:"NO_BUFFER"`
	Broadcasts	string	`json:"BROADCASTS"`
	Runts	string	`json:"RUNTS"`
	Giants	string	`json:"GIANTS"`
	In_errors	string	`json:"IN_ERRORS"`
	Crc	string	`json:"CRC"`
	Frame	string	`json:"FRAME"`
	Overrun	string	`json:"OVERRUN"`
	Ignored	string	`json:"IGNORED"`
	Abort	string	`json:"ABORT"`
	In_pause_resume	string	`json:"IN_PAUSE_RESUME"`
	In_pause	string	`json:"IN_PAUSE"`
	In_resume	string	`json:"IN_RESUME"`
	L2_decode_drops	string	`json:"L2_DECODE_DROPS"`
	Switch_ingress_policy_drops	string	`json:"SWITCH_INGRESS_POLICY_DROPS"`
	Out_packets	string	`json:"OUT_PACKETS"`
	Out_bytes	string	`json:"OUT_BYTES"`
	Underruns	string	`json:"UNDERRUNS"`
	Out_pause_resume	string	`json:"OUT_PAUSE_RESUME"`
	Out_pause	string	`json:"OUT_PAUSE"`
	Out_resume	string	`json:"OUT_RESUME"`
	Out_errors	string	`json:"OUT_ERRORS"`
	Collisions	string	`json:"COLLISIONS"`
	Interface_resets	string	`json:"INTERFACE_RESETS"`
	Late_collisions	string	`json:"LATE_COLLISIONS"`
	Deferred	string	`json:"DEFERRED"`
	Out_decode_drops	string	`json:"OUT_DECODE_DROPS"`
	Rate_limit_drops	string	`json:"RATE_LIMIT_DROPS"`
	Switch_egress_policy_drops	string	`json:"SWITCH_EGRESS_POLICY_DROPS"`
	In_reset_drops	string	`json:"IN_RESET_DROPS"`
	Out_reset_drops	string	`json:"OUT_RESET_DROPS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Control_interface_number	string	`json:"CONTROL_INTERFACE_NUMBER"`
	Control_interface_config_status	string	`json:"CONTROL_INTERFACE_CONFIG_STATUS"`
	Control_interface_state	string	`json:"CONTROL_INTERFACE_STATE"`
	Rxring_number	[]string	`json:"RXRING_NUMBER"`
	Rxring_packets	[]string	`json:"RXRING_PACKETS"`
	Rxring_bytes	[]string	`json:"RXRING_BYTES"`
	Rxring_overrun	[]string	`json:"RXRING_OVERRUN"`
	Rxring_blocks_free_curr	[]string	`json:"RXRING_BLOCKS_FREE_CURR"`
	Rxring_blocks_free_low	[]string	`json:"RXRING_BLOCKS_FREE_LOW"`
	Rxring_interface	[]string	`json:"RXRING_INTERFACE"`
	Txring_number	[]string	`json:"TXRING_NUMBER"`
	Txring_packets	[]string	`json:"TXRING_PACKETS"`
	Txring_bytes	[]string	`json:"TXRING_BYTES"`
	Txring_underrun	[]string	`json:"TXRING_UNDERRUN"`
	Txring_blocks_free_curr	[]string	`json:"TXRING_BLOCKS_FREE_CURR"`
	Txring_blocks_free_low	[]string	`json:"TXRING_BLOCKS_FREE_LOW"`
	Txring_interface	[]string	`json:"TXRING_INTERFACE"`
}

var ShowInterfaceDetail_Template = `Value Required INTERFACE (\S+)
Value INTERFACE_NAME (.*)
Value LINK_STATUS (.*)
Value PROTOCOL_STATUS (.*)
Value HARDWARE_TYPE ([\w ]+)
Value BANDWIDTH (\d+\s+\w+)
Value DELAY (\d+\s+\w+)
Value DUPLEX ([Ff][Uu][Ll]{2}|[Hh][Aa][Ll][Ff]|[Aa][Uu][Tt][Oo])
Value SPEED (\d+\s*\S+|[Aa][Uu][Tt][Oo])
Value FLOWCONTROL_IN (\w+)
Value FLOWCONTROL_OUT (\w+)
Value BONDED_PORT (.+?)
Value DESCRIPTION (.*)
Value MAC_ADDRESS ([a-zA-Z0-9]+\.[a-zA-Z0-9]+\.[a-zA-Z0-9]+)
Value MTU (.+)
Value IP_ADDRESS (.+?)
Value NETMASK (.+?)
Value IN_PACKETS (\d+)
Value IN_BYTES (\d+)
Value NO_BUFFER (\d+)
Value BROADCASTS (\d+)
Value RUNTS (\d+)
Value GIANTS (\d+)
Value IN_ERRORS (\d+)
Value CRC (\d+)
Value FRAME (\d+)
Value OVERRUN (\d+)
Value IGNORED (\d+)
Value ABORT (\d+)
Value IN_PAUSE_RESUME (\d+)
Value IN_PAUSE (\d+)
Value IN_RESUME (\d+)
Value L2_DECODE_DROPS (\d+)
Value SWITCH_INGRESS_POLICY_DROPS (\d+)
Value OUT_PACKETS (\d+)
Value OUT_BYTES (\d+)
Value UNDERRUNS (\d+)
Value OUT_PAUSE_RESUME (\d+)
Value OUT_PAUSE (\d+)
Value OUT_RESUME (\d+)
Value OUT_ERRORS (\d+)
Value COLLISIONS (\d+)
Value INTERFACE_RESETS (\d+)
Value LATE_COLLISIONS (\d+)
Value DEFERRED (\d+)
Value OUT_DECODE_DROPS (\d+)
Value RATE_LIMIT_DROPS (\d+)
Value SWITCH_EGRESS_POLICY_DROPS (\d+)
Value IN_RESET_DROPS (\d+)
Value OUT_RESET_DROPS (\d+)
Value VLAN_ID (\d+)
#
# Control Point Interface States
#
Value CONTROL_INTERFACE_NUMBER (\d+)
Value CONTROL_INTERFACE_CONFIG_STATUS (.+)
Value CONTROL_INTERFACE_STATE (.+)
#
# Queue Stats
#
# RXRING
Value List RXRING_NUMBER (\d+)
Value List RXRING_PACKETS (\d+)
Value List RXRING_BYTES (\d+)
Value List RXRING_OVERRUN (\d+)
Value List RXRING_BLOCKS_FREE_CURR (\d+)
Value List RXRING_BLOCKS_FREE_LOW (\d+)
Value List RXRING_INTERFACE (.+?)
#
# TXRING
Value List TXRING_NUMBER (\d+)
Value List TXRING_PACKETS (\d+)
Value List TXRING_BYTES (\d+)
Value List TXRING_UNDERRUN (\d+)
Value List TXRING_BLOCKS_FREE_CURR (\d+)
Value List TXRING_BLOCKS_FREE_LOW (\d+)
Value List TXRING_INTERFACE (.+?)


Start
  ^Interface -> Continue.Record
  ^Interface\s+${INTERFACE}\s+"${INTERFACE_NAME}",\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS}
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE}.*BW\s+${BANDWIDTH}.*DLY\s+${DELAY}
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE}.*MAC\s+address\s+${MAC_ADDRESS},\s+MTU\s+${MTU}
  ^\s+\S*${DUPLEX}\S*,\s+\S*?${SPEED}(?:\)|\S+)\s*$$
  ^\s+Input\s+flow\s+control\s+is\s+${FLOWCONTROL_IN},\s+output\s+flow\s+control\s+is\s+${FLOWCONTROL_OUT}
  ^\s+Active\s+member\s+of\s+${BONDED_PORT}\s*$$
  ^\s+Description:\s+${DESCRIPTION}
  ^\s+MAC\s+address\s+${MAC_ADDRESS}.*MTU\s+${MTU}\s*$$
  ^\s+IP\s+[Aa]ddress\s+${IP_ADDRESS}(?:,\s+[Ss]ubnet\s+[Mm]ask\s+${NETMASK})?\s*$$
  ^\s+${IN_PACKETS}\s+packets\s+input,\s+${IN_BYTES}\s+bytes,\s+${NO_BUFFER}\s+no\s+buffer
  ^\s+${IN_PACKETS}\s+packets\s+input,\s+${IN_BYTES}\s+bytes
  ^\s+Received\s+${BROADCASTS}\s+broadcasts,\s+${RUNTS}\s+runts,\s+${GIANTS}\s+giants
  ^\s+${IN_ERRORS}\s+input\s+errors,\s+${CRC}\s+CRC,\s+${FRAME}\s+frame,\s+${OVERRUN}\s+overrun,\s+${IGNORED}\s+ignored,\s+${ABORT}\s+abort
  ^\s+${IN_PAUSE_RESUME}\s+pause/resume\s+input
  ^\s+${IN_PAUSE}\s+pause\s+input,\s+${IN_RESUME}\s+resume\s+input
  ^\s+${L2_DECODE_DROPS}\s+L2\s+decode\s+drops
  ^\s+${SWITCH_INGRESS_POLICY_DROPS}\s+switch\s+ingress\s+policy\s+drops
  ^\s+${OUT_PACKETS}\s+packets\s+output,\s+${OUT_BYTES}\s+bytes,\s+${UNDERRUNS}\s+underruns
  ^\s+${OUT_PACKETS}\s+packets\s+output,\s+${OUT_BYTES}\s+bytes
  ^\s+${OUT_PAUSE_RESUME}\s+pause/resume\s+output
  ^\s+${OUT_PAUSE}\s+pause\s+output,\s+${OUT_RESUME}\s+resume\s+output
  ^\s+${OUT_ERRORS}\s+output\s+errors,\s+${COLLISIONS}\s+collisions,\s+${INTERFACE_RESETS}\s+interface\s+resets
  ^\s+${LATE_COLLISIONS}\s+late\s+collisions,\s+${DEFERRED}\s+deferred
  ^\s+${OUT_DECODE_DROPS}\s+output\s+decode\s+drops
  ^\s+${RATE_LIMIT_DROPS}\s+rate\s+limit\s+drops
  ^\s+${SWITCH_EGRESS_POLICY_DROPS}\s+switch\s+egress\s+policy\s+drops
  ^\s+${IN_RESET_DROPS}\s+input\s+reset\s+drops,\s+${OUT_RESET_DROPS}\s+output\s+reset\s+drops
  ^\s+input\s+queue
  ^\s+output\s+queue
  ^\s+Queue\s+Stats:
  ^\s+Traffic\s+Statistics
  ^\s+\d+\s+packets\s+dropped
  ^\s+\d+\s+minute\s+\w+\s+rate
  ^\s+VLAN\s+identifier\s+${VLAN_ID}
  ^\s+Control\s+Point\s+Interface\s+States -> ControlPoint
  ^\s+RX\[${RXRING_NUMBER}\]:\s+${RXRING_PACKETS}\s+packets,\s+${RXRING_BYTES}\s+bytes,\s+${RXRING_OVERRUN}\s+overrun -> RXRingBlocks
  ^\s+TX\[${TXRING_NUMBER}\]:\s+${TXRING_PACKETS}\s+packets,\s+${TXRING_BYTES}\s+bytes,\s+${TXRING_UNDERRUN}\s+underruns -> TXRingBlocks
  ^\s+Available\s+for\s+allocation\s+to\s+a\s+context
  ^Interface -> Continue.Record
  ^Interface\s+${INTERFACE}\s+"${INTERFACE_NAME}",\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS}
  ^\s+Topology\s+Information: -> Topology
  ^\s*$$
  ^. -> Error

ControlPoint
  ^\s+Interface\s+number\s+is\s+${CONTROL_INTERFACE_NUMBER}
  ^\s+Interface\s+config\s+status\s+is\s+${CONTROL_INTERFACE_CONFIG_STATUS}
  ^\s+Interface\s+state\s+is\s+${CONTROL_INTERFACE_STATE}
  ^Interface -> Continue.Record
  ^Interface\s+${INTERFACE}\s+"${INTERFACE_NAME}",\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS} -> Start
  ^\s+Queue\s+Stats:
  ^\s+RX\[${RXRING_NUMBER}\]:\s+${RXRING_PACKETS}\s+packets,\s+${RXRING_BYTES}\s+bytes,\s+${RXRING_OVERRUN}\s+overrun -> RXRingBlocks
  ^\s+TX\[${TXRING_NUMBER}\]:\s+${TXRING_PACKETS}\s+packets,\s+${TXRING_BYTES}\s+bytes,\s+${TXRING_UNDERRUN}\s+underruns -> TXRingBlocks

RXRingBlocks
  ^\s+RX\[${RXRING_NUMBER}\]:\s+${RXRING_PACKETS}\s+packets,\s+${RXRING_BYTES}\s+bytes,\s+${RXRING_OVERRUN}\s+overrun
  ^\s+Blocks\s+free\s+curr/low:\s+${RXRING_BLOCKS_FREE_CURR}/${RXRING_BLOCKS_FREE_LOW}
  ^\s+Used\s+by\s+${RXRING_INTERFACE}\s*$$
  ^\s+TX\[${TXRING_NUMBER}\]:\s+${TXRING_PACKETS}\s+packets,\s+${TXRING_BYTES}\s+bytes,\s+${TXRING_UNDERRUN}\s+underruns -> TXRingBlocks
  ^\s+Topology\s+Information: -> Topology
  ^\s+Control\s+Point\s+Interface\s+States -> ControlPoint
  ^Interface -> Continue.Record
  ^Interface\s+${INTERFACE}\s+"${INTERFACE_NAME}",\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS} -> Start
  ^\s*$$
  ^. -> Error

TXRingBlocks
  ^\s+TX\[${TXRING_NUMBER}\]:\s+${TXRING_PACKETS}\s+packets,\s+${TXRING_BYTES}\s+bytes,\s+${TXRING_UNDERRUN}\s+underruns
  ^\s+Blocks\s+free\s+curr/low:\s+${TXRING_BLOCKS_FREE_CURR}/${TXRING_BLOCKS_FREE_LOW}
  ^\s+Used\s+by\s+${TXRING_INTERFACE}\s*$$
  ^\s+Topology\s+Information: -> Topology
  ^\s+RX\[${RXRING_NUMBER}\]:\s+${RXRING_PACKETS}\s+packets,\s+${RXRING_BYTES}\s+bytes,\s+${RXRING_OVERRUN}\s+overrun -> RXRingBlocks
  ^\s+Control\s+Point\s+Interface\s+States -> ControlPoint
  ^Interface -> Continue.Record
  ^Interface\s+${INTERFACE}\s+"${INTERFACE_NAME}",\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS} -> Start
  ^\s*$$
  ^. -> Error

Topology
  ^\s{4,}\S+
  ^\s+Control\s+Point\s+Interface\s+States -> ControlPoint
  ^Interface -> Continue.Record
  ^Interface\s+${INTERFACE}\s+"${INTERFACE_NAME}",\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS} -> Start
  ^\s+RX\[${RXRING_NUMBER}\]:\s+${RXRING_PACKETS}\s+packets,\s+${RXRING_BYTES}\s+bytes,\s+${RXRING_OVERRUN}\s+overrun -> RXRingBlocks
  ^\s+TX\[${TXRING_NUMBER}\]:\s+${TXRING_PACKETS}\s+packets,\s+${TXRING_BYTES}\s+bytes,\s+${TXRING_UNDERRUN}\s+underruns -> TXRingBlocks
  ^\s*$$
  ^. -> Error
`