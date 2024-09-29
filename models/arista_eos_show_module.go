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

var AristaEosShowModule_Template = `Value MODULE (\S+)
Value PORTS (\d+)
Value CARD (.+?)
Value TYPE (\S+)
Value MODEL (\S+)
Value SERIAL_NUM (\S+)
Value Fillup MAC_ADDRESS_START (.+?)
Value Fillup MAC_ADDRESS_END (.+?)
Value Fillup HW_VER (\S+)
Value Fillup SW_VER (\S+|\s+)
Value Fillup STATUS (\S+)
Value Fillup UPTIME (.+)

Start
  ^-.+
  ^Module\s+Ports\s+Card\s+Type\s+Model\s+Serial\s+No\.\s*$$
  ^${MODULE}\s+${PORTS}\s+${CARD}\s+${TYPE}\s+${MODEL}\s+${SERIAL_NUM}\s*$$ -> Record
  ^Module\s+MAC\s+addresses\s+Hw\s+Sw\s*$$
  ^${MODULE}\s+(?:${MAC_ADDRESS_START}\s+-\s+${MAC_ADDRESS_END})?\s+${HW_VER}(\s+${SW_VER})?\s*$$
  ^Module\s+Status\s+Uptime\s*$$
  ^${MODULE}\s+${STATUS}(\s+${UPTIME})?\s*$$
  ^\s*$$
  ^. -> Error "LINE NOT FOUND"

EOF

`