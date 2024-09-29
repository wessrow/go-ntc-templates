package checkpoint_gaia 

type ShowLom struct {
	Version	string	`json:"VERSION"`
	Lomip	string	`json:"LOMIP"`
}

var ShowLom_Template = `Value VERSION (\S+)
Value LOMIP ((?:[0-9]{1,3}\.){3}[0-9]{1,3}(/\d{1,2})?|\w+)

Start
  ^Firmware Revision :\s+${VERSION}
  ^IP Address :\s+${LOMIP} -> Record
`