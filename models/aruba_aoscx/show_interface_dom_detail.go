package aruba_aoscx

type ShowInterfaceDomDetail struct {
	Tx_bias                      []string `json:"TX_BIAS"`
	Tx_bias_low_alarm_trig       []string `json:"TX_BIAS_LOW_ALARM_TRIG"`
	Temp_low_warn_trig           string   `json:"TEMP_LOW_WARN_TRIG"`
	Volt_high_alarm_trig         string   `json:"VOLT_HIGH_ALARM_TRIG"`
	Volt_low_warn_trig           string   `json:"VOLT_LOW_WARN_TRIG"`
	Tx_bias_high_warn_trig       []string `json:"TX_BIAS_HIGH_WARN_TRIG"`
	Tx_bias_low_warn_threshold   []string `json:"TX_BIAS_LOW_WARN_THRESHOLD"`
	Rx_low_alarm_trig            []string `json:"RX_LOW_ALARM_TRIG"`
	Tx_low_alarm_trig            []string `json:"TX_LOW_ALARM_TRIG"`
	Tx_low_warn_trig             []string `json:"TX_LOW_WARN_TRIG"`
	Temp_high_alarm_threshold    string   `json:"TEMP_HIGH_ALARM_THRESHOLD"`
	Tx_low_warn_threshold        []string `json:"TX_LOW_WARN_THRESHOLD"`
	Temp_high_alarm_trig         string   `json:"TEMP_HIGH_ALARM_TRIG"`
	Temp_high_warn_trig          string   `json:"TEMP_HIGH_WARN_TRIG"`
	Volt_low_alarm_trig          string   `json:"VOLT_LOW_ALARM_TRIG"`
	Interface                    string   `json:"INTERFACE"`
	Rx_mw                        []string `json:"RX_MW"`
	Tx_mw                        []string `json:"TX_MW"`
	Temperature                  string   `json:"TEMPERATURE"`
	Volt_high_alarm_threshold    string   `json:"VOLT_HIGH_ALARM_THRESHOLD"`
	Volt_low_alarm_threshold     string   `json:"VOLT_LOW_ALARM_THRESHOLD"`
	Volt_high_warn_threshold     string   `json:"VOLT_HIGH_WARN_THRESHOLD"`
	Tx_bias_high_alarm_trig      []string `json:"TX_BIAS_HIGH_ALARM_TRIG"`
	Tx_bias_low_warn_trig        []string `json:"TX_BIAS_LOW_WARN_TRIG"`
	Tx_bias_low_alarm_threshold  []string `json:"TX_BIAS_LOW_ALARM_THRESHOLD"`
	Rx_high_warn_trig            []string `json:"RX_HIGH_WARN_TRIG"`
	Temp_low_alarm_threshold     string   `json:"TEMP_LOW_ALARM_THRESHOLD"`
	Rx_low_warn_threshold        []string `json:"RX_LOW_WARN_THRESHOLD"`
	Tx_high_warn_trig            []string `json:"TX_HIGH_WARN_TRIG"`
	Rx_low_alarm_threshold       []string `json:"RX_LOW_ALARM_THRESHOLD"`
	Volt_high_warn_trig          string   `json:"VOLT_HIGH_WARN_TRIG"`
	Tx_bias_high_alarm_threshold []string `json:"TX_BIAS_HIGH_ALARM_THRESHOLD"`
	Rx_high_alarm_trig           []string `json:"RX_HIGH_ALARM_TRIG"`
	Rx_low_warn_trig             []string `json:"RX_LOW_WARN_TRIG"`
	Tx_low_alarm_threshold       []string `json:"TX_LOW_ALARM_THRESHOLD"`
	Voltage                      string   `json:"VOLTAGE"`
	Temp_high_warn_threshold     string   `json:"TEMP_HIGH_WARN_THRESHOLD"`
	Volt_low_warn_threshold      string   `json:"VOLT_LOW_WARN_THRESHOLD"`
	Rx_high_warn_threshold       []string `json:"RX_HIGH_WARN_THRESHOLD"`
	Tx_high_alarm_threshold      []string `json:"TX_HIGH_ALARM_THRESHOLD"`
	Tx_high_warn_threshold       []string `json:"TX_HIGH_WARN_THRESHOLD"`
	Channel                      []string `json:"CHANNEL"`
	Temp_low_alarm_trig          string   `json:"TEMP_LOW_ALARM_TRIG"`
	Temp_low_warn_threshold      string   `json:"TEMP_LOW_WARN_THRESHOLD"`
	Tx_bias_high_warn_threshold  []string `json:"TX_BIAS_HIGH_WARN_THRESHOLD"`
	Rx_high_alarm_threshold      []string `json:"RX_HIGH_ALARM_THRESHOLD"`
	Tx_high_alarm_trig           []string `json:"TX_HIGH_ALARM_TRIG"`
	Iface_type                   string   `json:"IFACE_TYPE"`
}

