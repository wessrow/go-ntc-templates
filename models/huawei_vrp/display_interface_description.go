package huawei_vrp 

type DisplayInterfaceDescription struct {
	Interface	string	`json:"INTERFACE"`
	Phy	string	`json:"PHY"`
	Protocol	string	`json:"PROTOCOL"`
	Description	string	`json:"DESCRIPTION"`
}

var DisplayInterfaceDescription_Template = `Value INTERFACE (\S+)
Value PHY (down|\*down|up|up\(s\))
Value PROTOCOL (down|\*down|up|up\(s\))
Value DESCRIPTION (\S+.*?)


Start
  ^Interface\s+PHY\s+Protocol\s+Description\s*$$ -> Begin
  ^\s*$$
  ^PHY:\s+Physical
  ^(?:\*|\^|\#|\-)down:
  ^\(\w+\):\s+\S+
  ^. -> Error

Begin
  ^${INTERFACE}\s+${PHY}\s+${PROTOCOL}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error`