package oneaccess_oneos 

type ShowVoiceVoicePortPriAll struct {
	Port	string	`json:"PORT"`
	Physical_type	string	`json:"PHYSICAL_TYPE"`
	Protocol_descriptor	string	`json:"PROTOCOL_DESCRIPTOR"`
	Config_state	string	`json:"CONFIG_STATE"`
	Loop_state	string	`json:"LOOP_STATE"`
	Framing	string	`json:"FRAMING"`
	L1_status	string	`json:"L1_STATUS"`
	Ais	string	`json:"AIS"`
	Los	string	`json:"LOS"`
	Rai	string	`json:"RAI"`
	Pri_ais_occurrences	string	`json:"PRI_AIS_OCCURRENCES"`
	Pri_los_occurrences	string	`json:"PRI_LOS_OCCURRENCES"`
	Pri_rai_occurrences	string	`json:"PRI_RAI_OCCURRENCES"`
	L2_status	string	`json:"L2_STATUS"`
	Pri_tx_frames_on_d_channel	string	`json:"PRI_TX_FRAMES_ON_D_CHANNEL"`
	Pri_rx_frames_on_d_channel	string	`json:"PRI_RX_FRAMES_ON_D_CHANNEL"`
	Attached_voip_dial_peer	string	`json:"ATTACHED_VOIP_DIAL_PEER"`
	Nbr_voice_communication	string	`json:"NBR_VOICE_COMMunication"`
	Outgoing_calls	string	`json:"OUTGOING_CALLS"`
	Outgoing_failures	string	`json:"OUTGOING_FAILURES"`
	Outgoing_failures_physical_intf_down	string	`json:"OUTGOING_FAILURES_PHYSICAL_INTF_DOWN"`
	Outgoing_failures_0_normal	string	`json:"OUTGOING_FAILURES_0_NORMAL"`
	Outgoing_failures_1_normal	string	`json:"OUTGOING_FAILURES_1_NORMAL"`
	Outgoing_failures_16_normal	string	`json:"OUTGOING_FAILURES_16_NORMAL"`
	Outgoing_failures_17_busy	string	`json:"OUTGOING_FAILURES_17_BUSY"`
	Outgoing_failures_18_no_answer	string	`json:"OUTGOING_FAILURES_18_NO_ANSWER"`
	Outgoing_failures_2_unavailable_resources	string	`json:"OUTGOING_FAILURES_2_UNAVAILABLE_RESOURCES"`
	Outgoing_failures_3_unavailable_service	string	`json:"OUTGOING_FAILURES_3_UNAVAILABLE_SERVICE"`
	Outgoing_failures_4_service_not_provided	string	`json:"OUTGOING_FAILURES_4_SERVICE_NOT_PROVIDED"`
	Outgoing_failures_5_invalid_message	string	`json:"OUTGOING_FAILURES_5_INVALID_MESSAGE"`
	Outgoing_failures_6_protocol	string	`json:"OUTGOING_FAILURES_6_PROTOCOL"`
	Outgoing_failures_7_interworking	string	`json:"OUTGOING_FAILURES_7_INTERWORKING"`
	Incoming_calls	string	`json:"INCOMING_CALLS"`
	Incoming_calls_backup_invoked	string	`json:"INCOMING_CALLS_BACKUP_INVOKED"`
	Incoming_failures	string	`json:"INCOMING_FAILURES"`
	Incoming_failures_remote	string	`json:"INCOMING_FAILURES_REMOTE"`
	Incoming_failures_unkown_number	string	`json:"INCOMING_FAILURES_UNKOWN_NUMBER"`
	Incoming_failures_dsp_unavailable	string	`json:"INCOMING_FAILURES_DSP_UNAVAILABLE"`
	Incoming_failures_no_voip_resource_available	string	`json:"INCOMING_FAILURES_NO_VOIP_RESOURCE_AVAILABLE"`
	Incoming_failures_not_specified	string	`json:"INCOMING_FAILURES_NOT_SPECIFIED"`
}

