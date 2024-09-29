package oneaccess_oneos 

type ShowIpSsh struct {
	Ssh	string	`json:"SSH"`
	Scp	string	`json:"SCP"`
	Encryption	string	`json:"ENCRYPTION"`
	Bits	string	`json:"BITS"`
	Fingerprint	string	`json:"FINGERPRINT"`
	Max_sessions	string	`json:"MAX_SESSIONS"`
	Max_channels_session	string	`json:"MAX_CHANNELS_SESSION"`
	Method	string	`json:"METHOD"`
	Session_timeout	string	`json:"SESSION_TIMEOUT"`
	Auth_timeout	string	`json:"AUTH_TIMEOUT"`
	Auth_retries	string	`json:"AUTH_RETRIES"`
}

var ShowIpSsh_Template = `Value SSH (\S+)
Value SCP (\S+)
Value ENCRYPTION (ssh\-\S+|\S+)
Value BITS (\d+)
Value FINGERPRINT ([a-f0-9:]+)
Value MAX_SESSIONS (\d+)
Value MAX_CHANNELS_SESSION (\d+)
Value METHOD (\S+)
Value SESSION_TIMEOUT (\d+)
Value AUTH_TIMEOUT (\d+)
Value AUTH_RETRIES (\d+)
 
Start
  ^SSH\s${SSH}
  ^SCP\sserver\s${SCP}
  ^Authentication\s.ethod:?\s${METHOD}
  ^Authentication\stimeout\s${AUTH_TIMEOUT}\ssecs,\sretries\s${AUTH_RETRIES}
  ^Session\stimeout\s${SESSION_TIMEOUT}\ssecs
  ^Maximum\snumber\sof\ssessions\s${MAX_SESSIONS}
  ^Maximum\snumber\sof\schannels\sper\ssession\s${MAX_CHANNELS_SESSION}
  ^Authorized\s+public\s+keys:
  ^none
  ^Key\sfingerprint -> FINGERPRINT
  ^\s*$$
  ^. -> Error
 
FINGERPRINT
  ^${ENCRYPTION}\s${BITS}\s${FINGERPRINT}\s*$$ -> Start
  ^${BITS}\sMD5:${FINGERPRINT}.*\s\(${ENCRYPTION}\)$$ -> Start
  ^\s*$$
  ^. -> Error
`