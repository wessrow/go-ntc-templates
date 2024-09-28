package models

type AristaEosShowMacSecurityMkaCounters struct {
	Interface	string	`json:"INTERFACE"`
	Rx_success	string	`json:"RX_SUCCESS"`
	Rx_failure	string	`json:"RX_FAILURE"`
	Tx_success	string	`json:"TX_SUCCESS"`
	Tx_failure	string	`json:"TX_FAILURE"`
}