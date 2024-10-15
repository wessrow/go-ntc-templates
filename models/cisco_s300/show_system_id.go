package cisco_s300

type ShowSystemId struct {
	Serial_number string `json:"SERIAL_NUMBER"`
}

var ShowSystemId_Template string = `Value SERIAL_NUMBER (\S+)

Start
  ^\s*Serial\s+number\s*:\s*${SERIAL_NUMBER}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
