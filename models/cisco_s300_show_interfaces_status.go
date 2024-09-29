package models

type CiscoS300ShowInterfacesStatus struct {
	Port	string	`json:"PORT"`
	Type	string	`json:"TYPE"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Neg	string	`json:"NEG"`
	Flowctrl	string	`json:"FLOWCTRL"`
	Linkstate	string	`json:"LINKSTATE"`
	Backpressure	string	`json:"BACKPRESSURE"`
	Mdixmode	string	`json:"MDIXMODE"`
}

var CiscoS300ShowInterfacesStatus_Template = `Value PORT (\S+)
Value TYPE (\S+)
Value DUPLEX (Full|Half|--)
Value SPEED (\d+|--)
Value NEG (Enabled|Disabled|--)
Value FLOWCTRL (Off|On|--)
Value LINKSTATE (Up|Down)
Value BACKPRESSURE (Disabled|Enabled|--)
Value MDIXMODE (Off|On|--)

Start
  ^\s+Flow\s+Link\s+Back\s+Mdix
  ^Port\s+Type\s+Duplex\s+Speed\s+Neg\s+ctrl\s+State\s+Pressure\s+Mode -> Begin
  ^\s*$$
  ^. -> Error

Begin
  ^-+
  ^${PORT}\s+${TYPE}\s+${DUPLEX}\s+${SPEED}\s+${NEG}\s+${FLOWCTRL}\s+${LINKSTATE}\s+${BACKPRESSURE}\s+${MDIXMODE} -> Record
  ^\s+Flow\s+Link
  ^Ch\s+Type\s+Duplex\s+Speed\s+Neg\s+control\s+State -> End

`