var ShowVoiceVoicePortPriAll_Template = `Value Required PORT (\d+\/\d+)
Value PHYSICAL_TYPE (\S+)
Value PROTOCOL_DESCRIPTOR (\S+)
Value CONFIG_STATE (\S+)
Value LOOP_STATE (\S+)
Value FRAMING (\S+)
Value L1_STATUS (\S+)
Value AIS (\S+)
Value LOS (\S+)
Value RAI (\S+)
Value PRI_AIS_OCCURRENCES (\d+)
Value PRI_LOS_OCCURRENCES (\d+)
Value PRI_RAI_OCCURRENCES (\d+)
Value L2_STATUS (\S+)
Value PRI_TX_FRAMES_ON_D_CHANNEL (\d+)
Value PRI_RX_FRAMES_ON_D_CHANNEL (\d+)
Value ATTACHED_VOIP_DIAL_PEER (\d+)
Value NBR_VOICE_COMMunication (\d+)
Value OUTGOING_CALLS (\d+)
Value OUTGOING_FAILURES (\d+)
Value OUTGOING_FAILURES_PHYSICAL_INTF_DOWN (\d+)
Value OUTGOING_FAILURES_0_NORMAL (\d+)
Value OUTGOING_FAILURES_1_NORMAL (\d+)
Value OUTGOING_FAILURES_16_NORMAL (\d+)
Value OUTGOING_FAILURES_17_BUSY (\d+)
Value OUTGOING_FAILURES_18_NO_ANSWER (\d+)
Value OUTGOING_FAILURES_2_UNAVAILABLE_RESOURCES (\d+)
Value OUTGOING_FAILURES_3_UNAVAILABLE_SERVICE (\d+)
Value OUTGOING_FAILURES_4_SERVICE_NOT_PROVIDED (\d+)
Value OUTGOING_FAILURES_5_INVALID_MESSAGE (\d+)
Value OUTGOING_FAILURES_6_PROTOCOL (\d+)
Value OUTGOING_FAILURES_7_INTERWORKING (\d+)
Value INCOMING_CALLS (\d+)
Value INCOMING_CALLS_BACKUP_INVOKED (\d+)
Value INCOMING_FAILURES (\d+)
Value INCOMING_FAILURES_REMOTE (\d+)
Value INCOMING_FAILURES_UNKOWN_NUMBER (\d+)
Value INCOMING_FAILURES_DSP_UNAVAILABLE (\d+)
Value INCOMING_FAILURES_NO_VOIP_RESOURCE_AVAILABLE (\d+)
Value INCOMING_FAILURES_NOT_SPECIFIED (\d+)

Start
  ^\s*voice\sport -> Continue.Record
  ^\s*voice\sport\s+${PORT}
  ^\s*physical\stype\s+${PHYSICAL_TYPE}
  ^\s*protocol\sdescriptor\s+${PROTOCOL_DESCRIPTOR}
  ^\s*config\sstate\s+${CONFIG_STATE}
  ^\s*loop\sstate\s+${LOOP_STATE}
  ^\s*framing\s+${FRAMING}
  ^\s*layer\s1\sstatus\s+${L1_STATUS}
  ^\s*Alarm\sIndication\sSignal\s+\(AIS\)\s+${AIS}
  ^\s*Loss\sOff\sSignal\s+\(LOS\)\s+${LOS}
  ^\s*Remote\sAlarm\sIndication\s+\(RAI\)\s+${RAI}
  ^\s*pri\sAIS\soccurrence\(s\)\s+${PRI_AIS_OCCURRENCES}
  ^\s*pri\sLOS\soccurrence\(s\)\s+${PRI_LOS_OCCURRENCES}
  ^\s*pri\sRAI\soccurrence\(s\)\s+${PRI_RAI_OCCURRENCES}
  ^\s*layer\s2\sstatus\s+${L2_STATUS}
  ^\s*pri\sTx\sframes\son\sD\schannel\s+${PRI_TX_FRAMES_ON_D_CHANNEL}
  ^\s*pri\sRx\sframes\son\sD\schannel\s+${PRI_RX_FRAMES_ON_D_CHANNEL}
  ^\s*attached\svoip\sdial\speer\s+${ATTACHED_VOIP_DIAL_PEER}
  ^\s*number\sof\svoice\scommunication\s+${NBR_VOICE_COMMunication}
  ^\s*Outgoing\scalls\s+${OUTGOING_CALLS}
  ^\s*Outgoing\scalls\sfailures\s+${OUTGOING_FAILURES}
  ^\s*Physical\sInterface\sdown\s+${OUTGOING_FAILURES_PHYSICAL_INTF_DOWN}
  ^\s*Cause\sClass\s0\s+.*${OUTGOING_FAILURES_0_NORMAL}
  ^\s*Cause\sClass\s1\s+.*${OUTGOING_FAILURES_1_NORMAL}
  ^\s*Normal\sCause\s\(16\)\s+${OUTGOING_FAILURES_16_NORMAL}
  ^\s*User\sbusy\s\(17\)\s+${OUTGOING_FAILURES_17_BUSY}
  ^\s*No\sanswer\s\(18\)\s+${OUTGOING_FAILURES_18_NO_ANSWER}
  ^\s*Cause\sClass\s2\s+.*${OUTGOING_FAILURES_2_UNAVAILABLE_RESOURCES}
  ^\s*Cause\sClass\s3\s+.*${OUTGOING_FAILURES_3_UNAVAILABLE_SERVICE}
  ^\s*Cause\sClass\s4\s+.*${OUTGOING_FAILURES_4_SERVICE_NOT_PROVIDED}
  ^\s*Cause\sClass\s5\s+.*${OUTGOING_FAILURES_5_INVALID_MESSAGE}
  ^\s*Cause\sClass\s6\s+.*${OUTGOING_FAILURES_6_PROTOCOL}
  ^\s*Cause\sClass\s7\s+.*${OUTGOING_FAILURES_7_INTERWORKING}
  ^\s*Incoming\scalls\s+${INCOMING_CALLS}
  ^\s*Incoming\scalls\sbackup\sinvoked\s+${INCOMING_CALLS_BACKUP_INVOKED}
  ^\s*Incoming\scalls\sfailures\s+${INCOMING_FAILURES}
  ^\s*Remote\sfailure\s+${INCOMING_FAILURES_REMOTE}
  ^\s*Unknown\snumber\s+${INCOMING_FAILURES_UNKOWN_NUMBER}
  ^\s*DSP\sunavailable\s+${INCOMING_FAILURES_DSP_UNAVAILABLE}
  ^\s*No\sVoIP\sresource\savailable\s+${INCOMING_FAILURES_NO_VOIP_RESOURCE_AVAILABLE}
  ^\s*Not\sspecified\s+${INCOMING_FAILURES_NOT_SPECIFIED}
  ^\s+Channel\(s\)\s+used
  ^\s+ces\s*:
  ^\s+Insufficient
  ^\s*$$
  ^. -> Error
`