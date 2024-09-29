package models

type ExtremeExosShowPortsDescription struct {
	Interface	string	`json:"INTERFACE"`
	Display_string	string	`json:"DISPLAY_STRING"`
	Description	string	`json:"DESCRIPTION"`
}

var ExtremeExosShowPortsDescription_Template = `Value INTERFACE (\d+)
Value DISPLAY_STRING (\S+)
Value DESCRIPTION (.*)

# Broken. There is no good way to differ 'Display String' and 'Description' since
# both columns can be anything and empty. Since 'Display String' is empty often,
# just skip it.

Start
  ^\s*Port\s+Display\s+String\s+Description\s+String\s*$$ -> Ports
  ^\s*
  ^. -> Error

Ports
  ^(\s*=+)+\s*$$
  ^\s*${INTERFACE}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*
  ^. -> Error

`