package arista_eos

type ShowEnvironmentTemperature struct {
	Critical_limit string `json:"CRITICAL_LIMIT"`
	Module         string `json:"MODULE"`
	Sensor         string `json:"SENSOR"`
	Description    string `json:"DESCRIPTION"`
	Current_value  string `json:"CURRENT_VALUE"`
	Setpoint       string `json:"SETPOINT"`
	Alert_limit    string `json:"ALERT_LIMIT"`
}

var ShowEnvironmentTemperature_Template string = `Value Required SENSOR (\S+)
Value DESCRIPTION ((\s|\S)+?)
Value CURRENT_VALUE (\S+)
Value SETPOINT (\S+)
Value ALERT_LIMIT (\S+)
Value CRITICAL_LIMIT (\S+)
Value Filldown MODULE (\S+\s\d+)

Start
  ^System\stemperature\sstatus\sis
  ^\s+Alert\s+Critical
  ^\s+Temp\s+Setpoint\s+
  ^Sensor\s+Description\s+\(C\)\s+
  ^-+\s+-+\s+-+\s+-+\s+-+\s+-+
  ^${MODULE}:
  ^${SENSOR}\s+${DESCRIPTION}\s+${CURRENT_VALUE}\s+\(${SETPOINT}\)\s+\S+\s+${ALERT_LIMIT}\s+${CRITICAL_LIMIT} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"`
