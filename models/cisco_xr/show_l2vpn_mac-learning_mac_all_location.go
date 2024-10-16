package cisco_xr

type ShowL2vpnMacLearningMacAllLocation struct {
	Topo_id     string `json:"TOPO_ID"`
	Producer    string `json:"PRODUCER"`
	Next_hop    string `json:"NEXT_HOP"`
	Mac_address string `json:"MAC_ADDRESS"`
}

var ShowL2vpnMacLearningMacAllLocation_Template string = `Value TOPO_ID (\d+)
Value PRODUCER (\S+)
Value NEXT_HOP (\S+)
Value MAC_ADDRESS (\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)

Start
  # Match the timestamp at beginning of command output
  ^\S+\s+\S+\s+\d+\s+\d+:\d+:\d+\.\d+\s+\S+\s*$$
  ^Topo ID\s+Producer\s+Next Hop\(s\)\s+Mac Address\s+IP Address$$
  ^-+\s+-+\s+-+\s+-+\s+-+$$
  ^${TOPO_ID}\s*${PRODUCER}\s*${NEXT_HOP}\s*${MAC_ADDRESS}\s* -> Record
  ^. -> Error
`
