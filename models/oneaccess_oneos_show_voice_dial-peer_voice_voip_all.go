package models

type OneaccessOneosShowVoiceDialPeerVoiceVoipAll struct {
	Dial_peer	string	`json:"DIAL_PEER"`
	Config_state	string	`json:"CONFIG_STATE"`
	Operational_status	string	`json:"OPERATIONAL_STATUS"`
	Registration_status	string	`json:"REGISTRATION_STATUS"`
	Current_protocol	string	`json:"CURRENT_PROTOCOL"`
	Last_ok	string	`json:"LAST_OK"`
	User_agent	string	`json:"USER_AGENT"`
	Bw_used	string	`json:"BW_USED"`
	Cac	string	`json:"CAC"`
	Bw_unused	string	`json:"BW_UNUSED"`
	Sip_protocol_mode	string	`json:"SIP_PROTOCOL_MODE"`
	Sip_protocol_configured	string	`json:"SIP_PROTOCOL_CONFIGURED"`
	Current_calls	string	`json:"CURRENT_CALLS"`
	Outgoing_calls	string	`json:"OUTGOING_CALLS"`
	Outgoing_bw_used	string	`json:"OUTGOING_BW_USED"`
	Outgoing_cac	string	`json:"OUTGOING_CAC"`
	Outgoing_bw_unused	string	`json:"OUTGOING_BW_UNUSED"`
	Outgoing_failures	string	`json:"OUTGOING_FAILURES"`
	Outgoing_failures_q931	string	`json:"OUTGOING_FAILURES_Q931"`
	Outgoing_failures_q931_0_normal	string	`json:"OUTGOING_FAILURES_Q931_0_NORMAL"`
	Outgoing_failures_q931_1_normal	string	`json:"OUTGOING_FAILURES_Q931_1_NORMAL"`
	Outgoing_failures_q931_16_normal	string	`json:"OUTGOING_FAILURES_Q931_16_NORMAL"`
	Outgoing_failures_q931_17_busy	string	`json:"OUTGOING_FAILURES_Q931_17_BUSY"`
	Outgoing_failures_q931_18_no_answer	string	`json:"OUTGOING_FAILURES_Q931_18_NO_ANSWER"`
	Outgoing_failures_q931_2_unavailable_resources	string	`json:"OUTGOING_FAILURES_Q931_2_UNAVAILABLE_RESOURCES"`
	Outgoing_failures_q931_3_unavailable_service	string	`json:"OUTGOING_FAILURES_Q931_3_UNAVAILABLE_SERVICE"`
	Outgoing_failures_q931_4_service_not_provided	string	`json:"OUTGOING_FAILURES_Q931_4_SERVICE_NOT_PROVIDED"`
	Outgoing_failures_q931_5_invalid_message	string	`json:"OUTGOING_FAILURES_Q931_5_INVALID_MESSAGE"`
	Outgoing_failures_q931_6_protocol	string	`json:"OUTGOING_FAILURES_Q931_6_PROTOCOL"`
	Outgoing_failures_q931_7_interworking	string	`json:"OUTGOING_FAILURES_Q931_7_INTERWORKING"`
	Outgoing_failures_sip_call	string	`json:"OUTGOING_FAILURES_SIP_CALL"`
	Outgoing_failures_capabilities	string	`json:"OUTGOING_FAILURES_CAPABILITIES"`
	Outgoing_failures_protocol	string	`json:"OUTGOING_FAILURES_PROTOCOL"`
	Outgoing_failures_internal_call	string	`json:"OUTGOING_FAILURES_INTERNAL_CALL"`
	Outgoing_failures_dsp_unavailable	string	`json:"OUTGOING_FAILURES_DSP_UNAVAILABLE"`
	Outgoing_failures_max_bw_exceeded	string	`json:"OUTGOING_FAILURES_MAX_BW_EXCEEDED"`
	Outgoing_failures_max_conn_exceeded	string	`json:"OUTGOING_FAILURES_MAX_CONN_EXCEEDED"`
	Outgoing_failures_rtp_payload	string	`json:"OUTGOING_FAILURES_RTP_PAYLOAD"`
	Outgoing_failures_not_specified	string	`json:"OUTGOING_FAILURES_NOT_SPECIFIED"`
	Incoming_calls	string	`json:"INCOMING_CALLS"`
	Incoming_bw_used	string	`json:"INCOMING_BW_USED"`
	Incoming_cac	string	`json:"INCOMING_CAC"`
	Incoming_bw_unused	string	`json:"INCOMING_BW_UNUSED"`
	Incoming_failures	string	`json:"INCOMING_FAILURES"`
	Incoming_failures_local_port	string	`json:"INCOMING_FAILURES_LOCAL_PORT"`
	Incoming_failures_sip_call	string	`json:"INCOMING_FAILURES_SIP_CALL"`
	Incoming_failures_capabilities	string	`json:"INCOMING_FAILURES_CAPABILITIES"`
	Incoming_failures_protocol	string	`json:"INCOMING_FAILURES_PROTOCOL"`
	Incoming_failures_internal_call	string	`json:"INCOMING_FAILURES_INTERNAL_CALL"`
	Incoming_failures_dsp_unavailable	string	`json:"INCOMING_FAILURES_DSP_UNAVAILABLE"`
	Incoming_failures_unknown_number	string	`json:"INCOMING_FAILURES_UNKNOWN_NUMBER"`
	Incoming_failures_channel_or_port_unavailable	string	`json:"INCOMING_FAILURES_CHANNEL_OR_PORT_UNAVAILABLE"`
	Incoming_failures_max_bw_exceeded	string	`json:"INCOMING_FAILURES_MAX_BW_EXCEEDED"`
	Incoming_failures_max_conn_exceeded	string	`json:"INCOMING_FAILURES_MAX_CONN_EXCEEDED"`
	Incoming_failures_rtp_payload	string	`json:"INCOMING_FAILURES_RTP_PAYLOAD"`
	Incoming_failures_not_specified	string	`json:"INCOMING_FAILURES_NOT_SPECIFIED"`
	Voice_fax_real_dsp_switching	string	`json:"VOICE_FAX_REAL_DSP_SWITCHING"`
	Rtp_packets_transmitted	string	`json:"RTP_PACKETS_TRANSMITTED"`
	Rtp_packets_received	string	`json:"RTP_PACKETS_RECEIVED"`
	Rtp_bytes_transmitted	string	`json:"RTP_BYTES_TRANSMITTED"`
	Rtp_bytes_received	string	`json:"RTP_BYTES_RECEIVED"`
	Rtp_excessive_jitter_events	string	`json:"RTP_EXCESSIVE_JITTER_EVENTS"`
	Rtp_packets_lost	string	`json:"RTP_PACKETS_LOST"`
	Rtp_packets_invalid	string	`json:"RTP_PACKETS_INVALID"`
	Frame_error_rate_total	string	`json:"FRAME_ERROR_RATE_TOTAL"`
	Frame_error_rate_lt_0_01	string	`json:"FRAME_ERROR_RATE_LT_0_01"`
	Frame_error_rate_lt_0_1	string	`json:"FRAME_ERROR_RATE_LT_0_1"`
	Frame_error_rate_lt_0_5	string	`json:"FRAME_ERROR_RATE_LT_0_5"`
	Frame_error_rate_lt_1	string	`json:"FRAME_ERROR_RATE_LT_1"`
	Frame_error_rate_lt_5	string	`json:"FRAME_ERROR_RATE_LT_5"`
	Frame_error_rate_ge_5	string	`json:"FRAME_ERROR_RATE_GE_5"`
	Modem_switches	string	`json:"MODEM_SWITCHES"`
	Fax_outgoing	string	`json:"FAX_OUTGOING"`
	Fax_incoming	string	`json:"FAX_INCOMING"`
	Fax_failures	string	`json:"FAX_FAILURES"`
	Fax_failures_request_mode	string	`json:"FAX_FAILURES_REQUEST_MODE"`
	Fax_failures_pre_msg	string	`json:"FAX_FAILURES_PRE_MSG"`
	Fax_failures_page	string	`json:"FAX_FAILURES_PAGE"`
	Fax_packets_transmitted	string	`json:"FAX_PACKETS_TRANSMITTED"`
	Fax_packets_received	string	`json:"FAX_PACKETS_RECEIVED"`
	Fax_bytes_transmitted	string	`json:"FAX_BYTES_TRANSMITTED"`
	Fax_bytes_received	string	`json:"FAX_BYTES_RECEIVED"`
	Fax_packets_lost	string	`json:"FAX_PACKETS_LOST"`
}

