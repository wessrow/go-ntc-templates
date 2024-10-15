package cisco_nxos

type ShowInterfaceTransceiver struct {
	Serial       string `json:"SERIAL"`
	Part_number  string `json:"PART_NUMBER"`
	Product_id   string `json:"PRODUCT_ID"`
	Interface    string `json:"INTERFACE"`
	Status       string `json:"STATUS"`
	Manufacturer string `json:"MANUFACTURER"`
	Type         string `json:"TYPE"`
}

var ShowInterfaceTransceiver_Template string = `Value INTERFACE (\S+/\S+)
Value STATUS (not\s\w+|\w+)
Value MANUFACTURER (\S+)
Value TYPE (.+)
Value SERIAL (\S+)
Value PART_NUMBER (\S+)
Value PRODUCT_ID (\S+)

Start
  ^${INTERFACE}
  ^\s+transceiver\sis\s+${STATUS}
  ^\s+type\s+is(\s+${TYPE})?
  ^\s+name\s+is(\s+${MANUFACTURER})?
  ^\s+part\s+number\s+is(\s+${PART_NUMBER})?
  ^\s+serial\s+number\s+is(\s+${SERIAL})?
  ^\s+transceiver\s+
  ^\s+nominal\s+
  ^\s+revision\s+
  ^\s+Link\s+
  ^\s+cable\s+type\s+is
  ^\s+cisco\s+id
  ^\s+cisco\s+extended
  ^\s+cisco\s+part\s+number
  ^\s+cisco\s+product\s+id\sis(\s+${PRODUCT_ID})?
  ^\s+cisco\s+version\s+id
  ^\s+cisco\s+vendor\s+id
  ^\s*$$ -> Record
  ^. -> Error
`
