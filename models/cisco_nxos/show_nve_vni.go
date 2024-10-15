package cisco_nxos

type ShowNveVni struct {
	Mcast_group string `json:"MCAST_GROUP"`
	State       string `json:"STATE"`
	Mode        string `json:"MODE"`
	Type        string `json:"TYPE"`
	Bd_vrf      string `json:"BD_VRF"`
	Interface   string `json:"INTERFACE"`
	Vni         string `json:"VNI"`
}

var ShowNveVni_Template string = `Value INTERFACE (\S+)
Value VNI (\d+)
Value MCAST_GROUP (\S+)
Value STATE (Up|Down)
Value MODE (\S+)
Value TYPE (L2|L3)
Value BD_VRF (\S+)

Start
  ^Codes:\s+CP.*
  ^\s+UC\s+-\s+Unconfigured.*
  ^\s+SU\s+-\s+Suppress.*
  ^\s+Xconn\s+-\s+Crossconnect.*
  ^\s+MS-IR\s+-\s+Multisite.*
  ^Interface\s+VNI\s+Multicast-group\s+State\s+Mode\s+Type\s+\[BD/VRF\]\s+Flags\s*$$
  ^-+
  ^${INTERFACE}\s+${VNI}\s+${MCAST_GROUP}\s+${STATE}\s+${MODE}\s+${TYPE}\s+\[${BD_VRF}\] -> Record
  ^\s*$$
  ^. -> Error
`
