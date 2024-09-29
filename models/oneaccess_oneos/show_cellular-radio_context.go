package oneaccess_oneos 

type ShowCellularRadioContext struct {
	Context	string	`json:"CONTEXT"`
	Controller	string	`json:"CONTROLLER"`
	Interface	string	`json:"INTERFACE"`
	Cellular_profile	string	`json:"CELLULAR_PROFILE"`
	Bearer_technology	string	`json:"BEARER_TECHNOLOGY"`
	Ipv4_address	string	`json:"IPV4_ADDRESS"`
	Ipv4_subnet	string	`json:"IPV4_SUBNET"`
	Ipv4_gateway	string	`json:"IPV4_GATEWAY"`
	Ipv4_dns1	string	`json:"IPV4_DNS1"`
	Ipv4_dns2	string	`json:"IPV4_DNS2"`
	Data_conn	string	`json:"DATA_CONN"`
	Lastcall_end_reason	string	`json:"LASTCALL_END_REASON"`
	Tx_packets	string	`json:"TX_PACKETS"`
	Rx_packets	string	`json:"RX_PACKETS"`
	Tx_packet_errors	string	`json:"TX_PACKET_ERRORS"`
	Rx_packet_errors	string	`json:"RX_PACKET_ERRORS"`
	Tx_packet_overflow	string	`json:"TX_PACKET_OVERFLOW"`
	Rx_packet_overflow	string	`json:"RX_PACKET_OVERFLOW"`
	Tx_bytes	string	`json:"TX_BYTES"`
	Rx_bytes	string	`json:"RX_BYTES"`
}

var ShowCellularRadioContext_Template = `Value Required CONTEXT (\d+)
Value CONTROLLER (\d+)
Value INTERFACE (Virtual\-Ethernet\s\S+)
Value CELLULAR_PROFILE (\d+)
Value BEARER_TECHNOLOGY (\S+)
Value IPV4_ADDRESS (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value IPV4_SUBNET (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value IPV4_GATEWAY (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value IPV4_DNS1 (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value IPV4_DNS2 (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value DATA_CONN (\S+)
Value LASTCALL_END_REASON (.*)
Value TX_PACKETS (\d+)
Value RX_PACKETS (\d+)
Value TX_PACKET_ERRORS (\d+)
Value RX_PACKET_ERRORS (\d+)
Value TX_PACKET_OVERFLOW (\d+)
Value RX_PACKET_OVERFLOW (\d+)
Value TX_BYTES (\d+)
Value RX_BYTES (\d+)

Start
  ^\s*Context -> Continue.Record
  ^\s*Context\s${CONTEXT} -> CONTEXT
  ^.*$$
  ^. -> Error

CONTEXT
  ^\s+Cellular\scontroller:\s${CONTROLLER}
  ^\s+Interface\s*:\s*${INTERFACE},
  ^\s+Cellular\sprofile\s*:\s*${CELLULAR_PROFILE}
  ^\s+Data\sbearer\stechnology\s*:\s*${BEARER_TECHNOLOGY}
  ^\s+Ipv4\sconnection -> IPV4
  ^\s+Ipv6\sconnection -> IPV6
  ^\s*Internal\s+context\s+Id:
  ^\s*[rt]xSpeed/max\s+:
  ^\s*Call\s+manager\s+state\s+:
  ^\s*Mtu\s+:
  ^\s*Data\s+bearer\s+technology\s+:
  ^\s*=+\s*$$
  ^\s*$$
  ^. -> Error

IPV4
  ^\s+Address\s*:\s*${IPV4_ADDRESS}
  ^\s+Subnet\s*:\s*${IPV4_SUBNET}
  ^\s+Gateway\s*:\s*${IPV4_GATEWAY}
  ^\s+Primary\sDns\s*:\s*${IPV4_DNS1}
  ^\s+Secondary\sDns\s*:\s*${IPV4_DNS2}
  ^\s+Packet\sdata\sconnection\s*:\s*${DATA_CONN}
  ^\s+Last\scall\send\sreason\s*:\s*${LASTCALL_END_REASON}
  ^\s+Tx\spackets\s*:\s*${TX_PACKETS}
  ^\s+Rx\spackets\s*:\s*${RX_PACKETS}
  ^\s+Tx\spacket\serrors\s*:\s*${TX_PACKET_ERRORS}
  ^\s+Rx\spacket\serrors\s*:\s*${RX_PACKET_ERRORS}
  ^\s+Tx\spacket\soverflows\s*:\s*${TX_PACKET_OVERFLOW}
  ^\s+Rx\spacket\soverflows\s*:\s*${RX_PACKET_OVERFLOW}
  ^\s+Tx\sbytes\s*:\s*${TX_BYTES}
  ^\s+Rx\sbytes\s*:\s*${RX_BYTES}
  ^\s*Data\s+statistics\s+:
  ^\s*$$ -> CONTEXT
  ^. -> Error

IPV6
  ^\s*Address\s+\s+:
  ^\s*Gateway\s+\s+:
  ^\s*Primary\s+Dns\s+:
  ^\s*Secondary\s+Dns\s+:
  ^\s*Packet\s+data\s+connection\s+:
  ^\s*Last\s+call\s+end\s+reason\s+:
  ^\s*Data\s+statistics\s+:
  ^\s*$$ -> CONTEXT
  ^. -> Error
`