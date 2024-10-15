package oneaccess_oneos

type ShowVoiceVoicePortAll struct {
	Port            string `json:"PORT"`
	Lp              string `json:"LP"`
	Sense           string `json:"SENSE"`
	Interface_state string `json:"INTERFACE_STATE"`
	Vp_state        string `json:"VP_STATE"`
}

var ShowVoiceVoicePortAll_Template string = `Value Required PORT (\d+)
Value Required LP (\d+)
Value SENSE (\S+|\S+\s\[\S+\])
Value INTERFACE_STATE (\w+)
Value VP_STATE (\w+)

Start
  ^\s*port\s+:\s+${PORT}\s+lp\s+:\s+${LP}\s+sense\s+:\s+${SENSE}\s+(if-state\s+:\s+${INTERFACE_STATE})?\s*\(?vp-state\s*:\s+${VP_STATE}\s*\)? -> Record
  ^\s*$$
  ^. -> Error
`
