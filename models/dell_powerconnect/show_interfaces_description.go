package dell_powerconnect 

type ShowInterfacesDescription struct {
	Interface	string	`json:"INTERFACE"`
	Description	string	`json:"DESCRIPTION"`
}

var ShowInterfacesDescription_Template = `Value INTERFACE (\S+)
Value DESCRIPTION (\S*)

Start
  ^Port\s+Description -> Begin

Begin
  ^-+
  ^Ch\s+Description -> End
  ^${INTERFACE}\s*${DESCRIPTION}\s*$$ -> Record
  ^\s*
  ^. -> Error
`