var OneaccessOneosShowVoiceDialPeerVoiceVoipAll_Template = `# global vars
Value DIAL_PEER (\d+)
Value CONFIG_STATE (\S+)
Value OPERATIONAL_STATUS (\S+)
Value REGISTRATION_STATUS (\S+)
Value CURRENT_PROTOCOL (\S+)
Value LAST_OK (\d+)
Value USER_AGENT (\w+)
Value BW_USED (\d+)
Value CAC (\d+)
Value BW_UNUSED (\d+)
Value SIP_PROTOCOL_MODE (\S+)
Value SIP_PROTOCOL_CONFIGURED (\w+)
Value CURRENT_CALLS (\d+)
# outgoing call vars
Value OUTGOING_CALLS (\d+)
Value OUTGOING_BW_USED (\d+)
Value OUTGOING_CAC (\d+)
Value OUTGOING_BW_UNUSED (\d+)
Value OUTGOING_FAILURES (\d+)
Value OUTGOING_FAILURES_Q931 (\d+)
Value OUTGOING_FAILURES_Q931_0_NORMAL (\d+)
Value OUTGOING_FAILURES_Q931_1_NORMAL (\d+)
Value OUTGOING_FAILURES_Q931_16_NORMAL (\d+)
Value OUTGOING_FAILURES_Q931_17_BUSY (\d+)
Value OUTGOING_FAILURES_Q931_18_NO_ANSWER (\d+)
Value OUTGOING_FAILURES_Q931_2_UNAVAILABLE_RESOURCES (\d+)
Value OUTGOING_FAILURES_Q931_3_UNAVAILABLE_SERVICE (\d+)
Value OUTGOING_FAILURES_Q931_4_SERVICE_NOT_PROVIDED (\d+)
Value OUTGOING_FAILURES_Q931_5_INVALID_MESSAGE (\d+)
Value OUTGOING_FAILURES_Q931_6_PROTOCOL (\d+)
Value OUTGOING_FAILURES_Q931_7_INTERWORKING (\d+)
Value OUTGOING_FAILURES_SIP_CALL (\d+)
Value OUTGOING_FAILURES_CAPABILITIES (\d+)
Value OUTGOING_FAILURES_PROTOCOL (\d+)
Value OUTGOING_FAILURES_INTERNAL_CALL (\d+)
Value OUTGOING_FAILURES_DSP_UNAVAILABLE (\d+)
Value OUTGOING_FAILURES_MAX_BW_EXCEEDED (\d+)
Value OUTGOING_FAILURES_MAX_CONN_EXCEEDED (\d+)
Value OUTGOING_FAILURES_RTP_PAYLOAD (\d+)
Value OUTGOING_FAILURES_NOT_SPECIFIED (\d+)
# incoming call vars
Value INCOMING_CALLS (\d+)
Value INCOMING_BW_USED (\d+)
Value INCOMING_CAC (\d+)
Value INCOMING_BW_UNUSED (\d+)
Value INCOMING_FAILURES (\d+)
Value INCOMING_FAILURES_LOCAL_PORT (\d+)
Value INCOMING_FAILURES_SIP_CALL (\d+)
Value INCOMING_FAILURES_CAPABILITIES (\d+)
Value INCOMING_FAILURES_PROTOCOL (\d+)
Value INCOMING_FAILURES_INTERNAL_CALL (\d+)
Value INCOMING_FAILURES_DSP_UNAVAILABLE (\d+)
Value INCOMING_FAILURES_UNKNOWN_NUMBER (\d+)
Value INCOMING_FAILURES_CHANNEL_OR_PORT_UNAVAILABLE (\d+)
Value INCOMING_FAILURES_MAX_BW_EXCEEDED (\d+)
Value INCOMING_FAILURES_MAX_CONN_EXCEEDED (\d+)
Value INCOMING_FAILURES_RTP_PAYLOAD (\d+)
Value INCOMING_FAILURES_NOT_SPECIFIED (\d+)
# VOICE & FAX STATS
Value VOICE_FAX_REAL_DSP_SWITCHING (\d+)
# RTP vars
Value RTP_PACKETS_TRANSMITTED (\d+)
Value RTP_PACKETS_RECEIVED (\d+)
Value RTP_BYTES_TRANSMITTED (\d+)
Value RTP_BYTES_RECEIVED (\d+)
Value RTP_EXCESSIVE_JITTER_EVENTS (\d+)
Value RTP_PACKETS_LOST (\d+)
Value RTP_PACKETS_INVALID (\d+)
# frame error call vars
Value FRAME_ERROR_RATE_TOTAL (\d+)
Value FRAME_ERROR_RATE_LT_0_01 (\d+)
Value FRAME_ERROR_RATE_LT_0_1 (\d+)
Value FRAME_ERROR_RATE_LT_0_5 (\d+)
Value FRAME_ERROR_RATE_LT_1 (\d+)
Value FRAME_ERROR_RATE_LT_5 (\d+)
Value FRAME_ERROR_RATE_GE_5 (\d+)
# modem passthrough vars
Value MODEM_SWITCHES (\d+)
# FAX vars
Value FAX_OUTGOING (\d+)
Value FAX_INCOMING (\d+)
Value FAX_FAILURES (\d+)
Value FAX_FAILURES_REQUEST_MODE (\d+)
Value FAX_FAILURES_PRE_MSG (\d+)
Value FAX_FAILURES_PAGE (\d+)
Value FAX_PACKETS_TRANSMITTED (\d+)
Value FAX_PACKETS_RECEIVED (\d+)
Value FAX_BYTES_TRANSMITTED (\d+)
Value FAX_BYTES_RECEIVED (\d+)
Value FAX_PACKETS_LOST (\d+)


Start
  ^\s*Dial\sPeer\s+${DIAL_PEER}
  ^\s*Config\sstate\s+${CONFIG_STATE}
  ^\s*Operstatus\s+${OPERATIONAL_STATUS}
  ^\s*Registration\sstatus\s+${REGISTRATION_STATUS}
  ^\s*Current\sprotocol\s+${CURRENT_PROTOCOL}
  ^\s*lastOK\s+${LAST_OK}
  ^\s*User\sAgent\s+\(?${USER_AGENT}\)?
  ^\s*Bandwidth\sreally\sused\/CAC\svalue\/unused\s+${BW_USED}\s*\/\s*${CAC}\s*\/\s*${BW_UNUSED}
  ^\s*Current\ssip-protocol-mode\s+${SIP_PROTOCOL_MODE}.*:\s${SIP_PROTOCOL_CONFIGURED}
  ^\s*Current\sCalls\s+${CURRENT_CALLS}
  ^\s*Outgoing\sCalls\s*$$ -> OutgoingCalls
  ^\s*Incoming\sCalls\s*$$ -> IncomingCalls
  ^\s*Voice\s\&\sFAX\sstatistics\s*$$ -> VoiceFax
  ^\s*RTP\sstatistics\s*$$ -> RTP
  ^\s*Number\sof\scalls\swith\sframe\serror\srate\s*$$ -> CallsWithFrameError
  ^\s*Modem\spassthrough\s*$$ -> MODEM
  ^\s*T38\sFAX\sCalls\s*$$ -> FAX
  ^\s*$$
  ^. -> Error

OutgoingCalls
  ^\s*Outgoing\sCalls\s+${OUTGOING_CALLS}
  ^\s*Bandwidth\sreally\sused\/CAC\svalue\/unused\s+${OUTGOING_BW_USED}\s*\/\s*${OUTGOING_CAC}\s*\/\s*${OUTGOING_BW_UNUSED}
  ^\s*Outgoing\scalls\sfailures\s+${OUTGOING_FAILURES}
  ^\s*Q931\sCall\sfailures\s+${OUTGOING_FAILURES_Q931}
  ^\s*Cause\sClass\s0\s+.*${OUTGOING_FAILURES_Q931_0_NORMAL}
  ^\s*Cause\sClass\s1\s+.*${OUTGOING_FAILURES_Q931_1_NORMAL}
  ^\s*Normal\sCause\s\(16\)\s+${OUTGOING_FAILURES_Q931_16_NORMAL}
  ^\s*User\sbusy\s\(17\)\s+${OUTGOING_FAILURES_Q931_17_BUSY}
  ^\s*No\sanswer\s\(18\)\s+${OUTGOING_FAILURES_Q931_18_NO_ANSWER}
  ^\s*Cause\sClass\s2\s+.*${OUTGOING_FAILURES_Q931_2_UNAVAILABLE_RESOURCES}
  ^\s*Cause\sClass\s3\s+.*${OUTGOING_FAILURES_Q931_3_UNAVAILABLE_SERVICE}
  ^\s*Cause\sClass\s4\s+.*${OUTGOING_FAILURES_Q931_4_SERVICE_NOT_PROVIDED}
  ^\s*Cause\sClass\s5\s+.*${OUTGOING_FAILURES_Q931_5_INVALID_MESSAGE}
  ^\s*Cause\sClass\s6\s+.*${OUTGOING_FAILURES_Q931_6_PROTOCOL}
  ^\s*Cause\sClass\s7\s+.*${OUTGOING_FAILURES_Q931_7_INTERWORKING}
  ^\s*SIP\sCall\sfailures\s+${OUTGOING_FAILURES_SIP_CALL}
  ^\s*Incompatible\scapabilities\s+${OUTGOING_FAILURES_CAPABILITIES}
  ^\s*Protocol\serrors\s+${OUTGOING_FAILURES_PROTOCOL}
  ^\s*Internal\scall\sfailures\s+${OUTGOING_FAILURES_INTERNAL_CALL}
  ^\s*DSP\sunavailable\s+${OUTGOING_FAILURES_DSP_UNAVAILABLE}
  ^\s*Max-bandwidth\sexceeded\s+${OUTGOING_FAILURES_MAX_BW_EXCEEDED}
  ^\s*Max-connection\sexceeded\s+${OUTGOING_FAILURES_MAX_CONN_EXCEEDED}
  ^\s*RTP\sdynamic-payload\serror\s+${OUTGOING_FAILURES_RTP_PAYLOAD}
  ^\s*Not\sspecified\s+${OUTGOING_FAILURES_NOT_SPECIFIED}
  ^\s*Incoming\sCalls\s*$$ -> IncomingCalls
  ^\s*$$
  ^. -> Error

IncomingCalls
  ^\s*Incoming\scalls\s+${INCOMING_CALLS}
  ^\s*Bandwidth\sreally\sused\/CAC\svalue\/unused\s+${INCOMING_BW_USED}\s*\/\s*${INCOMING_CAC}\s*\/\s*${INCOMING_BW_UNUSED}
  ^\s*Incoming\scalls\sfailures\s+${INCOMING_FAILURES}
  ^\s*Local\sPort\sCall\sfailures\s+${INCOMING_FAILURES_LOCAL_PORT}
  ^\s*SIP\sCall\sfailures\s+${INCOMING_FAILURES_SIP_CALL}
  ^\s*Incompatible\scapabilities\s+${INCOMING_FAILURES_CAPABILITIES}
  ^\s*Protocol\serrors\s+${INCOMING_FAILURES_PROTOCOL}
  ^\s*Internal\scall\sfailures\s+${INCOMING_FAILURES_INTERNAL_CALL}
  ^\s*DSP\sunavailable\s+${INCOMING_FAILURES_DSP_UNAVAILABLE}
  ^\s*Unknown\snumber\s+${INCOMING_FAILURES_UNKNOWN_NUMBER}
  ^\s*Channel\s\/\sport\sunavailable\s+${INCOMING_FAILURES_CHANNEL_OR_PORT_UNAVAILABLE}
  ^\s*Max-bandwidth\sexceeded\s+${INCOMING_FAILURES_MAX_BW_EXCEEDED}
  ^\s*Max-connection\sexceeded\s+${INCOMING_FAILURES_MAX_CONN_EXCEEDED}
  ^\s*RTP\sdynamic-payload\serror\s+${INCOMING_FAILURES_RTP_PAYLOAD}
  ^\s*Not\sspecified\s+${INCOMING_FAILURES_NOT_SPECIFIED}
  ^\s*Voice\s.\sFax\sstatistics\s*$$ -> VoiceFax
  ^\s*RTP\sstatistics\s*$$ -> RTP
  ^\s*$$
  ^. -> Error

VoiceFax
  ^\s*Number\sof\sreal\sdsp\sswitching\s+${VOICE_FAX_REAL_DSP_SWITCHING}
  ^\s*RTP\sstatistics\s*$$ -> RTP
  ^\s*$$
  ^. -> Error

RTP
  ^\s*Number\sof\stransmitted\spackets\s+${RTP_PACKETS_TRANSMITTED}
  ^\s*Number\sof\sreceived\spackets\s+${RTP_PACKETS_RECEIVED}
  ^\s*Number\sof\stransmitted\sbytes\s+${RTP_BYTES_TRANSMITTED}
  ^\s*Number\sof\sreceived\sbytes\s+${RTP_BYTES_RECEIVED}
  ^\s*Number\sof\sexcessive\sjitter\sevents\s+${RTP_EXCESSIVE_JITTER_EVENTS}
  ^\s*Number\sof\slost\spackets\s+${RTP_PACKETS_LOST}
  ^\s*Number\sof\sinvalid\spackets\s+${RTP_PACKETS_INVALID}
  ^\s*Number\sof\scalls\swith\sframe\serror\srate\s*$$ -> CallsWithFrameError
  ^\s*$$
  ^. -> Error

CallsWithFrameError
  ^\s*total
  ^\s*${FRAME_ERROR_RATE_TOTAL}\s*${FRAME_ERROR_RATE_LT_0_01}\s*${FRAME_ERROR_RATE_LT_0_1}\s*${FRAME_ERROR_RATE_LT_0_5}\s*${FRAME_ERROR_RATE_LT_1}\s*${FRAME_ERROR_RATE_LT_5}\s*${FRAME_ERROR_RATE_GE_5}\s*
  ^\s*Modem\spassthrough\s*$$ -> MODEM
  ^\s*$$
  ^. -> Error

MODEM
  ^\s*Number\sof\sswitching\sto\smodem\smode\s+${MODEM_SWITCHES}
  ^\s*T38\sFAX\sCalls\s*$$ -> FAX
  ^\s*$$
  ^. -> Error

FAX
  ^\s*Number\sof\soutgoing\sfax\s+${FAX_OUTGOING}
  ^\s*Number\sof\sincoming\sfax\s+${FAX_INCOMING}
  ^\s*Number\sof\sfailures\s+${FAX_FAILURES}
  ^\s*Request\sMode\sfailure\s+${FAX_FAILURES_REQUEST_MODE}
  ^\s*Pre-message\sprocedure\sfailure\s+${FAX_FAILURES_PRE_MSG}
  ^\s*Page\sfailure\s+${FAX_FAILURES_PAGE}
  ^\s*Number\sof\stransmitted\spackets\s+${FAX_PACKETS_TRANSMITTED}
  ^\s*Number\sof\sreceived\spackets\s+${FAX_PACKETS_RECEIVED}
  ^\s*Number\sof\stransmitted\sbytes\s+${FAX_BYTES_TRANSMITTED}
  ^\s*Number\sof\sreceived\sbytes\s+${FAX_BYTES_RECEIVED}
  ^\s*Number\sof\slost\spackets\s+${FAX_PACKETS_LOST}
  ^\s*$$
  ^. -> Error

`