package cisco_s300

type ShowInterfacesDescription struct {
	Interface   string `json:"INTERFACE"`
	Description string `json:"DESCRIPTION"`
}

var ShowInterfacesDescription_Template string = `Value INTERFACE (\S+)
Value DESCRIPTION (\S.*?)

Start
  ^\s*Port\s+Description\s*$$ -> Column1
  ^\s*Ch\s+Description\s*$$ -> Column2
  ^\s*$$
  ^. -> Error

Column1
  ^(\s*-*)*\s*$$
  ^\s*Ch\s+Description\s*$$ -> Column2
  ^\s*${INTERFACE}(\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error

Column2
  ^(\s*-*)*\s*$$
  ^\s*${INTERFACE}(\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
