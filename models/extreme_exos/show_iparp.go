package extreme_exos

type ShowIparp struct {
	Static      string `json:"STATIC"`
	Age         string `json:"AGE"`
	Vlan        string `json:"VLAN"`
	Vid         string `json:"VID"`
	Vr          string `json:"VR"`
	Destination string `json:"DESTINATION"`
	Mac_address string `json:"MAC_ADDRESS"`
	Port        string `json:"PORT"`
}

var ShowIparp_Template string = `Value VR (\S+)
Value DESTINATION (\d+\.\d+\.\d+\.\d+)
Value MAC_ADDRESS ([0-9A-Fa-f:]+|incomplete)
Value PORT (\S+)
Value STATIC (\S+)
Value AGE (\d+)
Value VLAN (\S+)
Value VID (\S+)

Start
  ^VR\s+Destination\s+Mac\s+Age\s+Static\s+VLAN\s+VID\s+Port -> Start_record
  ^. -> Error

Start_record
  ^\s*---
  ^${VR}\s+${DESTINATION}\s+\(*${MAC_ADDRESS}\)*\s+${AGE}\s+${STATIC}\s+${VLAN}\s+${VID}(\s+${PORT})?\s*$$ -> Record
  ^(\S+\s*){1,4}:
  ^\s*$$
  ^. -> Error
`
