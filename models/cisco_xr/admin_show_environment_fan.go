package cisco_xr

type AdminShowEnvironmentFan struct {
	Module    string   `json:"MODULE"`
	Type      string   `json:"TYPE"`
	Fan_speed []string `json:"FAN_SPEED"`
}

var AdminShowEnvironmentFan_Template string = `Value MODULE (\S+)
Value TYPE (\S+)
Value List FAN_SPEED (\d+)

Start
  ^.+UTC
  ^=+
  ^\s+Fan\s+speed\s+\(rpm\)\s*$$
  ^Location\s+FRU\s+Type.+FAN_\d
  ^\s+FAN_.+$$
  ^-+ -> FM
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
  
FM
  ^${MODULE}\s+${TYPE}\s+${FAN_SPEED}\s*$$ -> Record
  ^${MODULE}\s+${TYPE}\s+${FAN_SPEED} -> Continue
  ^${MODULE}\s+${TYPE}\s+\d+\s+${FAN_SPEED}\s*$$ -> Record
  ^${MODULE}\s+${TYPE}\s+\d+\s+${FAN_SPEED} -> Continue
  ^${MODULE}\s+${TYPE}\s+(?:\d+\s+){2}${FAN_SPEED}\s*$$ -> Record
  ^${MODULE}\s+${TYPE}\s+(?:\d+\s+){2}${FAN_SPEED} -> Continue
  ^${MODULE}\s+${TYPE}\s+(?:\d+\s+){3}${FAN_SPEED}\s*$$ -> Record
  ^${MODULE}\s+${TYPE}\s+(?:\d+\s+){3}${FAN_SPEED} -> Continue
  ^${MODULE}\s+${TYPE}\s+(?:\d+\s+){4}${FAN_SPEED}\s*$$ -> Record
  ^${MODULE}\s+${TYPE}\s+(?:\d+\s+){4}${FAN_SPEED} -> Continue
  ^${MODULE}\s+${TYPE}\s+(?:\d+\s+){5}${FAN_SPEED}
  ^\s+${FAN_SPEED}\s*$$ -> Record
  ^\s+${FAN_SPEED} -> Continue
  ^\s+\d+\s+${FAN_SPEED}\s*$$ -> Record
  ^\s+\d+\s+${FAN_SPEED} -> Continue
  ^(?:\s+\d+\s+){2}${FAN_SPEED}\s*$$ -> Record
  ^(?:\s+\d+\s+){2}${FAN_SPEED} -> Continue
  ^(?:\s+\d+\s+){3}${FAN_SPEED}\s*$$ -> Record
  ^(?:\s+\d+\s+){3}${FAN_SPEED} -> Continue
  ^(?:\s+\d+\s+){4}${FAN_SPEED}\s*$$ -> Record
  ^(?:\s+\d+\s+){4}${FAN_SPEED} -> Continue
  ^(?:\s+\d+\s+){5}${FAN_SPEED} -> Record
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
`
