package arista_eos

type ShowInterfacesTransceiverDetail struct {
	Port                 string `json:"PORT"`
	Type                 string `json:"TYPE"`
	Current_value        string `json:"CURRENT_VALUE"`
	High_alarm_threshold string `json:"HIGH_ALARM_THRESHOLD"`
	High_warn_threshold  string `json:"HIGH_WARN_THRESHOLD"`
	Low_alarm_threshold  string `json:"LOW_ALARM_THRESHOLD"`
	Low_warn_threshold   string `json:"LOW_WARN_THRESHOLD"`
}

var ShowInterfacesTransceiverDetail_Template string = `Value Required PORT (\S+)
Value Filldown TYPE ((\S|\s)+?)
Value CURRENT_VALUE (\S+)
Value HIGH_ALARM_THRESHOLD (\S+)
Value HIGH_WARN_THRESHOLD (\S+)
Value LOW_ALARM_THRESHOLD (\S+)
Value LOW_WARN_THRESHOLD (\S+)


Start
  ^mA:\smilliamperes,\s+
  ^\+\+\s:\shigh\salarm,\s+
  ^A2D\sreadouts\s+
  ^The\sthreshold\svalues\sare\scalibrated\.
  ^\s+High\sAlarm\s+\S+
  ^\s+${TYPE}\s+Threshold\s+ -> Values
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
Values
  ^Port\s+\S+
  ^-+\s+-+\s+-+\s+-+\s+-+\s+-+\s+
  ^${PORT}\s+${CURRENT_VALUE}\s+${HIGH_ALARM_THRESHOLD}\s+${HIGH_WARN_THRESHOLD}\s+${LOW_ALARM_THRESHOLD}\s+${LOW_WARN_THRESHOLD} -> Record
  ^\s+High\sAlarm\s+High Warn\s+Low\sAlarm\s+Low\sWarn -> Start
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
`
