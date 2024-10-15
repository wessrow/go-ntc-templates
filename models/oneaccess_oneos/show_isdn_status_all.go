package oneaccess_oneos

type ShowIsdnStatusAll struct {
	Pri_rdi_occurrences   string   `json:"PRI_RDI_OCCURRENCES"`
	Protocol_descriptor   string   `json:"PROTOCOL_DESCRIPTOR"`
	Protocol_framing      string   `json:"PROTOCOL_FRAMING"`
	Config_state          string   `json:"CONFIG_STATE"`
	L1_status             string   `json:"L1_STATUS"`
	Ais                   string   `json:"AIS"`
	Los                   string   `json:"LOS"`
	Pri_ais_occurrences   string   `json:"PRI_AIS_OCCURRENCES"`
	Pri_los_occurrences   string   `json:"PRI_LOS_OCCURRENCES"`
	Protocol_linecode     string   `json:"PROTOCOL_LINECODE"`
	L3_active_calls       string   `json:"L3_ACTIVE_CALLS"`
	L3_call_details       []string `json:"L3_CALL_DETAILS"`
	L2_status             string   `json:"L2_STATUS"`
	Physical_type         string   `json:"PHYSICAL_TYPE"`
	Loop_state            string   `json:"LOOP_STATE"`
	Rai                   string   `json:"RAI"`
	Isdn_line             string   `json:"ISDN_LINE"`
	L2_rx_frames_dchannel string   `json:"L2_RX_FRAMES_DCHANNEL"`
	L2_tx_frames_dchannel string   `json:"L2_TX_FRAMES_DCHANNEL"`
}

var ShowIsdnStatusAll_Template string = `Value Required ISDN_LINE (\d+\/\d+)
Value PHYSICAL_TYPE (\S+)
Value PROTOCOL_DESCRIPTOR (\S+)
Value PROTOCOL_LINECODE (\S+)
Value PROTOCOL_FRAMING (\S+)
Value CONFIG_STATE (\S+)
Value LOOP_STATE (\S+)
Value L1_STATUS (\S+)
Value AIS (\S+)
Value LOS (\S+)
Value RAI (\S+)
Value PRI_AIS_OCCURRENCES (\d+)
Value PRI_LOS_OCCURRENCES (\d+)
Value PRI_RDI_OCCURRENCES (\d+)
Value L2_STATUS (\S+)
Value L2_TX_FRAMES_DCHANNEL (\d+)
Value L2_RX_FRAMES_DCHANNEL (\d+)
Value L3_ACTIVE_CALLS (\d+|no active call)
Value List L3_CALL_DETAILS (.*)

Start
  ^\s*isdn\sline\s+${ISDN_LINE} -> ISDNLINE
  ^\s*$$
  ^. -> Error

ISDNLINE
  ^\s+physical\stype\s+${PHYSICAL_TYPE}
  ^\s+protocol\sdescriptor\s+${PROTOCOL_DESCRIPTOR}
  ^\s+linecode\s+${PROTOCOL_LINECODE}
  ^\s+framing\s+${PROTOCOL_FRAMING}
  ^\s+config\sstate\s+${CONFIG_STATE}
  ^\s+loop\sstate\s+${LOOP_STATE}
  ^\s+\-layer\s1\sstatus\s+${L1_STATUS}
  ^\s*Alarm\sIndication\sSignal\s+\(AIS\)\s+${AIS}
  ^\s*Loss\sOff\sSignal\s+\(LOS\)\s+${LOS}
  ^\s*Remote\sIndication\sSignal\s+\(RAI\)\s+${RAI}
  ^\s*pri\sAIS\soccurrence\(s\)\s+${PRI_AIS_OCCURRENCES}
  ^\s*pri\sLOS\soccurrence\(s\)\s+${PRI_LOS_OCCURRENCES}
  ^\s*pri\sRDI\soccurrence\(s\)\s+${PRI_RDI_OCCURRENCES}
  ^\s+\-layer\s2\sstatus\s+${L2_STATUS}
  ^\s+Tx\sframes\son\sD\schannel\s+${L2_TX_FRAMES_DCHANNEL}
  ^\s+Rx\sframes\son\sD\schannel\s+${L2_RX_FRAMES_DCHANNEL}
  ^\s+-layer\s3\sstatus -> L3Status
  ^\s*$$ -> Record
  ^\s*isdn\sline\s+${ISDN_LINE} -> ISDNLINE
  ^\s+ces=\s+
  ^. -> Error

L3Status
  ^\s*(active\scall\s*|)${L3_ACTIVE_CALLS}
  ^\s*->\s${L3_CALL_DETAILS}
  ^\s*$$ -> Record ISDNLINE
  ^. -> Error

`
