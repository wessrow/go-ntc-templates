package hp_procurve

type ShowPortSecurity struct {
	Learn_mode           string `json:"LEARN_MODE"`
	Action               string `json:"ACTION"`
	Eavesdrop_prevention string `json:"EAVESDROP_PREVENTION"`
	Port                 string `json:"PORT"`
}

var ShowPortSecurity_Template string = `Value PORT (\S+)
Value LEARN_MODE (\S+)
Value ACTION (None|Send\s+Alarm|Send\s+Alarm\S+\s+Disable\s+Port)
Value EAVESDROP_PREVENTION (Enabled|Disabled)

Start
  ^\s+Port\s+Security\s*
  ^\s*$$
  ^\s+Port\s+Learn\s+Mode\s+\|\s+Action\s+Eavesdrop\s+Prevention\s*$$
  ^\s*-+(?:\s|-|\+)+$$ -> Show_Port_Security
  ^. -> Error

Show_Port_Security
  ^\s+${PORT}\s+${LEARN_MODE}\s+\|\s+${ACTION}\s+${EAVESDROP_PREVENTION}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
