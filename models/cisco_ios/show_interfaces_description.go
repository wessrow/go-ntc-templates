package cisco_ios

type ShowInterfacesDescription struct {
	Port        string `json:"PORT"`
	Status      string `json:"STATUS"`
	Protocol    string `json:"PROTOCOL"`
	Description string `json:"DESCRIPTION"`
}

var ShowInterfacesDescription_Template string = `Value PORT (\S+)
Value STATUS (up|down|deleted|admin\s+down|reset)
Value PROTOCOL (up|down)
Value DESCRIPTION (\S.*?)

Start
  ^Interface\s+Status\s+Protocol\s+Description\s*$$ -> Begin
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

Begin
  ^${PORT}\s+${STATUS}\s+${PROTOCOL}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
