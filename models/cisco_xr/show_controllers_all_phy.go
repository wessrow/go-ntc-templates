package cisco_xr 

type ShowControllersAllPhy struct {
	Interface	string	`json:"INTERFACE"`
	Lane	string	`json:"LANE"`
	Temperature_value	string	`json:"TEMPERATURE_VALUE"`
	Temperature_alarm_high	string	`json:"TEMPERATURE_ALARM_HIGH"`
	Temperature_alarm_low	string	`json:"TEMPERATURE_ALARM_LOW"`
	Temperature_warn_high	string	`json:"TEMPERATURE_WARN_HIGH"`
	Temperature_warn_low	string	`json:"TEMPERATURE_WARN_LOW"`
	Voltage_value	string	`json:"VOLTAGE_VALUE"`
	Voltage_alarm_high	string	`json:"VOLTAGE_ALARM_HIGH"`
	Voltage_alarm_low	string	`json:"VOLTAGE_ALARM_LOW"`
	Voltage_warn_high	string	`json:"VOLTAGE_WARN_HIGH"`
	Voltage_warn_low	string	`json:"VOLTAGE_WARN_LOW"`
	Rx_value	string	`json:"RX_VALUE"`
	Rx_alarm_high	string	`json:"RX_ALARM_HIGH"`
	Rx_alarm_low	string	`json:"RX_ALARM_LOW"`
	Rx_warn_high	string	`json:"RX_WARN_HIGH"`
	Rx_warn_low	string	`json:"RX_WARN_LOW"`
	Tx_value	string	`json:"TX_VALUE"`
	Tx_alarm_high	string	`json:"TX_ALARM_HIGH"`
	Tx_alarm_low	string	`json:"TX_ALARM_LOW"`
	Tx_warn_high	string	`json:"TX_WARN_HIGH"`
	Tx_warn_low	string	`json:"TX_WARN_LOW"`
}

var ShowControllersAllPhy_Template = `Value Filldown INTERFACE (\S+/\S+)
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
Value RX_VALUE (.+?)
Value Filldown RX_ALARM_HIGH (([-\+]?\d+\.\d+))
Value Filldown RX_ALARM_LOW (([-\+]?\d+\.\d+))
Value Filldown RX_WARN_HIGH (([-\+]?\d+\.\d+))
Value Filldown RX_WARN_LOW (([-\+]?\d+\.\d+))
Value TX_VALUE (.+?)
Value Filldown TX_ALARM_HIGH (([-\+]?\d+\.\d+))
Value Filldown TX_ALARM_LOW (([-\+]?\d+\.\d+))
Value Filldown TX_WARN_HIGH (([-\+]?\d+\.\d+))
Value Filldown TX_WARN_LOW (([-\+]?\d+\.\d+))

Start
  ^PHY\s+data\s+for\s+interface:\s+${INTERFACE}
  ^\s+Thresholds:\s+Alarm\s+High\s+Warning\s+High -> Controllers

Controllers
  ^\s+Temperature:\s+${TEMPERATURE_ALARM_HIGH}\s+\w+\s+${TEMPERATURE_WARN_HIGH}\s+\w+\s+${TEMPERATURE_WARN_LOW}\s+\w+\s+${TEMPERATURE_ALARM_LOW}
  ^\s+Voltage:\s+${VOLTAGE_ALARM_HIGH}\s+\w+\s+${VOLTAGE_WARN_HIGH}\s+\w+\s+${VOLTAGE_WARN_LOW}\s+\w+\s+${VOLTAGE_ALARM_LOW}
  ^\s+Transmit Power:.+\(${TX_ALARM_HIGH} dBm\).+\(${TX_WARN_HIGH} dBm\).+\(${TX_WARN_LOW} dBm\).+\(${TX_ALARM_LOW} dBm\)
  ^\s+Receive Power:.+\(${RX_ALARM_HIGH} dBm\).+\(${RX_WARN_HIGH} dBm\).+\(${RX_WARN_LOW} dBm\).+\(${RX_ALARM_LOW} dBm\)
  ^\s+Temperature:\s+${TEMPERATURE_VALUE}
  ^\s+Voltage:\s+${VOLTAGE_VALUE}
  ^\s+Tx\s+Power:.+?\(${TX_VALUE}\)\s*$$
  ^\s+Rx\s+Power:.+?\(${RX_VALUE}\)\s*$$
  ^PHY\s+data\s+for\s+interface: -> Continue.Record
  ^PHY\s+data\s+for\s+interface:\s+${INTERFACE} 
  ^\s+Lane -> Record Lanes

Lanes
  ^\s+Transmit Power:.+\(${TX_ALARM_HIGH} dBm\).+\(${TX_WARN_HIGH} dBm\).+\(${TX_WARN_LOW} dBm\).+\(${TX_ALARM_LOW} dBm\)
  ^\s+Receive Power:.+\(${RX_ALARM_HIGH} dBm\).+\(${RX_WARN_HIGH} dBm\).+\(${RX_WARN_LOW} dBm\).+\(${RX_ALARM_LOW} dBm\)
  ^\s+${LANE}\s+${TEMPERATURE_VALUE}.*\(${TX_VALUE}\s+\S+\).*\(${RX_VALUE}\s+\S+\) -> Record
  ^\S -> Clearall Start
`