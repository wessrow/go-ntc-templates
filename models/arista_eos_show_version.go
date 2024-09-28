package models

type AristaEosShowVersion struct {
	Model	string	`json:"MODEL"`
	Hw_version	string	`json:"HW_VERSION"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Sys_mac	string	`json:"SYS_MAC"`
	Image	string	`json:"IMAGE"`
	Total_memory	string	`json:"TOTAL_MEMORY"`
	Free_memory	string	`json:"FREE_MEMORY"`
}