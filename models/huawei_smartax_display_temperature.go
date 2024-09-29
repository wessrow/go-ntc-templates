package models

type HuaweiSmartaxDisplayTemperature struct {
	Slotid	string	`json:"SLOTID"`
	Boardname	string	`json:"BOARDNAME"`
	Temperature_c	string	`json:"TEMPERATURE_C"`
	Temperature_d	string	`json:"TEMPERATURE_D"`
}

var HuaweiSmartaxDisplayTemperature_Template = `Value SLOTID (\d+)
Value BOARDNAME (\S+)
Value TEMPERATURE_C (\d+)
Value TEMPERATURE_D (\d+)

Start
  ^\s+SlotID:\s+${SLOTID}\s+BoardName:\s+${BOARDNAME}\s+ Temperature:\s+${TEMPERATURE_C}C\(\s*${TEMPERATURE_D}F\) -> Record
  ^\s*$$
  ^. -> Error
`