package arista_eos 

type ShowBootConfig struct {
	Version	string	`json:"VERSION"`
	Console	string	`json:"CONSOLE"`
}

var ShowBootConfig_Template = `Value VERSION (\S+)
Value CONSOLE (.*)

Start
  ^Software image:\s+${VERSION}
  ^Console speed:\s+${CONSOLE} -> Record
  `