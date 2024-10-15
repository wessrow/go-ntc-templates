package cisco_nxos

type ShowRouteMap struct {
	Name          string   `json:"NAME"`
	Action        string   `json:"ACTION"`
	Seq           string   `json:"SEQ"`
	Match_clauses []string `json:"MATCH_CLAUSES"`
	Set_clauses   []string `json:"SET_CLAUSES"`
}

var ShowRouteMap_Template string = `Value Required NAME (\S+)
Value Required ACTION (\S+)
Value Required SEQ (\d+)
Value List MATCH_CLAUSES (.+?)
Value List SET_CLAUSES (.+?)

Start
  ^route-map\s+${NAME},\s+${ACTION},\s+sequence\s+${SEQ}\s*$$
  ^\s+Match\s+clauses -> Match
  ^\s+Set\s+clauses -> Set
  ^\s*$$
  ^. -> Error

Match
  ^\s*$$
  ^\s+Set\s+clauses -> Set
  ^\s+Policy\s+routing
  ^\s+${MATCH_CLAUSES}\s*$$
  ^route-map -> Continue.Record
  ^route-map\s+${NAME},\s+${ACTION},\s+sequence\s+${SEQ}\s*$$ -> Start
  ^\s*$$
  ^. -> Error

Set
  ^\s*$$
  ^\s+Policy\s+routing
  ^\s+${SET_CLAUSES}\s*$$
  ^route-map -> Continue.Record
  ^route-map\s+${NAME},\s+${ACTION},\s+sequence\s+${SEQ}\s*$$ -> Start
  ^\s*$$
  ^. -> Error
`
