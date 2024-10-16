package oneaccess_oneos

type ShowIsdnLedStatus struct {
	Sip_server_status          string `json:"SIP_SERVER_STATUS"`
	Voice_com_status           string `json:"VOICE_COM_STATUS"`
	Sysled_maintenance_state   string `json:"SYSLED_MAINTENANCE_STATE"`
	Sysled_bri1_state          string `json:"SYSLED_BRI1_STATE"`
	Dialpeer_voip_up           string `json:"DIALPEER_VOIP_UP"`
	Sip_gateway_nbr_registered string `json:"SIP_GATEWAY_NBR_REGISTERED"`
	Sip_gateway_nbr_available  string `json:"SIP_GATEWAY_NBR_AVAILABLE"`
	Sip_gateway_registrar      string `json:"SIP_GATEWAY_REGISTRAR"`
	Sip_gateway_intf           string `json:"SIP_GATEWAY_INTF"`
	Sip_gateway_intf_ip        string `json:"SIP_GATEWAY_INTF_IP"`
	Track_cond_voice_gw        string `json:"TRACK_COND_VOICE_GW"`
	Sysled_com_state           string `json:"SYSLED_COM_STATE"`
	Dialpeer_pots_up           string `json:"DIALPEER_POTS_UP"`
	Sip_gateway_status         string `json:"SIP_GATEWAY_STATUS"`
	Sysled_bri1_color          string `json:"SYSLED_BRI1_COLOR"`
	Sysled_bri2_state          string `json:"SYSLED_BRI2_STATE"`
	Sysled_maintenance_color   string `json:"SYSLED_MAINTENANCE_COLOR"`
	Sysled_pri_color           string `json:"SYSLED_PRI_COLOR"`
	Sysled_pri_state           string `json:"SYSLED_PRI_STATE"`
	Track_cond_voice_port      string `json:"TRACK_COND_VOICE_PORT"`
	Sysled_voip_color          string `json:"SYSLED_VOIP_COLOR"`
	Sysled_voip_state          string `json:"SYSLED_VOIP_STATE"`
	Sysled_com_color           string `json:"SYSLED_COM_COLOR"`
	Sysled_bri2_color          string `json:"SYSLED_BRI2_COLOR"`
	Sip_gateway_registered     string `json:"SIP_GATEWAY_REGISTERED"`
	Sip_gateway_intf_status    string `json:"SIP_GATEWAY_INTF_STATUS"`
}

var ShowIsdnLedStatus_Template string = `Value DIALPEER_POTS_UP (\d+)
Value DIALPEER_VOIP_UP (\d+)
Value SIP_GATEWAY_STATUS (no shutdown|shutdown|[nN]o Sip-gateway)
Value SIP_GATEWAY_REGISTERED (.*registered.*)
Value SIP_GATEWAY_NBR_REGISTERED (\d+)
Value SIP_GATEWAY_NBR_AVAILABLE (\d+)
Value SIP_GATEWAY_REGISTRAR (\S+:\d+)
Value SIP_GATEWAY_INTF (\S+\s\S+)
Value SIP_GATEWAY_INTF_IP (\S+)
Value SIP_GATEWAY_INTF_STATUS (\S+)
Value SIP_SERVER_STATUS (no shutdown|shutdown|[nN]o Sip-server)
Value VOICE_COM_STATUS (no shutdown|shutdown|[nN]o voice com)
Value TRACK_COND_VOICE_GW (\S+)
Value TRACK_COND_VOICE_PORT (\S+)
Value SYSLED_VOIP_COLOR (\S+|\-)
Value SYSLED_VOIP_STATE (\S+)
Value SYSLED_COM_COLOR (\S+|\-)
Value SYSLED_COM_STATE (\S+)
Value SYSLED_MAINTENANCE_COLOR (\S+|\-)
Value SYSLED_MAINTENANCE_STATE (\S+)
Value SYSLED_PRI_COLOR (\S+|\-)
Value SYSLED_PRI_STATE (\S+)
Value SYSLED_BRI1_COLOR (\S+|\-)
Value SYSLED_BRI1_STATE (\S+)
Value SYSLED_BRI2_COLOR (\S+|\-)
Value SYSLED_BRI2_STATE (\S+)


Start
  ^\s*${DIALPEER_POTS_UP}\sDial\-peer\spots\sup
  ^\s*${DIALPEER_VOIP_UP}\sDial\-peer\svoip\sup
  ^\s*Sip\-gateway\sstatus\sis\s${SIP_GATEWAY_STATUS} -> SipGateway
  ^\s*${SIP_GATEWAY_STATUS}
  ^\s*Sip\-server\sstatus\sis\s${SIP_SERVER_STATUS}
  ^\s*${SIP_SERVER_STATUS}
  ^\s*voice\s+com\sstatus\sis\s${VOICE_COM_STATUS}
  ^\s*${VOICE_COM_STATUS}
  ^\s*voice\sled\strack-conditions -> Track
  ^\s*Sys\s+LEDs -> SysLed
  ^\s*Status\s+of\s+VOIP\s+service
  ^\s*Isdn\s+virtual\s+LEDs
  ^\s*(isdnshut|Layer1State|nbcomm)
  ^\s*$$
  ^. -> Error

SipGateway
  ^\s*->\sis\s${SIP_GATEWAY_REGISTERED}
  ^\s*${SIP_GATEWAY_NBR_REGISTERED}\/${SIP_GATEWAY_NBR_AVAILABLE}\sep\sis.*with\sregistrar\s${SIP_GATEWAY_REGISTRAR}
  ^\s*->\sIF\s${SIP_GATEWAY_INTF}\s\[${SIP_GATEWAY_INTF_IP}\]\sis\s${SIP_GATEWAY_INTF_STATUS}\.
  ^\s*$$ -> Start
  ^. -> Error

Track
  ^\s*voice-gw\s+${TRACK_COND_VOICE_GW}
  ^\s*voice-port\s+${TRACK_COND_VOICE_PORT}
  ^\s*$$ -> Start
  ^. -> Error

SysLed
  ^\s*SYS\sLED\s(VoIP|VOIP)\s*color=\s*${SYSLED_VOIP_COLOR}\s*,?\sstate=${SYSLED_VOIP_STATE}
  ^\s*SYS\sLED\sCOM\s*color=\s*${SYSLED_COM_COLOR}\s*,?\sstate=${SYSLED_COM_STATE}
  ^\s*SYS\sLED\sMaintenance\s*color=\s*${SYSLED_MAINTENANCE_COLOR}\s*,?\sstate=${SYSLED_MAINTENANCE_STATE}
  ^\s*SYS\sLED\sPRI\s*color=\s*${SYSLED_PRI_COLOR}\s*,?\sstate=${SYSLED_PRI_STATE}
  ^\s*SYS\sLED\sBRI1\s*color=\s*${SYSLED_BRI1_COLOR}\s*,?\sstate=${SYSLED_BRI1_STATE}
  ^\s*SYS\sLED\sBRI2\s*color=\s*${SYSLED_BRI2_COLOR}\s*,?\sstate=${SYSLED_BRI2_STATE}
  ^\s*$$
  ^. -> Error
`
