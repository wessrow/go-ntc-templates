package cisco_nvfis

type ShowRunningConfigSnmpUser struct {
	Priv_protocol string `json:"PRIV_PROTOCOL"`
	Auth_key      string `json:"AUTH_KEY"`
	Priv_key      string `json:"PRIV_KEY"`
	Username      string `json:"USERNAME"`
	Version       string `json:"VERSION"`
	User_group    string `json:"USER_GROUP"`
	Auth_protocol string `json:"AUTH_PROTOCOL"`
}

var ShowRunningConfigSnmpUser_Template string = `Value USERNAME ([-\w]+)
Value VERSION (\d+)
Value USER_GROUP (\w+)
Value AUTH_PROTOCOL (\w+)
Value PRIV_PROTOCOL (\w+)
Value AUTH_KEY ([0-9A-f:]+)
Value PRIV_KEY ([0-9A-f:]+)


Start
  ^snmp\suser\s${USERNAME}
  ^\suser-version\s+${VERSION}
  ^\suser-group\s+${USER_GROUP}
  ^\sauth-protocol\s+${AUTH_PROTOCOL}
  ^\spriv-protocol\s+${PRIV_PROTOCOL}
  ^\sauth-key\s+${AUTH_KEY}
  ^\spriv-key\s+${PRIV_KEY} -> Record

EOF`
