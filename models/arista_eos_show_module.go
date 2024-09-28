package models

type AristaEosShowModule struct {
	Module	string	`json:"MODULE"`
	Ports	string	`json:"PORTS"`
	Card	string	`json:"CARD"`
	Type	string	`json:"TYPE"`
	Model	string	`json:"MODEL"`
	Serial_num	string	`json:"SERIAL_NUM"`
	Mac_address_start	string	`json:"MAC_ADDRESS_START"`
	Mac_address_end	string	`json:"MAC_ADDRESS_END"`
	Hw_ver	string	`json:"HW_VER"`
	Sw_ver	string	`json:"SW_VER"`
	Status	string	`json:"STATUS"`
	Uptime	string	`json:"UPTIME"`
}