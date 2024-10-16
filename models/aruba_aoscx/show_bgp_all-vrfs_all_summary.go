package aruba_aoscx

type ShowBgpAllVrfsAllSummary struct {
	Vrf          string `json:"VRF"`
	Af           string `json:"AF"`
	Neighbour    string `json:"NEIGHBOUR"`
	Remote_as    string `json:"REMOTE_AS"`
	State        string `json:"STATE"`
	Admin_status string `json:"ADMIN_STATUS"`
}

var ShowBgpAllVrfsAllSummary_Template string = `Value Filldown VRF (\S+)
Value Filldown AF (\S+)
Value NEIGHBOUR (\d+\.\d+\.\d+\.\d+)
Value REMOTE_AS (\d+)
Value STATE (\w+)
Value ADMIN_STATUS (\w+)

Start
  ^VRF\s+:\s+${VRF}
  ^BGP\sSummary
  ^-+
  ^\s*Local\s*AS.*
  ^\s*Peers.*
  ^\s*Cfg.*
  ^\s*$$
  ^Address-family\s:\s${AF} -> Af
  ^. -> Error

Af
  ^-+
  ^\s+Neighbor\s+Remote-AS\s+MsgRcvd\s+MsgSent\s+Up/Down\s+Time\s+State\s+AdminStatus
  ^\s+${NEIGHBOUR}\s+${REMOTE_AS}\s+\d+\s+\d+\s+\S+\s+${STATE}\s+${ADMIN_STATUS} -> Record
  ^\s*$$ -> Clear Start
  ^. -> Error

EOF`
