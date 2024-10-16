package oneaccess_oneos

type ShowRunningConfigBind struct {
	Protocol  string   `json:"PROTOCOL"`
	Interface []string `json:"INTERFACE"`
	Acl       string   `json:"ACL"`
	Vrf       []string `json:"VRF"`
}

var ShowRunningConfigBind_Template string = `Value Required PROTOCOL (ssh|telnet)
Value List INTERFACE (\S+.*\d+)
Value ACL (\S+)
Value List VRF (\S+)
 
Start
  ^bind\sssh\sacl\s${ACL} -> SSH
  ^bind\sssh\s${INTERFACE}$$ -> SSH
  ^bind\sssh\svrf\s${VRF} -> SSH
  ^bind\stelnet\sacl\s${ACL} -> TELNET
  ^bind\stelnet\s${INTERFACE} -> TELNET
  ^bind\stelnet\svrf\s${VRF} -> TELNET
  ^\s*$$
  ^. -> Error
 
SSH
  # at EOF or new protocol, start a record
  # not all entries (ACL, VRF) may be present
  ^bind\stelnet -> Continue.Record
  ^bind\stelnet\sacl\s${ACL} -> Start
  ^bind\stelnet\s${INTERFACE} -> Start
  ^bind\stelnet\svrf\s${VRF} -> Start
  ^bind\s${PROTOCOL}\svrf\s${VRF}
  ^bind\s${PROTOCOL}\s${INTERFACE}
  ^\s*$$
  ^. -> Error
 
TELNET
  # at EOF or new protocol, start a record
  # not all entries (ACL, VRF) may be present
  ^bind\sssh -> Continue.Record
  ^bind\sssh\sacl\s${ACL} -> Start
  ^bind\sssh\s${INTERFACE}$$ -> Start
  ^bind\sssh\svrf\s${VRF} -> Start
  ^bind\s${PROTOCOL}\svrf\s${VRF}
  ^bind\s${PROTOCOL}\s${INTERFACE}
  ^\s*$$
  ^. -> Error
`
