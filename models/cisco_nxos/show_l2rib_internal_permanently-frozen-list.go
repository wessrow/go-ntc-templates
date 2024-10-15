package cisco_nxos

type ShowL2ribInternalPermanentlyFrozenList struct {
	Topology    string `json:"TOPOLOGY"`
	Mac_address string `json:"MAC_ADDRESS"`
	Frozen_time string `json:"FROZEN_TIME"`
}

var ShowL2ribInternalPermanentlyFrozenList_Template string = `Value Required TOPOLOGY (\d+)
Value Required MAC_ADDRESS (\S+\.\S+\.\S+)
Value Required FROZEN_TIME (.*\S)

Start
  ^${TOPOLOGY}\s+${MAC_ADDRESS}\s+${FROZEN_TIME} -> Record
  ^Topology\s+Mac Address\s+Frozen time
  ^-+
  ^\s*$$
  ^.* -> Error "LINE NOT FOUND"
`
