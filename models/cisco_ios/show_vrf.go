package cisco_ios

type ShowVrf struct {
	Name       string   `json:"NAME"`
	Default_rd string   `json:"DEFAULT_RD"`
	Protocols  string   `json:"PROTOCOLS"`
	Interfaces []string `json:"INTERFACES"`
}

var ShowVrf_Template string = `Value Required NAME (\S+)
Value DEFAULT_RD ((\d+|\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}):\d+|<not set>)
Value PROTOCOLS (\S+)
Value List INTERFACES (\S+)

Start
  ^\s*Name\s+Default RD\s+Protocols\s+Interfaces -> Start_record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^\s*__Platform_iVRF\:_ID00_
  ^\s*mgmtVrf\s+<not set>
  ^. -> Error

Start_record
  ^\s{2}\S+ -> Continue.Record
  ^\s{60}\s+${INTERFACES}
  ^\s+${NAME}\s+${DEFAULT_RD}\s+${PROTOCOLS}\s+${INTERFACES}\s*$$
  ^\s+${NAME}\s+${DEFAULT_RD}\s+${PROTOCOLS}\s*$$
  ^\s+${NAME}\s+${DEFAULT_RD}\s*$$
  ^\s*Platform\s+iVRF\s+Name -> Start
  ^. -> Error "VRFError"
`
