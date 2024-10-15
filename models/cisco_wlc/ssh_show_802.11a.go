package cisco_wlc

type SshShow80211a struct {
	Data_rate_value         []string `json:"DATA_RATE_VALUE"`
	Mcs_rate                []string `json:"MCS_RATE"`
	Voice_cac_method        string   `json:"VOICE_CAC_METHOD"`
	Video_amc               string   `json:"VIDEO_AMC"`
	Video_roaming_bandwidth string   `json:"VIDEO_ROAMING_BANDWIDTH"`
	Status                  string   `json:"STATUS"`
	Mcs_rate_value          []string `json:"MCS_RATE_VALUE"`
	Ss_mcs_rate             []string `json:"SS_MCS_RATE"`
	Voice_amc               string   `json:"VOICE_AMC"`
	Specs_value             []string `json:"SPECS_VALUE"`
	Video_cac_method        string   `json:"VIDEO_CAC_METHOD"`
	Video_max_bandwidth     string   `json:"VIDEO_MAX_BANDWIDTH"`
	Max_client              string   `json:"MAX_CLIENT"`
	Rssi_thres              string   `json:"RSSI_THRES"`
	Edca                    string   `json:"EDCA"`
	Voice_max_bandwidth     string   `json:"VOICE_MAX_BANDWIDTH"`
	Specs                   []string `json:"SPECS"`
	Data_rate               []string `json:"DATA_RATE"`
	Ss_mcs_rate_value       []string `json:"SS_MCS_RATE_VALUE"`
	Beacon_interval         string   `json:"BEACON_INTERVAL"`
	Rssi_low_check          string   `json:"RSSI_LOW_CHECK"`
	Voice_roaming_bandwidth string   `json:"VOICE_ROAMING_BANDWIDTH"`
}

var SshShow80211a_Template string = `Value STATUS (\w+)
Value List SPECS (\w+)
Value List SPECS_VALUE (\w+)
Value List DATA_RATE (\d+\.?\d?\w)
Value List DATA_RATE_VALUE (\w+)
Value List MCS_RATE (\d+)
Value List MCS_RATE_VALUE (\w+)
Value List SS_MCS_RATE (Nss=\d\:\sMCS\s0-\d)
Value List SS_MCS_RATE_VALUE (\w+)
Value BEACON_INTERVAL (\d+)
Value RSSI_LOW_CHECK (\w+)
Value RSSI_THRES (-\d+)
Value EDCA (.+?)
Value VOICE_AMC (\w+)
Value VOICE_CAC_METHOD (.+?)
Value VOICE_MAX_BANDWIDTH (\d+)
Value VOICE_ROAMING_BANDWIDTH (\d+|)
Value VIDEO_AMC (\w+)
Value VIDEO_CAC_METHOD (.+?)
Value VIDEO_MAX_BANDWIDTH (\d+)
Value VIDEO_ROAMING_BANDWIDTH (\d+|)
Value MAX_CLIENT (\d+)



Start
  ^802.11(a|b)\sNetwork\.*\s${STATUS}s*$$ -> Spec_support
  ^Beacon\sInterval\.*\s${BEACON_INTERVAL}s*$$
  ^RSSI\sLow\sCheck\.*\s${RSSI_LOW_CHECK}s*$$
  ^RSSI\sThreshold\.*\s${RSSI_THRES}s*$$
  ^EDCA\sprofile\stype\.*\s${EDCA}s*$$
  ^Voice\s+AC:\s*$$ -> Voice
  ^Video\s+AC:\s*$$ -> Video
  ^802\.11n\s+Status:\s*$$ -> 80211n
  ^Maximum\sNumber\sof\sClients\sper\sAP(\sRadio)?\.*\s${MAX_CLIENT}s*$$
  ^CF\s+Poll
  ^CFP\s+[Period|Maximum]
  ^Default\s+Channel
  ^Default\s+Tx\s+Power\s+Level
  ^DTPC\s+Status
  ^Fragmentation\s+Threshold
  ^RSSI\s+
  ^TI\s+Threshold
  ^Legacy\s+Tx\s+Beamforming setting
  ^Traffic\s+Stream\s+Metrics Status
  ^Expedited\s+BW\s+Request
  ^World\s+Mode
  ^dfs-peakdetect
  ^Voice\s+MAC\s+optimization
  ^Call\s+Admission\s+Control
  ^\s*$$
  ^. -> Error

Spec_support
  ^11${SPECS}Support\.*\s${SPECS_VALUE}s*$$
  ^802.11(a|b|b\/g)\sOperational\sRates -> Data_Rates
  ^\s*802.11\S\s+(?:Low|Mid|High)\s+Band
  ^\s*$$
  ^. -> Error

Data_Rates
  ^\s+802.11(a|g|b|(b\/g))\s${DATA_RATE}\s+Rate\.*\s+${DATA_RATE_VALUE}s*$$
  ^802.11n\s+MCS\s+Settings: -> MCS_Rates
  ^\s*
  ^. -> Error

MCS_Rates
  ^\s+MCS\s${MCS_RATE}\.*\s+${MCS_RATE_VALUE}s*$$
  ^802.11ac\s+MCS\s+Settings: -> AC_MCS_Rates
  ^802.11n Status: -> 80211n
  ^\s*$$
  ^. -> Error

AC_MCS_Rates
  ^\s+${SS_MCS_RATE}\s+\.*\s${SS_MCS_RATE_VALUE}s*$$
  ^802.11n\s+Status:\s*$$ -> 80211n
  ^\s*$$
  ^. -> Error

80211n
  ^\s+A-M[P|S]DU\s+Tx
  ^\s+Priority\s+\d+
  ^\s+Aggregation\s+scheduler
  ^\s+Frame\s+Burst
  ^\s+Realtime\s+Timeout
  ^\s+Non\s+Realtime\s+Timeout
  ^Beacon\sInterval\.*\s${BEACON_INTERVAL}s*$$ -> Start
  ^EDCA\sprofile\stype\.*\s${EDCA}s*$$
  ^Maximum\sNumber\sof\sClients\sper\sAP(\sRadio)?\.*\s${MAX_CLIENT}s*$$
  ^\S -> Start

Voice
  ^\s+Voice\sAC\s-\sAdmission\scontrol\s\(ACM\)\.*\s${VOICE_AMC}s*$$
  ^\s+Voice\sCAC\sMethod.*\s${VOICE_CAC_METHOD}s*$$
  ^\s+Voice\smax\sRF\sbandwidth\.*\s${VOICE_MAX_BANDWIDTH}s*$$
  ^\s+Voice\sreserved\sroaming\sbandwidth\.*\s${VOICE_ROAMING_BANDWIDTH}s*$$
  ^Video\s+AC:\s*$$ -> Video
  ^Maximum\sNumber\sof\sClients\sper\sAP(\sRadio)?\.*\s${MAX_CLIENT}s*$$
  ^\S -> Start

Video
  ^\s+Video\sAC\s-\sAdmission\scontrol\s\(ACM\)\.*\s${VIDEO_AMC}s*$$
  ^\s+Video\sCAC\sMethod\s\.*\s${VIDEO_CAC_METHOD}s*$$
  ^\s+Video\smax\sRF\sbandwidth\.*\s${VIDEO_MAX_BANDWIDTH}s*$$
  ^\s+Video\sreserved\sroaming\sbandwidth\.*\s${VIDEO_ROAMING_BANDWIDTH}s*$$
  ^Voice\s+AC:\s*$$ -> Voice
  ^Maximum\sNumber\sof\sClients\sper\sAP(\sRadio)?\.*\s${MAX_CLIENT}s*$$
  ^\S -> Start
`
