package oneaccess_oneos

type ShowTacacsServer struct {
	Vrf           string `json:"VRF"`
	Tacacs_server string `json:"TACACS_SERVER"`
	Server_port   string `json:"SERVER_PORT"`
	Secret_key    string `json:"SECRET_KEY"`
	Source        string `json:"SOURCE"`
}

var ShowTacacsServer_Template string = `Value TACACS_SERVER (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value SERVER_PORT (\d+)
Value SECRET_KEY (\w+)
Value SOURCE (\w+\s[\w\/\.]+)
Value VRF ([\w\-]+)


Start
  ^\s*${TACACS_SERVER}\s+${SERVER_PORT}\s+${SECRET_KEY}\s+${SOURCE}\s*(?:Default-Router)?${VRF}?$$ -> Record
  ^\s*-+\s+List\s+of\s+TACACS\+\s+server\s+-+\s*$$
  ^\s*IP\s+address\s+Port(\s+number)?\s+Secret\s+Key\s+(Source\s+address|Interface)\s+VRF\s*$$
  ^\s*-+\s*$$
  ^\s*$$
  ^. -> Error
`
