package cisco_nxos

type ShowInterfaceDescription struct {
	Port        string `json:"PORT"`
	Type        string `json:"TYPE"`
	Speed       string `json:"SPEED"`
	Description string `json:"DESCRIPTION"`
}

var ShowInterfaceDescription_Template string = `Value PORT (\S+)
Value TYPE (\S+)
Value SPEED (\d+G?)
Value DESCRIPTION (\S.*?)

Start
  ^(Interface\s+Description|Port\s+Type\s+Speed\s+Description)\s*$$ -> Begin

Begin
  ^-+
  ^Port\s+Type\s+Speed\s+Description\s*$$
  ^Interface\s+Description\s*$$
  ^${PORT}\s+${TYPE}\s+${SPEED}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^${PORT}\s+(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
