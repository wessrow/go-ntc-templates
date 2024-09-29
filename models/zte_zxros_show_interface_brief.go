package models

type ZteZxrosShowInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Attribute	string	`json:"ATTRIBUTE"`
	Mode	string	`json:"MODE"`
	Bw	string	`json:"BW"`
	Admin	string	`json:"ADMIN"`
	Physical	string	`json:"PHYSICAL"`
	Protocol	string	`json:"PROTOCOL"`
	Description	string	`json:"DESCRIPTION"`
}

var ZteZxrosShowInterfaceBrief_Template = `Value INTERFACE (\S+)
Value ATTRIBUTE (\S+)
Value MODE (\S+)
Value BW (\S+)
Value ADMIN (up|down|administratively down)
Value PHYSICAL (up|down|administratively down)
Value PROTOCOL (up|down|administratively down)
Value DESCRIPTION (\S.*?)

Start
  ^Interface\s+Attribute\s+Mode\s+BW\s+Admin\s+Phy\s+Prot\s+Description
  ^${INTERFACE}\s+${ATTRIBUTE}\s+${MODE}\s+${BW}\s+${ADMIN}\s+${PHYSICAL}\s+${PROTOCOL}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`