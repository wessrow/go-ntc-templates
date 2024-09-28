package models

type CiscoNxosShowInterfaceTransceiverDetails struct {
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
	Amps_value	string	`json:"AMPS_VALUE"`
	Amps_alarm_high	string	`json:"AMPS_ALARM_HIGH"`
	Amps_alarm_low	string	`json:"AMPS_ALARM_LOW"`
	Amps_warn_high	string	`json:"AMPS_WARN_HIGH"`
	Amps_warn_low	string	`json:"AMPS_WARN_LOW"`
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