var ShowInterfaceDomDetail_Template string = `Value INTERFACE (\S+)
Value IFACE_TYPE (\S+)
Value List CHANNEL (\d+)
Value TEMPERATURE (-?\d+\.\d+\w)
Value TEMP_HIGH_ALARM_TRIG (\w+)
Value TEMP_LOW_ALARM_TRIG (\w+)
Value TEMP_HIGH_WARN_TRIG (\w+)
Value TEMP_LOW_WARN_TRIG (\w+)
Value TEMP_HIGH_ALARM_THRESHOLD (-?\d+\.\d+\w)
Value TEMP_LOW_ALARM_THRESHOLD (-?\d+\.\d+\w)
Value TEMP_HIGH_WARN_THRESHOLD (-?\d+\.\d+\w)
Value TEMP_LOW_WARN_THRESHOLD (-?\d+\.\d+\w)
Value VOLTAGE (\d+\.\d+\w)
Value VOLT_HIGH_ALARM_TRIG (\w+)
Value VOLT_LOW_ALARM_TRIG (\w+)
Value VOLT_HIGH_WARN_TRIG (\w+)
Value VOLT_LOW_WARN_TRIG (\w+)
Value VOLT_HIGH_ALARM_THRESHOLD (\d+\.\d+\w)
Value VOLT_LOW_ALARM_THRESHOLD (\d+\.\d+\w)
Value VOLT_HIGH_WARN_THRESHOLD (\d+\.\d+\w)
Value VOLT_LOW_WARN_THRESHOLD (\d+\.\d+\w)
Value List TX_BIAS (\d+\.\d+\w+)
Value List TX_BIAS_HIGH_ALARM_TRIG (\w+)
Value List TX_BIAS_LOW_ALARM_TRIG (\w+)
Value List TX_BIAS_HIGH_WARN_TRIG (\w+)
Value List TX_BIAS_LOW_WARN_TRIG (\w+)
Value List TX_BIAS_HIGH_ALARM_THRESHOLD (\d+\.\d+\w+)
Value List TX_BIAS_LOW_ALARM_THRESHOLD (\d+\.\d+\w+)
Value List TX_BIAS_HIGH_WARN_THRESHOLD (\d+\.\d+\w+)
Value List TX_BIAS_LOW_WARN_THRESHOLD (\d+\.\d+\w+)
Value List RX_MW (\d+\.\d+\w+)
Value List RX_HIGH_ALARM_TRIG (\w+)
Value List RX_LOW_ALARM_TRIG (\w+)
Value List RX_HIGH_WARN_TRIG (\w+)
Value List RX_LOW_WARN_TRIG (\w+)
Value List RX_HIGH_ALARM_THRESHOLD (\d+\.\d+\w+)
Value List RX_LOW_ALARM_THRESHOLD (\d+\.\d+\w+)
Value List RX_HIGH_WARN_THRESHOLD (\d+\.\d+\w+)
Value List RX_LOW_WARN_THRESHOLD (\d+\.\d+\w+)
Value List TX_MW (\d+\.\d+\w+)
Value List TX_HIGH_ALARM_TRIG (\w+)
Value List TX_LOW_ALARM_TRIG (\w+)
Value List TX_HIGH_WARN_TRIG (\w+)
Value List TX_LOW_WARN_TRIG (\w+)
Value List TX_HIGH_ALARM_THRESHOLD (\d+\.\d+\w+)
Value List TX_LOW_ALARM_THRESHOLD (\d+\.\d+\w+)
Value List TX_HIGH_WARN_THRESHOLD (\d+\.\d+\w+)
Value List TX_LOW_WARN_THRESHOLD (\d+\.\d+\w+)


Start
  ^Transceiver -> Continue.Record
  ^Transceiver\s+in\s+${INTERFACE} 
  ^\s+Type:\s${IFACE_TYPE}
  ^\s+% No DOM information available -> Next.Clear
  ^\s+Temperature:\s+${TEMPERATURE} -> Common
  ^\s+Channel\s+${CHANNEL}: -> Common
  ^\. -> Error

Common
  ^\s+Temperature high alarm: ${TEMP_HIGH_ALARM_TRIG}
  ^\s+Temperature low alarm: ${TEMP_LOW_ALARM_TRIG}
  ^\s+Temperature high warning: ${TEMP_HIGH_WARN_TRIG}
  ^\s+Temperature low warning: ${TEMP_LOW_WARN_TRIG}
  ^\s+Temperature high alarm threshold: ${TEMP_HIGH_ALARM_THRESHOLD}
  ^\s+Temperature low alarm threshold: ${TEMP_LOW_ALARM_THRESHOLD}
  ^\s+Temperature high warning threshold: ${TEMP_HIGH_WARN_THRESHOLD}
  ^\s+Temperature low warning threshold: ${TEMP_LOW_WARN_THRESHOLD}
  ^\s+Voltage:\s+${VOLTAGE} 
  ^\s+Voltage high alarm: ${VOLT_HIGH_ALARM_TRIG}
  ^\s+Voltage low alarm: ${VOLT_LOW_ALARM_TRIG}
  ^\s+Voltage high warning: ${VOLT_HIGH_WARN_TRIG}
  ^\s+Voltage low warning: ${VOLT_LOW_WARN_TRIG}
  ^\s+Voltage high alarm threshold: ${VOLT_HIGH_ALARM_THRESHOLD}
  ^\s+Voltage low alarm threshold: ${VOLT_LOW_ALARM_THRESHOLD}
  ^\s+Voltage high warning threshold: ${VOLT_HIGH_WARN_THRESHOLD}
  ^\s+Voltage low warning threshold: ${VOLT_LOW_WARN_THRESHOLD}
  ^\s+Tx Bias: ${TX_BIAS} -> Details
  ^\s*$$ -> Start
  ^\. -> Error

Details
  ^\s+Channel\s+${CHANNEL}:
  ^\s+Tx Bias: ${TX_BIAS}
  ^\s+Tx Bias high alarm: ${TX_BIAS_HIGH_ALARM_TRIG}
  ^\s+Tx Bias low alarm: ${TX_BIAS_LOW_ALARM_TRIG}
  ^\s+Tx Bias high warning: ${TX_BIAS_HIGH_WARN_TRIG}
  ^\s+Tx Bias low warning: ${TX_BIAS_LOW_WARN_TRIG}
  ^\s+Tx Bias high alarm threshold: ${TX_BIAS_HIGH_ALARM_THRESHOLD}
  ^\s+Tx Bias low alarm threshold: ${TX_BIAS_LOW_ALARM_THRESHOLD}
  ^\s+Tx Bias high warning threshold: ${TX_BIAS_HIGH_WARN_THRESHOLD}
  ^\s+Tx Bias low warning threshold: ${TX_BIAS_LOW_WARN_THRESHOLD}
  ^\s+Rx Power: ${RX_MW}
  ^\s+Rx Power high alarm: ${RX_HIGH_ALARM_TRIG}
  ^\s+Rx Power low alarm: ${RX_LOW_ALARM_TRIG}
  ^\s+Rx Power high warning: ${RX_HIGH_WARN_TRIG}
  ^\s+Rx Power low warning: ${RX_LOW_WARN_TRIG}
  ^\s+Rx Power high alarm threshold: ${RX_HIGH_ALARM_THRESHOLD}
  ^\s+Rx Power low alarm threshold: ${RX_LOW_ALARM_THRESHOLD}
  ^\s+Rx Power high warning threshold: ${RX_HIGH_WARN_THRESHOLD}
  ^\s+Rx Power low warning threshold: ${RX_LOW_WARN_THRESHOLD}
  ^\s+Tx Power: ${TX_MW}
  ^\s+Tx Power high alarm: ${TX_HIGH_ALARM_TRIG}
  ^\s+Tx Power low alarm: ${TX_LOW_ALARM_TRIG}
  ^\s+Tx Power high warning: ${TX_HIGH_WARN_TRIG}
  ^\s+Tx Power low warning: ${TX_LOW_WARN_TRIG}
  ^\s+Tx Power high alarm threshold: ${TX_HIGH_ALARM_THRESHOLD}
  ^\s+Tx Power low alarm threshold: ${TX_LOW_ALARM_THRESHOLD}
  ^\s+Tx Power high warning threshold: ${TX_HIGH_WARN_THRESHOLD}
  ^\s+Tx Power low warning threshold: ${TX_LOW_WARN_THRESHOLD}
  ^\s*$$
  ^Transceiver -> Continue.Record
  ^Transceiver\s+in\s+${INTERFACE} -> Start
  ^\. -> Error`
