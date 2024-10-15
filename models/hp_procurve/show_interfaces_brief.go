package hp_procurve

type ShowInterfacesBrief struct {
	Type            string `json:"TYPE"`
	Enabled         string `json:"ENABLED"`
	Mode            string `json:"MODE"`
	Mdi_mode        string `json:"MDI_MODE"`
	Flow_ctrl       string `json:"FLOW_CTRL"`
	Port            string `json:"PORT"`
	Status          string `json:"STATUS"`
	Bcast_limit     string `json:"BCAST_LIMIT"`
	Intrusion_alert string `json:"INTRUSION_ALERT"`
}

var ShowInterfacesBrief_Template string = `Value PORT ([a-zA-Z0-9\-]+)
Value TYPE (\S*)
Value INTRUSION_ALERT (Yes|No)
Value ENABLED (Yes|No)
Value STATUS (Up|Down)
Value MODE ([a-zA-Z0-9\.]+)
Value MDI_MODE (MDI\S*|Auto|NA)
Value FLOW_CTRL (on|off)
Value BCAST_LIMIT (\d+)

Start
  ^\s+Status.*Status\s*
  ^\s*$$
  ^\s*Note:\s*\*.*$$
  ^\s+\|\s+Intrusion\s+MDI\s+Flow(?:\s+Bcast)?\s*$$
  ^\s+Port\s+Type\s+\|\s+Alert\s+Enabled\s+Status\s+Mode\s+Mode\s+Ctrl\s*$$ -> showintbrief1
  ^\s+Port\s+Type\s+\|\s+Alert\s+Enabled\s+Status\s+Mode\s+Mode\s+Ctrl\s+Limit\s*$$ -> showintbrief2
  ^. -> Error

showintbrief1
  ^\s+${PORT}\*?\s+${TYPE}\s+\|\s+${INTRUSION_ALERT}\s+${ENABLED}\s+${STATUS}\s+${MODE}?\*?\s+(?:${MDI_MODE}\s+)?${FLOW_CTRL}\s*$$ -> Record
  ^\s*-+(?:\s|-|\+)+$$
  ^\s*\*\s+third-party.*$$
  ^\s*$$
  ^. -> Error

showintbrief2
  ^\s+${PORT}\*?\s+${TYPE}\s+\|\s+${INTRUSION_ALERT}\s+${ENABLED}\s+${STATUS}\s+${MODE}?\*?\s+(?:${MDI_MODE}\s+)?${FLOW_CTRL}\s+${BCAST_LIMIT}\s*$$ -> Record
  ^\s*-+(?:\s|-|\+)+$$
  ^\s*\*\s+third-party.*$$
  ^\s*$$
  ^. -> Error

`
