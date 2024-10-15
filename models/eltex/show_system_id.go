package eltex

type ShowSystemId struct {
	Unit             string `json:"UNIT"`
	Mac_address      string `json:"MAC_ADDRESS"`
	Hardware_version string `json:"HARDWARE_VERSION"`
	Serial_number    string `json:"SERIAL_NUMBER"`
}

var ShowSystemId_Template string = `Value UNIT (\d+)
Value MAC_ADDRESS ([a-zA-Z0-9]{2}(?::[a-zA-Z0-9]{2}){5})
Value HARDWARE_VERSION (\S+)
Value SERIAL_NUMBER (\S+)

Start
  ^\s*Unit\s+MAC\s+address\s+Hardware\s+version\s+Serial\s+number\s*$$ -> Firmware1
  ^\s*Serial\s+number\s*:\s*${SERIAL_NUMBER}\s*$$
  ^\s*$$
  ^. -> Error

Firmware1
  ^(\s*-+)*\s*$$
  ^\s*${UNIT}\s+${MAC_ADDRESS}\s+${HARDWARE_VERSION}\s+${SERIAL_NUMBER}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
