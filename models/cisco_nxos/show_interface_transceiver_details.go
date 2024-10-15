package cisco_nxos

type ShowInterfaceTransceiverDetails struct {
	Voltage_value          string `json:"VOLTAGE_VALUE"`
	Tx_value               string `json:"TX_VALUE"`
	Tx_alarm_high          string `json:"TX_ALARM_HIGH"`
	Temperature_alarm_low  string `json:"TEMPERATURE_ALARM_LOW"`
	Temperature_warn_high  string `json:"TEMPERATURE_WARN_HIGH"`
	Amps_value             string `json:"AMPS_VALUE"`
	Voltage_alarm_low      string `json:"VOLTAGE_ALARM_LOW"`
	Temperature_warn_low   string `json:"TEMPERATURE_WARN_LOW"`
	Amps_warn_high         string `json:"AMPS_WARN_HIGH"`
	Rx_value               string `json:"RX_VALUE"`
	Temperature_alarm_high string `json:"TEMPERATURE_ALARM_HIGH"`
	Temperature_value      string `json:"TEMPERATURE_VALUE"`
	Voltage_alarm_high     string `json:"VOLTAGE_ALARM_HIGH"`
	Amps_alarm_low         string `json:"AMPS_ALARM_LOW"`
	Rx_warn_low            string `json:"RX_WARN_LOW"`
	Interface              string `json:"INTERFACE"`
	Voltage_warn_high      string `json:"VOLTAGE_WARN_HIGH"`
	Voltage_warn_low       string `json:"VOLTAGE_WARN_LOW"`
	Amps_alarm_high        string `json:"AMPS_ALARM_HIGH"`
	Amps_warn_low          string `json:"AMPS_WARN_LOW"`
	Rx_alarm_high          string `json:"RX_ALARM_HIGH"`
	Rx_alarm_low           string `json:"RX_ALARM_LOW"`
	Rx_warn_high           string `json:"RX_WARN_HIGH"`
	Tx_warn_high           string `json:"TX_WARN_HIGH"`
	Lane                   string `json:"LANE"`
	Tx_warn_low            string `json:"TX_WARN_LOW"`
	Tx_alarm_low           string `json:"TX_ALARM_LOW"`
}

var ShowInterfaceTransceiverDetails_Template string = `Value Filldown INTERFACE (\S+/\S+)
Value LANE (\d+)
Value TEMPERATURE_VALUE (([-\+]?\d+\.\d+|(N\/A)))
Value TEMPERATURE_ALARM_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value TEMPERATURE_ALARM_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value TEMPERATURE_WARN_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value TEMPERATURE_WARN_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value VOLTAGE_VALUE (([-\+]?\d+\.\d+|(N\/A)))
Value VOLTAGE_ALARM_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value VOLTAGE_ALARM_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value VOLTAGE_WARN_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value VOLTAGE_WARN_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value AMPS_VALUE (([-\+]?\d+\.\d+|(N\/A)))
Value AMPS_ALARM_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value AMPS_ALARM_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value AMPS_WARN_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value AMPS_WARN_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value RX_VALUE (([-\+]?\d+\.\d+|(N\/A)))
Value RX_ALARM_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value RX_ALARM_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value RX_WARN_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value RX_WARN_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value TX_VALUE (([-\+]?\d+\.\d+|(N\/A)))
Value TX_ALARM_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value TX_ALARM_LOW (([-\+]?\d+\.\d+|(N\/A)))
Value TX_WARN_HIGH (([-\+]?\d+\.\d+|(N\/A)))
Value TX_WARN_LOW (([-\+]?\d+\.\d+|(N\/A)))


Start
  ^${INTERFACE}
  ^Lane Number:${LANE}
  ^\s+Temperature\s+${TEMPERATURE_VALUE}\s+(\w+\s+)?[+-]*\s+${TEMPERATURE_ALARM_HIGH}\s+(\w+\s+)?[+-]*\s+${TEMPERATURE_ALARM_LOW}\s+(\w+\s+)?[+-]*\s+${TEMPERATURE_WARN_HIGH}\s+(\w+\s+)?[+-]*\s+${TEMPERATURE_WARN_LOW}
  ^\s+Voltage\s+${VOLTAGE_VALUE}\s+(\w+\s+)?[+-]*\s+${VOLTAGE_ALARM_HIGH}\s+(\w+\s+)?[+-]*\s+${VOLTAGE_ALARM_LOW}\s+(\w+\s+)?[+-]*\s+${VOLTAGE_WARN_HIGH}\s+(\w+\s+)?[+-]*\s+${VOLTAGE_WARN_LOW}
  ^\s+Current\s+${AMPS_VALUE}\s+(\w+\s+)?[+-]*\s+${AMPS_ALARM_HIGH}\s+(\w+\s+)?[+-]*\s+${AMPS_ALARM_LOW}\s+(\w+\s+)?[+-]*\s+${AMPS_WARN_HIGH}\s+(\w+\s+)?[+-]*\s+${AMPS_WARN_LOW}
  ^\s+Rx\s+Power\s+${RX_VALUE}\s+(\w+\s+)?[+-]*\s+${RX_ALARM_HIGH}\s+(\w+\s+)?[+-]*\s+${RX_ALARM_LOW}\s+(\w+\s+)?[+-]*\s+${RX_WARN_HIGH}\s+(\w+\s+)?[+-]*\s+${RX_WARN_LOW}
  ^\s+Tx\s+Power\s+${TX_VALUE}\s+(\w+\s+)?[+-]*\s+${TX_ALARM_HIGH}\s+(\w+\s+)?[+-]*\s+${TX_ALARM_LOW}\s+(\w+\s+)?[+-]*\s+${TX_WARN_HIGH}\s+(\w+\s+)?[+-]*\s+${TX_WARN_LOW}
  ^\s+Note: -> Record

EOF`
