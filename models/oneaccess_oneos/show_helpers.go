package oneaccess_oneos

type ShowHelpers struct {
	Interface string   `json:"INTERFACE"`
	Helpers   []string `json:"HELPERS"`
}

var ShowHelpers_Template string = `Value Required INTERFACE (.*)
Value List HELPERS (\d+\.\d+\.\d+\.\d+)

Start
  ^ip\s+forward-protocol
  ^${INTERFACE}:\s\d+\sbroadcasts\sforwarded
  ^\s*${HELPERS}\s*$$ -> Record
  ^\s*${HELPERS} -> Continue
  ^(\s*\d+\.\d+\.\d+\.\d+\s+){1}${HELPERS}\s*$$ -> Record
  ^(\s*\d+\.\d+\.\d+\.\d+\s+){1}${HELPERS} -> Continue
  ^(\s*\d+\.\d+\.\d+\.\d+\s+){2}${HELPERS}\s*$$ -> Record
  ^(\s*\d+\.\d+\.\d+\.\d+\s+){2}${HELPERS} -> Continue
  ^(\s*\d+\.\d+\.\d+\.\d+\s+){3}${HELPERS}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
