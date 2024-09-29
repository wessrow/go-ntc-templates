package models

type MikrotikRouterosSystemIdentityPrint struct {
	Name	string	`json:"NAME"`
}

var MikrotikRouterosSystemIdentityPrint_Template = `Value NAME (\S+)

Start
  ^\s*name:\s+${NAME}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`