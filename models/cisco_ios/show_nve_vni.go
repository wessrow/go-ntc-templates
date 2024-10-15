package cisco_ios

type ShowNveVni struct {
	State       string `json:"STATE"`
	Mode        string `json:"MODE"`
	Bd          string `json:"BD"`
	Vrf         string `json:"VRF"`
	Interface   string `json:"INTERFACE"`
	Vni         string `json:"VNI"`
	Mcast_group string `json:"MCAST_GROUP"`
}

var ShowNveVni_Template string = `Value INTERFACE (\S+)
Value VNI (\d+)
Value MCAST_GROUP (\S+)
Value STATE (Up|Down)
Value MODE (L2CP|L3CP)
Value BD (\d+)
Value VRF (\S+)

Start
  ^Interface\s+VNI\s+Multicast-group\s+VNI\s+state\s+Mode\s+BD\s+cfg\s+vrf\s*$$
  ^${INTERFACE}\s+${VNI}\s+${MCAST_GROUP}\s+${STATE}\s+${MODE}\s+${BD}\s+CLI\s+${VRF} -> Record
  ^\s*$$
  ^. -> Error
  `
