package models

type AristaEosShowBootConfig struct {
	Version	string	`json:"VERSION"`
	Console	string	`json:"CONSOLE"`
}

var AristaEosShowBootConfig_Template = `Value VERSION (\S+)
Value CONSOLE (.*)

Start
  ^Software image:\s+${VERSION}
  ^Console speed:\s+${CONSOLE} -> Record
  
`