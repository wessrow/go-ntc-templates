package models

type OneaccessOneosShowVoiceSipGateway struct {
	Gateway_state	string	`json:"GATEWAY_STATE"`
	Gateway_operational_status	string	`json:"GATEWAY_OPERATIONAL_STATUS"`
	Sockets	[]string	`json:"SOCKETS"`
	Registration_state	string	`json:"REGISTRATION_STATE"`
	Rtp_monitoring	string	`json:"RTP_MONITORING"`
	Registrar	string	`json:"REGISTRAR"`
	Registrar_neighbor_registered	string	`json:"REGISTRAR_NEIGHBOR_REGISTERED"`
	Registrar_neighbor_available	string	`json:"REGISTRAR_NEIGHBOR_AVAILABLE"`
	Bw_used	string	`json:"BW_USED"`
	Cac	string	`json:"CAC"`
	Bw_unused	string	`json:"BW_UNUSED"`
	Thold_bw_to_switch	string	`json:"THOLD_BW_TO_SWITCH"`
	Max_bw_exceeded	string	`json:"MAX_BW_EXCEEDED"`
	Neighbor_lower_switching	string	`json:"NEIGHBOR_LOWER_SWITCHING"`
	Registration_errors	string	`json:"REGISTRATION_ERRORS"`
	Sip_protocol_mode	string	`json:"SIP_PROTOCOL_MODE"`
	Sip_protocol_configured	string	`json:"SIP_PROTOCOL_CONFIGURED"`
	Current_call	string	`json:"CURRENT_CALL"`
	Calls_released_rtp_monitoring	string	`json:"CALLS_RELEASED_RTP_MONITORING"`
	Authentication_rejects	string	`json:"AUTHENTICATION_REJECTS"`
	Dropped_packet	string	`json:"DROPPED_PACKET"`
	Dropped_packet_rate_limit	string	`json:"DROPPED_PACKET_RATE_LIMIT"`
	Dropped_packet_memory_limit	string	`json:"DROPPED_PACKET_MEMORY_LIMIT"`
	Dropped_packet_cpu_limit	string	`json:"DROPPED_PACKET_CPU_LIMIT"`
	Dropped_packet_acl	string	`json:"DROPPED_PACKET_ACL"`
	Dropped_packet_unknown_proxy	string	`json:"DROPPED_PACKET_UNKNOWN_PROXY"`
}

var OneaccessOneosShowVoiceSipGateway_Template = `Value GATEWAY_STATE (\S+)
Value GATEWAY_OPERATIONAL_STATUS (\S+)
Value List SOCKETS (\S+:\d+)
Value REGISTRATION_STATE (\S+)
Value RTP_MONITORING (\S+)
Value REGISTRAR (\S+:\d+)
Value REGISTRAR_NEIGHBOR_REGISTERED (\d+)
Value REGISTRAR_NEIGHBOR_AVAILABLE (\d+)
Value BW_USED (\d+)
Value CAC (\d+)
Value BW_UNUSED (\d+)
Value THOLD_BW_TO_SWITCH (\S+)
Value MAX_BW_EXCEEDED (\d+)
Value NEIGHBOR_LOWER_SWITCHING (\d+)
Value REGISTRATION_ERRORS (\d+)
Value SIP_PROTOCOL_MODE (\S+)
Value SIP_PROTOCOL_CONFIGURED (\S+)
Value CURRENT_CALL (\d+)
Value CALLS_RELEASED_RTP_MONITORING (\d+)
Value AUTHENTICATION_REJECTS (\d+)
Value DROPPED_PACKET (\d+)
Value DROPPED_PACKET_RATE_LIMIT (\d+)
Value DROPPED_PACKET_MEMORY_LIMIT (\d+)
Value DROPPED_PACKET_CPU_LIMIT (\d+)
Value DROPPED_PACKET_ACL (\d+)
Value DROPPED_PACKET_UNKNOWN_PROXY (\d+)

Start
  ^\s*Sip-Gateway\sstatistics
  ^\s*Gateway\sstate\s+${GATEWAY_STATE}
  ^\s*Operstatus\s+${GATEWAY_OPERATIONAL_STATUS}
  ^\s*Sockidx:\s*\d+,\s*${SOCKETS}
  ^\s*Registration\sstate\s+${REGISTRATION_STATE}
  ^\s*RTP\smonitoring\s+${RTP_MONITORING}
  ^\s*Nb\sRegistered\sendpoints
  ^\s*\[${REGISTRAR_NEIGHBOR_REGISTERED}\/${REGISTRAR_NEIGHBOR_AVAILABLE}\]\s+${REGISTRAR}
  ^\s*Bandwidth\sreally\sused\/CAC\svalue\/unused\s+${BW_USED}\s*\/\s*${CAC}\s*\/\s*${BW_UNUSED}
  ^\s*Threshold\sof\sbandwidth\sto\sswitch\s+${THOLD_BW_TO_SWITCH}
  ^\s*Max\sBandwidth\sexceeded\s+${MAX_BW_EXCEEDED}
  ^\s*Number\sof\slower\sswitching\s+${NEIGHBOR_LOWER_SWITCHING}
  ^\s*Registration\serrors\s+${REGISTRATION_ERRORS}
  ^\s*Current\ssip-protocol-mode\s+${SIP_PROTOCOL_MODE}\s*\(config\s*:\s*${SIP_PROTOCOL_CONFIGURED}\)
  ^\s*Current\scall\s+${CURRENT_CALL}
  ^\s*Calls\sreleased\sby\srtp\smonitoring\s+${CALLS_RELEASED_RTP_MONITORING}
  ^\s*Authentication\sRejects\s+${AUTHENTICATION_REJECTS}
  ^\s*Dropped\spackets\s+${DROPPED_PACKET}
  ^\s*due\sto\srate\slimitation\s+${DROPPED_PACKET_RATE_LIMIT}
  ^\s*due\sto\smemory\slimitation\s+${DROPPED_PACKET_MEMORY_LIMIT}
  ^\s*due\sto\sCPU\slimitation\s+${DROPPED_PACKET_CPU_LIMIT}
  ^\s*due\sto\sdenied\sby\sacl\s+${DROPPED_PACKET_ACL}
  ^\s*due\sto\sunknown\sproxy\s+${DROPPED_PACKET_UNKNOWN_PROXY}
  ^\s*SIP-GW\s+entity\s+opened\s+sockets:
  ^\s+\S+P\s+sockets:
  ^\s*$$
  ^. -> Error

`