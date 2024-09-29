package cisco_ios 

type ShowDmvpn struct {
	Peer_nbma	string	`json:"PEER_NBMA"`
	Peer_tunnel	string	`json:"PEER_TUNNEL"`
	State	string	`json:"STATE"`
	Uptime	string	`json:"UPTIME"`
	Attribute	string	`json:"ATTRIBUTE"`
}

var ShowDmvpn_Template = `Value PEER_NBMA (\S*?)
Value PEER_TUNNEL (\S*?)
Value STATE (\S*?)
Value UPTIME (\S*?)
Value ATTRIBUTE (\S*?)

Start
  ^\s+\d+\s+${PEER_NBMA}\s+${PEER_TUNNEL}\s+${STATE}\s+${UPTIME}\s+${ATTRIBUTE}$$ -> Record
  ^Legend.*
  ^\s+\D.*.*$$
  ^=*$$
  ^Interface.*$$
  ^Type.*$$
  ^\s*$$
  ^\s+#\s+Ent\s+Peer\s+NBMA\s+Addr\s+Peer\s+Tunnel\s+Add\s+State\s+UpDn\s+Tm\s+Attrb
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error
`