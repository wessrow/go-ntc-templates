package cisco_wlc

type SshShowMobilityAnchor struct {
	Status     string `json:"STATUS"`
	Priority   string `json:"PRIORITY"`
	Wlan_id    string `json:"WLAN_ID"`
	Ip_address string `json:"IP_ADDRESS"`
}

var SshShowMobilityAnchor_Template string = `Value WLAN_ID (\d+)
Value IP_ADDRESS (\S+)
Value STATUS (\S+)
Value PRIORITY ([1-3])

Start
  ^\s*WLAN ID\s+IP Address\s+Status\s+Priority\s*$$ -> Mobility_Anchor

Mobility_Anchor
  # WLAN Mobility Anchor List
  ^\s+${WLAN_ID}\s+${IP_ADDRESS}\s+${STATUS}\s+${PRIORITY} -> Record
  #
  # also handling the similar Guest LAN (GLAN) output
  ^\s+${WLAN_ID}\s+${IP_ADDRESS}\s+${STATUS} -> Record
  #
  ^\s*GLAN ID\s+IP Address\s+Status\s*$$
  ^\s+[-\s]+$$
  ^\s*$$
  ^. -> Error
`
