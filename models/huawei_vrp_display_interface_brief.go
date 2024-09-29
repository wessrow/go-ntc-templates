package models

type HuaweiVrpDisplayInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Phy	string	`json:"PHY"`
	Protocol	string	`json:"PROTOCOL"`
	Inuti	string	`json:"INUTI"`
	Oututi	string	`json:"OUTUTI"`
	Inerrors	string	`json:"INERRORS"`
	Outerrors	string	`json:"OUTERRORS"`
}

var HuaweiVrpDisplayInterfaceBrief_Template = `Value INTERFACE (\S+)
Value PHY (down|[\*\^]down|up|up\(\w+\))
Value PROTOCOL (down|[\*\^]down|up|up\(\w+\))
Value INUTI (\d*\.?\d*%|\-\-)
Value OUTUTI (\d*\.?\d*%|\-\-)
Value INERRORS (\d+)
Value OUTERRORS (\d+)


Start
  ^\s*${INTERFACE}\s+${PHY}\s+${PROTOCOL}\s+${INUTI}\s+${OUTUTI}\s+${INERRORS}\s+${OUTERRORS} -> Record
  ^PHY:\s+Physical
  ^(?:\*|\^|\#)down:
  ^\(\w+\):\s+\S+
  ^InUti/OutUti:
  ^Interface\s+PHY\s+Protocol\s+InUti\s+OutUti\s+inErrors\s+outErrors\s*$$
  ^\s*$$
  ^. -> Error

`