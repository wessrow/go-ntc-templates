package cisco_ios

type ShowSwitchDetail struct {
	Switch      string `json:"SWITCH"`
	Role        string `json:"ROLE"`
	Mac_address string `json:"MAC_ADDRESS"`
	Priority    string `json:"PRIORITY"`
	Version     string `json:"VERSION"`
	State       string `json:"STATE"`
}

var ShowSwitchDetail_Template string = `Value Key SWITCH (\d+)
Value ROLE (\w+)
Value MAC_ADDRESS (\S+)
Value PRIORITY (\d+)
Value VERSION (\S+)
Value STATE (\w+)

Start
  ^Switch/Stack
  ^Mac\s+persistency\s+wait\s+time
  ^\s+H/W\s+Current\s*$$
  ^Switch#\s+Role\s+Mac\s+Address\s+Priority\s+Version\s+State\s*$$
  ^\s*-+ -> Status
  ^\s*$$
  ^. -> Error

Status
  ^\*?\s*${SWITCH}\s+${ROLE}\s+${MAC_ADDRESS}\s+${PRIORITY}\s+${STATE}\s*$$ -> Record
  ^\*?\s*${SWITCH}\s+${ROLE}\s+${MAC_ADDRESS}\s+${PRIORITY}\s+${VERSION}\s+${STATE}\s*$$ -> Record
  ^\s*Stack\s+Port\s+Status -> Stack
  ^\s*$$
  ^. -> Error

Stack
  ^.
`
