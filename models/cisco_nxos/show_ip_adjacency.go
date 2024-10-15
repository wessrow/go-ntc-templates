package cisco_nxos

type ShowIpAdjacency struct {
	Source      string `json:"SOURCE"`
	Interface   string `json:"INTERFACE"`
	Flags       string `json:"FLAGS"`
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Pref        string `json:"PREF"`
}

var ShowIpAdjacency_Template string = `Value Required IP_ADDRESS (\S+)
Value Required MAC_ADDRESS (\S+)
Value Required PREF (\d+)
Value Required SOURCE (\S+)
Value Required INTERFACE (\S+)
# Flags:
# # - Adjacencies Throttled for Glean
# G - Adjacencies of vPC peer with G/W bit
Value FLAGS ([G#]*)

Start
  # Ignore headers
  ^.*?-\s+Adjacencies
  ^IP\s+Adjacency\s+Table
  ^Total\s+number\s+of\s+entries:
  # Jump to a list of entries:
  ^Address\s+MAC\s+Address\s+Pref\s+Source\s+Interface -> Entries
  # Process empty and unknown lines
  ^\s*$$
  ^. -> Error

Entries
  # Entry with optional flags
  ^${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${PREF}\s+${SOURCE}\s+${INTERFACE}\s*${FLAGS}$$ -> Record
  # Process empty and unknown lines
  ^\s*$$
  ^. -> Error
`
