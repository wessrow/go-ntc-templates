package models

type CiscoIosShowPowerStatus struct {
	Ps	string	`json:"PS"`
	Input_type	string	`json:"INPUT_TYPE"`
	Input_status	string	`json:"INPUT_STATUS"`
	Model	string	`json:"MODEL"`
	Type	string	`json:"TYPE"`
	Status	string	`json:"STATUS"`
	Fan_sensor	string	`json:"FAN_SENSOR"`
	Inline_status	string	`json:"INLINE_STATUS"`
}