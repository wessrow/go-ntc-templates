package ubiquiti_edgeswitch

type ShowVersion struct {
	Hw_revision  string `json:"HW_REVISION"`
	Version      string `json:"VERSION"`
	Switch_type  string `json:"SWITCH_TYPE"`
	Switch_model string `json:"SWITCH_MODEL"`
	Serial       string `json:"SERIAL"`
	Mac_address  string `json:"MAC_ADDRESS"`
}

var ShowVersion_Template string = `Value VERSION (.*)
Value SWITCH_TYPE (\S+\s\S+\s\S+)
Value SWITCH_MODEL (\S+)
Value SERIAL (\S+)
Value MAC_ADDRESS (\S+)
Value HW_REVISION (\S+)

Start
  ^Switch:.+
  ^Machine\s+Type\.+\s+${SWITCH_TYPE}
  ^Machine\s+Model\.+\s+${SWITCH_MODEL}
  ^Serial\s+Number\.+\s+${SERIAL}
  ^Burned\s+In\s+MAC\s+Address\.+\s+${MAC_ADDRESS}
  ^Software\s+Version\.+\s+${VERSION}
  ^Hardware\s+Revision\.+\s+${HW_REVISION}
  # netmiko's send_command starts command output with the last symbol from command sent
  ^\w$$
  ^System\s+Description
  ^Part\s+Number
  ^\s*$$
  ^. -> Error
`
