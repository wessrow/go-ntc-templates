package models

type HuaweiVrpDisplayTemperature struct {
	Slotid	string	`json:"SLOTID"`
	Pcb	string	`json:"PCB"`
	Status	string	`json:"STATUS"`
	Temperature	string	`json:"TEMPERATURE"`
}

var HuaweiVrpDisplayTemperature_Template = `Value Filldown SLOTID (\S+)
Value Required PCB (\w+)
Value Required STATUS (\w+)
Value TEMPERATURE (\d+)

Start
  ^${SLOTID}\s+: 
  ^${PCB}\s+\d+\s+\d+\s+\d+\s+${STATUS}\s+\d+\s+\d+\s+\d+\s+\d+\s+\d+\s+${TEMPERATURE} -> Record
  ^\s*Base-Board,
  ^PCB\s+I2C\s+Addr\s+Chl\s+Status\s+Minor\s+Major\s+Fatal\s+Adj_speed\s+Temp\s*$$
  ^\s+TMin\s+Tmax\s+\(\S+\)\s*$$
  ^-+
  ^\s*$$
  ^. -> Error

`