package aruba_aoscx

type ShowAaaAuthenticationPortAccessInterfaceAllClientStatus struct {
	Client_mac   string `json:"CLIENT_MAC"`
	Client_name  string `json:"CLIENT_NAME"`
	Session_port string `json:"SESSION_PORT"`
	Role         string `json:"ROLE"`
}

var ShowAaaAuthenticationPortAccessInterfaceAllClientStatus_Template string = `Value CLIENT_MAC (\S+)
Value CLIENT_NAME (\S+)
Value SESSION_PORT (\S+)
Value ROLE (\S+)

Start
  ^\s*Port\s+Access\s+Client\s+Status\s+Details
  ^\s*Client\s+${CLIENT_MAC},\s+${CLIENT_NAME}
  ^\s*=+
  ^\s+Session\s+Details -> Session
  ^\s+Authentication\s+Details -> Authentication
  ^\s+Authorization\s+Details -> Authorization
  ^\s*$$
  ^. -> Error

Session
  ^\s+-+
  ^\s+Port\s+:\s*${SESSION_PORT}
  ^\s+Session.*$$
  ^\s+IPv4.*$$
  ^\s+IPv6.*$$
  ^\s*$$ -> Start
  ^. -> Error

Authentication
  ^\s+-+
  ^\s+Status.*$$
  ^\s+Auth\s*Precedence.*$$ -> Start
  ^\s*$$ -> Start
  ^. -> Error

Authorization
  ^\s*-+
  ^\s+Role\s+:\s+${ROLE} 
  ^\s+Status.*$$
  ^\s*$$ -> Record Start
  ^. -> Error`
