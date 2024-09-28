package models

type AristaEosShowInterfacesTransceiver struct {
	Port	string	`json:"PORT"`
	Temp	string	`json:"TEMP"`
	Voltage	string	`json:"VOLTAGE"`
	Bias_current	string	`json:"BIAS_CURRENT"`
	Tx_power	string	`json:"TX_POWER"`
	Rx_power	string	`json:"RX_POWER"`
}