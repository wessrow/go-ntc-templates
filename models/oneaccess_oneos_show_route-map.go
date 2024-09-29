package models

type OneaccessOneosShowRouteMap struct {
	Name	string	`json:"NAME"`
	Action	string	`json:"ACTION"`
	Sequence	string	`json:"SEQUENCE"`
	Match_clauses	[]string	`json:"MATCH_CLAUSES"`
	Set_clauses	[]string	`json:"SET_CLAUSES"`
}

var OneaccessOneosShowRouteMap_Template = `Value Required NAME (\S+)
Value Required ACTION (\S+)
Value Required SEQUENCE (\d+)
Value List MATCH_CLAUSES (.+)
Value List SET_CLAUSES (.+)

Start
  ^route-map\s+${NAME}\s${ACTION}\s+${SEQUENCE}\s*$$
  ^\s+match\s+${MATCH_CLAUSES}
  ^\s+set\s+${SET_CLAUSES}
  ^exit -> Record
  ^\s*$$
  ^. -> Error

`