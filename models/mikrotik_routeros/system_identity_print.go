package mikrotik_routeros 

type SystemIdentityPrint struct {
	Name	string	`json:"NAME"`
}

var SystemIdentityPrint_Template = `Value NAME (\S+)

Start
  ^\s*name:\s+${NAME}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`