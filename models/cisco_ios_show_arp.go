package models

type CiscoIosShowArp struct {
	Protocol	string	`json:"PROTOCOL"`
	Address	string	`json:"ADDRESS"`
	Age_min	string	`json:"AGE_MIN"`
	Hardware_address	string	`json:"HARDWARE_ADDRESS"`
	Type	string	`json:"TYPE"`
	Interface	string	`json:"INTERFACE"`
}

var CiscoIosShowArp_Template = `Value PROTOCOL (\S+)
Value ADDRESS (\S+)
Value AGE_MIN (\S+)
Value HARDWARE_ADDRESS ((?:([a-f0-9]{4}\.){2}[a-f0-9]{4})|Incomplete)
Value TYPE (\S+)
Value INTERFACE (\S+)

Start
  ^\s*Protocol\s+Address\s+Age\s+\(min\)\s+Hardware\s+Addr\s+Type\s+Interface\s*$$ -> ArpTable
  ^\s*$$
  ^. -> Error

ArpTable
  ^\s*${PROTOCOL}\s+${ADDRESS}\s+${AGE_MIN}\s+${HARDWARE_ADDRESS}\s+${TYPE}(?:\s+${INTERFACE})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`