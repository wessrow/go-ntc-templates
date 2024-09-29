package models

type CiscoS300ShowSystemId struct {
	Serial_number	string	`json:"SERIAL_NUMBER"`
}

var CiscoS300ShowSystemId_Template = `Value SERIAL_NUMBER (\S+)

Start
  ^\s*Serial\s+number\s*:\s*${SERIAL_NUMBER}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`