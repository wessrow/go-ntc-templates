package cisco_xr

type ShowInterfacesDescription struct {
	Status      string `json:"STATUS"`
	Protocol    string `json:"PROTOCOL"`
	Description string `json:"DESCRIPTION"`
	Interface   string `json:"INTERFACE"`
}

var ShowInterfacesDescription_Template string = `Value INTERFACE (\S+)
Value STATUS (up|down|deleted|admin\-down|reset)
Value PROTOCOL (up|down|deleted|admin\-down)
Value DESCRIPTION (.+?)

Start
  ^${INTERFACE}\s+${STATUS}\s+${PROTOCOL}\s*(?:${DESCRIPTION})?\s*$$ -> Record
  ^\S+\s+\S+\s+\d+
  ^Interface\s+Status\s+Protocol\s+Description\s*$$
  ^-+\s*$$
  ^\s*$$
  ^. -> Error
`
