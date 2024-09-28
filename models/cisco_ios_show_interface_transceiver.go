package models

type CiscoIosShowInterfaceTransceiver struct {
	Interface	string	`json:"INTERFACE"`
	Temperature	string	`json:"TEMPERATURE"`
	Voltage	string	`json:"VOLTAGE"`
	Tx_pwr	string	`json:"TX_PWR"`
	Rx_pwr	string	`json:"RX_PWR"`
}