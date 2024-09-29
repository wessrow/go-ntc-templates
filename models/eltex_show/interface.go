package eltex_show 

type Interface struct {
	Interface	string	`json:"INTERFACE"`
	Status	string	`json:"STATUS"`
	Iface_index	string	`json:"IFACE_INDEX"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Description	string	`json:"DESCRIPTION"`
	Mtu	string	`json:"MTU"`
	Duplex_type	string	`json:"DUPLEX_TYPE"`
	Speed	string	`json:"SPEED"`
	Link_type	string	`json:"LINK_TYPE"`
	Media_type	string	`json:"MEDIA_TYPE"`
	Uptime	string	`json:"UPTIME"`
	Downtime	string	`json:"DOWNTIME"`
	Link_aggregation_type	string	`json:"LINK_AGGREGATION_TYPE"`
	Link_modes	[]string	`json:"LINK_MODES"`
	Flow_control_status	string	`json:"FLOW_CONTROL_STATUS"`
	Fec_status	string	`json:"FEC_STATUS"`
	Mdix_mode_status	string	`json:"MDIX_MODE_STATUS"`
	Iface_in_po_num	string	`json:"IFACE_IN_PO_NUM"`
	Active_iface_in_po_num	string	`json:"ACTIVE_IFACE_IN_PO_NUM"`
	Min_iface_required_in_po_num	string	`json:"MIN_IFACE_REQUIRED_IN_PO_NUM"`
	Iface_in_po	[]string	`json:"IFACE_IN_PO"`
	Iface_in_po_bandwidth	[]string	`json:"IFACE_IN_PO_BANDWIDTH"`
	Iface_in_po_status	[]string	`json:"IFACE_IN_PO_STATUS"`
	Sum_bandwidth	string	`json:"SUM_BANDWIDTH"`
	Input_rate	string	`json:"INPUT_RATE"`
	Output_rate	string	`json:"OUTPUT_RATE"`
	Input_packets	string	`json:"INPUT_PACKETS"`
	Bytes_received	string	`json:"BYTES_RECEIVED"`
	Broadcasts_num	string	`json:"BROADCASTS_NUM"`
	Multicasts_num	string	`json:"MULTICASTS_NUM"`
	Input_err_num	string	`json:"INPUT_ERR_NUM"`
	Fcs_num	string	`json:"FCS_NUM"`
	Alignments_num	string	`json:"ALIGNMENTS_NUM"`
	Oversize_num	string	`json:"OVERSIZE_NUM"`
	Internal_mac_num	string	`json:"INTERNAL_MAC_NUM"`
	Pause_frames_received	string	`json:"PAUSE_FRAMES_RECEIVED"`
	Output_packets	string	`json:"OUTPUT_PACKETS"`
	Bytes_sent	string	`json:"BYTES_SENT"`
	Output_err_num	string	`json:"OUTPUT_ERR_NUM"`
	Collisions_num	string	`json:"COLLISIONS_NUM"`
	Excessive_collisions_num	string	`json:"EXCESSIVE_COLLISIONS_NUM"`
	Late_collisions_num	string	`json:"LATE_COLLISIONS_NUM"`
	Pause_frames_sent	string	`json:"PAUSE_FRAMES_SENT"`
	Symbol_err_num	string	`json:"SYMBOL_ERR_NUM"`
	Carrier_num	string	`json:"CARRIER_NUM"`
	Sqe_test_err_num	string	`json:"SQE_TEST_ERR_NUM"`
	Output_queue_num	[]string	`json:"OUTPUT_QUEUE_NUM"`
	Output_queue_packets_passed	[]string	`json:"OUTPUT_QUEUE_PACKETS_PASSED"`
	Output_queue_packets_dropped	[]string	`json:"OUTPUT_QUEUE_PACKETS_DROPPED"`
}

var Interface_Template = `Value INTERFACE ([a-zA-Z\-]+\s*(?:\d+(?:/\d+)*)?)
Value STATUS (.*)
Value IFACE_INDEX (\d+)
Value MAC_ADDRESS ([a-zA-Z0-9]{2}(:[a-zA-Z0-9]{2}){5})
Value DESCRIPTION (.*)
Value MTU (\d+)
Value DUPLEX_TYPE (\S+)
Value SPEED (\S+)
Value LINK_TYPE (.*)
Value MEDIA_TYPE (\S+)
Value UPTIME (.*)
Value DOWNTIME (.*)
Value LINK_AGGREGATION_TYPE (.*)
Value List LINK_MODES (\d+baseT/(?:Full|Half))
Value FLOW_CONTROL_STATUS (\S+)
Value FEC_STATUS (\S+)
Value MDIX_MODE_STATUS (\S+)
Value IFACE_IN_PO_NUM (\d+)
Value ACTIVE_IFACE_IN_PO_NUM (\d+)
Value MIN_IFACE_REQUIRED_IN_PO_NUM (\d+)
Value List IFACE_IN_PO ([a-zA-Z\-]+\s*(?:\d+(?:/\d+)*)?)
Value List IFACE_IN_PO_BANDWIDTH (\d+[a-zA-Z]+)
Value List IFACE_IN_PO_STATUS (.*)
Value SUM_BANDWIDTH ([a-zA-Z0-9]+)
Value INPUT_RATE (-?\d+)
Value OUTPUT_RATE (-?\d+)
Value INPUT_PACKETS (-?\d+)
Value BYTES_RECEIVED (-?\d+)
Value BROADCASTS_NUM (-?\d+)
Value MULTICASTS_NUM (-?\d+)
Value INPUT_ERR_NUM (-?\d+)
Value FCS_NUM (-?\d+)
Value ALIGNMENTS_NUM (-?\d+)
Value OVERSIZE_NUM (-?\d+)
Value INTERNAL_MAC_NUM (-?\d+)
Value PAUSE_FRAMES_RECEIVED (-?\d+)
Value OUTPUT_PACKETS (-?\d+)
Value BYTES_SENT (-?\d+)
Value OUTPUT_ERR_NUM (-?\d+)
Value COLLISIONS_NUM (-?\d+)
Value EXCESSIVE_COLLISIONS_NUM (-?\d+)
Value LATE_COLLISIONS_NUM (-?\d+)
Value PAUSE_FRAMES_SENT (-?\d+)
Value SYMBOL_ERR_NUM (-?\d+)
Value CARRIER_NUM (-?\d+)
Value SQE_TEST_ERR_NUM (-?\d+)
Value List OUTPUT_QUEUE_NUM (\d+)
Value List OUTPUT_QUEUE_PACKETS_PASSED (\d+)
Value List OUTPUT_QUEUE_PACKETS_DROPPED (\d+)

Start
  ^\s*Load\s+balancing\s*:.*$$
  ^\s*Gathering\s+information...\s*$$
  ^\s*Channel\s+Ports\s*$$ -> Columns2
  ^\s*-+\s*show\s+interfaces\s+\S+\s*-+\s*$$ -> Continue.Record
  ^\s*-+\s*show\s+interfaces\s+\S+\s*-+\s*$$
  ^\s*${INTERFACE}\s+is\s+${STATUS}\s+\(.*\)\s*$$
  ^\s*Interface\s+index\s+is\s+${IFACE_INDEX}\s*$$
  ^\s*.*,\s*MAC\s+address\s+is\s+${MAC_ADDRESS}\s*$$
  ^\s*Description:\s*${DESCRIPTION}\s*$$
  ^\s*Interface\s+MTU\s+is\s+${MTU}\s*$$
  ^\s*Port\s+is\s+\S+\s*$$
  ^\s*${DUPLEX_TYPE},\s+${SPEED},\s+link\s+type\s+is\s+${LINK_TYPE},\s+media\s+type\s+is\s+${MEDIA_TYPE}\s*$$
  ^\s*Link\s+is\s+up\s+for\s+${UPTIME}\s*$$
  ^\s*Link\s+is\s+down\s+for\s+${DOWNTIME}\s*$$
  ^\s*Advertised\s+link\s+modes\s*:\s*${LINK_MODES}.*$$ -> Continue
  ^\s*Advertised\s+link\s+modes\s*:\s*(?:\d+baseT/(?:Full|Half))\s+${LINK_MODES}\s*$$ -> Continue
  ^\s*Advertised\s+link\s+modes\s*:\s*(?:\d+baseT/(?:Full|Half)).*$$
  ^\s*${LINK_MODES}.*$$ -> Continue
  ^\s*(?:\d+baseT/(?:Full|Half))\s+${LINK_MODES}\s*$$ -> Continue
  ^\s*(?:\d+baseT/(?:Full|Half)).*$$
  ^\s*Link\s+aggregation\s+type\s+is\s+${LINK_AGGREGATION_TYPE}\s*$$
  ^\s*No.\s+of\s+members\s+in\s+this\s+port-channel:\s+${IFACE_IN_PO_NUM}\s+\(active\s+${ACTIVE_IFACE_IN_PO_NUM}(?:,\s+minimum\s+required\s+${MIN_IFACE_REQUIRED_IN_PO_NUM})?\)\s*$$
  ^\s*${IFACE_IN_PO},.*,\s*${IFACE_IN_PO_BANDWIDTH}\s+\(${IFACE_IN_PO_STATUS}\)\s*$$
  ^\s*Active\s+bandwidth\s+is\s+${SUM_BANDWIDTH}\s*$$
  ^\s*Flow\s+control\s+is\s+${FLOW_CONTROL_STATUS},\s+MDIX\s+mode\s+is\s+${MDIX_MODE_STATUS}\s*$$
  ^\s*FEC\s+is\s+${FEC_STATUS}\s*$$
  ^\s*1?5\s+second\s+input\s+rate\s+is\s+${INPUT_RATE}\s+Kbit/s\s*$$
  ^\s*1?5\s+second\s+output\s+rate\s+is\s+${OUTPUT_RATE}\s+Kbit/s\s*$$
  ^\s*${INPUT_PACKETS}\s+packets\s+input,\s+${BYTES_RECEIVED}\s+bytes\s+received\s*$$
  ^\s*${BROADCASTS_NUM}\s+broadcasts,\s+${MULTICASTS_NUM}\s+multicasts\s*$$
  ^\s*${INPUT_ERR_NUM}\s+input\s+errors,\s+${FCS_NUM}\s+FCS,\s+${ALIGNMENTS_NUM}\s+alignment\s*$$
  ^\s*${OVERSIZE_NUM}\s+oversize,\s+${INTERNAL_MAC_NUM}\s+internal\s+MAC\s*$$
  ^\s*${PAUSE_FRAMES_RECEIVED}\s+pause\s+frames\s+received\s*$$
  ^\s*${OUTPUT_PACKETS}\s+packets\s+output,\s+${BYTES_SENT}\s+bytes\s+sent\s*$$
  ^\s*${OUTPUT_ERR_NUM}\s+output\s+errors,\s+${COLLISIONS_NUM}\s+collisions\s*$$
  ^\s*${EXCESSIVE_COLLISIONS_NUM}\s+excessive\s+collisions,\s+${LATE_COLLISIONS_NUM}\s+late\s+collisions\s*$$
  ^\s*${PAUSE_FRAMES_SENT}\s+pause\s+frames\s+transmitted\s*$$
  ^\s*${SYMBOL_ERR_NUM}\s+symbol\s+errors,\s+${CARRIER_NUM}\s+carrier,\s+${SQE_TEST_ERR_NUM}\s+SQE\s+test\s+error\s*$$
  ^\s*Output\s+queues\s*:\s*\(queue\s+#:\s+packets\s+passed/packets\s+dropped\)
  ^\s*${OUTPUT_QUEUE_NUM}\s*:\s*${OUTPUT_QUEUE_PACKETS_PASSED}/${OUTPUT_QUEUE_PACKETS_DROPPED}\s*$$
  ^\s*$$
  ^. -> Error

Columns2
  ^\s*-+(?:\s+-+)*\s*$$
  ^[a-zA-Z\-]+\s*(?:\d+(?:/\d+)*)?.* -> Continue.Record
  ^${INTERFACE}(?:\s+.*?\:\s*${IFACE_IN_PO})?.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){1}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){2}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){3}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){4}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){5}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){6}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){7}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){8}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){9}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){10}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){1}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){2}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){3}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){4}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){5}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){6}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){7}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){8}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){9}${IFACE_IN_PO}.*$$ -> Continue
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,){10}${IFACE_IN_PO}.*$$ -> Continue
  ^[a-zA-Z\-]+\s*(\d+(/\d+)*)?(?:\s+.*?\:\s*(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,)*[a-zA-Z\-]+\s*(\d+(/\d+)*),?)?\s*$$
  ^\s+(?:[a-zA-Z\-]+\s*(\d+(/\d+)*)?,)*[a-zA-Z\-]+\s*(\d+(/\d+)*)\s*$$
  ^\s*$$
  ^. -> Error
`