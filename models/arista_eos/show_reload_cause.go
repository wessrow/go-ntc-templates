package arista_eos 

type ShowReloadCause struct {
	Reload_cause	string	`json:"RELOAD_CAUSE"`
	Reload_time	string	`json:"RELOAD_TIME"`
	Recommended_action	string	`json:"RECOMMENDED_ACTION"`
	Debug_info	string	`json:"DEBUG_INFO"`
}

var ShowReloadCause_Template = `Value RELOAD_CAUSE (.+)
Value RELOAD_TIME (.+)
Value RECOMMENDED_ACTION (.+)
Value DEBUG_INFO (.+)


Start
  ^Reload Cause 1:
  ^--.+
  ^${RELOAD_CAUSE} ->  RT
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
RT  
  ^\s+
  ^Reload Time:
  ^--.+
  ^${RELOAD_TIME} ->  RA
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
RA 
  ^\s+
  ^Recommended Action:
  ^--.+
  ^${RECOMMENDED_ACTION} -> DI
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
DI  
  ^\s+
  ^Debugging Information:
  ^--.+
  ^${DEBUG_INFO} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF  
`